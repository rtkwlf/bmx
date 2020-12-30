/*
Copyright 2019 D2L Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package okta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/rtkwlf/bmx/console"
	"github.com/rtkwlf/bmx/saml/identityProviders/okta/file"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

const (
	applicationJson = "application/json"
)

func NewOktaClient(org string, consolerw console.ConsoleReader) (*OktaClient, error) {
	// All users of cookiejar should import "golang.org/x/net/publicsuffix"
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
	}

	oktaSessionStorage := &file.OktaSessionStorage{}

	client := &OktaClient{
		HttpClient:    httpClient,
		SessionCache:  oktaSessionStorage,
		ConsoleReader: consolerw,
		Timeout:       2 * time.Second,
		Retries:       15,
	}

	client.BaseUrl, _ = url.Parse(fmt.Sprintf("https://%s.okta.com/api/v1/", org))

	return client, nil
}

type SessionCache interface {
	SaveSessions(sessions []file.OktaSessionCache)
	Sessions() ([]file.OktaSessionCache, error)
}

type OktaClient struct {
	HttpClient    *http.Client
	SessionCache  SessionCache
	ConsoleReader console.ConsoleReader
	BaseUrl       *url.URL
	Timeout       time.Duration
	Retries       int
}

func (o *OktaClient) GetSaml(appLink OktaAppLink) (string, error) {
	appResponse, err := o.HttpClient.Get(appLink.LinkUrl)
	if err != nil {
		log.Fatal(err)
	}

	return GetSaml(appResponse.Body)
}

func (o *OktaClient) Authenticate(username, password, org string) (string, error) {
	rel, err := url.Parse("authn")
	if err != nil {
		return "", err
	}

	url := o.BaseUrl.ResolveReference(rel)
	if err != nil {
		return "", err
	}

	body := fmt.Sprintf(`{"username":"%s", "password":"%s"}`, username, password)
	authResponse, err := o.HttpClient.Post(url.String(), applicationJson, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	if authResponse.StatusCode != 200 {
		z, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			return "", err
		}
		eResp := &errorResponse{}
		err = json.Unmarshal(z, &eResp)
		if err != nil {
			return "", fmt.Errorf("Received invalid response from okta.\nReponse code: %q\nBody:%s", authResponse.Status, body)
		}
		return "", fmt.Errorf("%s. Response code: %q", eResp.Summary, authResponse.Status)
	}

	oktaAuthResponse := &OktaAuthResponse{}
	z, err := ioutil.ReadAll(authResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(z, &oktaAuthResponse)
	if err != nil {
		log.Fatal(err)
	}

	if err := o.doMfa(oktaAuthResponse); err != nil {
		log.Fatal(err)
	}

	oktaSessionResponse, err := o.startSession(oktaAuthResponse.SessionToken)
	o.setSessionId(oktaSessionResponse.Id)
	o.CacheOktaSession(username, org, oktaSessionResponse.Id, oktaSessionResponse.ExpiresAt)

	return oktaSessionResponse.UserId, err
}

func (o *OktaClient) AuthenticateFromCache(username, org string) (string, bool) {
	sessionID, ok := o.GetCachedOktaSession(username, org)
	if !ok {
		return "", false
	}

	o.setSessionId(sessionID)

	rel, _ := url.Parse(fmt.Sprintf("users/me"))
	url := o.BaseUrl.ResolveReference(rel)

	meRequest, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", false
	}
	meResponse, err := o.HttpClient.Do(meRequest)
	if err != nil {
		return "", false
	}
	var me OktaMeResponse
	b, err := ioutil.ReadAll(meResponse.Body)
	if err != nil {
		return "", false
	}
	err = json.Unmarshal(b, &me)
	if err != nil {
		return "", false
	}
	return me.Id, true
}

func (o *OktaClient) ListApplications(userId string) ([]OktaAppLink, error) {
	rel, _ := url.Parse(fmt.Sprintf("users/%s/appLinks", userId))
	url := o.BaseUrl.ResolveReference(rel)

	listApplicationRequest, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	listApplicationsResponse, err := o.HttpClient.Do(listApplicationRequest)
	if err != nil {
		return nil, err
	}
	var oktaApplications []OktaAppLink
	b, err := ioutil.ReadAll(listApplicationsResponse.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &oktaApplications)
	if err != nil {
		return nil, err
	}

	return oktaApplications, nil
}

func (o *OktaClient) startSession(sessionToken string) (*OktaSessionResponse, error) {
	rel, err := url.Parse("sessions")
	if err != nil {
		return nil, err
	}
	url := o.BaseUrl.ResolveReference(rel)
	if err != nil {
		return nil, err
	}
	oktaSessionsRequest := OktaSessionsRequest{
		SessionToken: sessionToken,
	}
	b, err := json.Marshal(oktaSessionsRequest)
	if err != nil {
		return nil, err
	}
	sessionResponse, err := o.HttpClient.Post(url.String(), applicationJson, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	oktaSessionResponse := &OktaSessionResponse{}
	b, err = ioutil.ReadAll(sessionResponse.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, oktaSessionResponse)
	if err != nil {
		return nil, err
	}

	return oktaSessionResponse, nil
}

func (o *OktaClient) CacheOktaSession(userId, org, sessionId, expiresAt string) {
	session := file.OktaSessionCache{
		Userid:    userId,
		Org:       org,
		SessionId: sessionId,
		ExpiresAt: expiresAt,
	}
	existingSessions, err := readOktaCacheSessionsFile(o)
	if err != nil {
		return
	}
	existingSessions = append(existingSessions, session)
	o.SessionCache.SaveSessions(existingSessions)
}

func (o *OktaClient) GetCachedOktaSession(userid, org string) (string, bool) {
	oktaSessions, err := readOktaCacheSessionsFile(o)
	if err != nil {
		return "", false
	}
	for _, oktaSession := range oktaSessions {
		if oktaSession.Userid == userid &&
			oktaSession.Org == org {
			return oktaSession.SessionId, true
		}
	}
	return "", false
}

func readOktaCacheSessionsFile(o *OktaClient) ([]file.OktaSessionCache, error) {
	sessions, err := o.SessionCache.Sessions()
	if err != nil {
		return nil, err
	}
	return removeExpiredOktaSessions(sessions), nil
}

func (o *OktaClient) setSessionId(id string) {
	cookies := o.HttpClient.Jar.Cookies(o.BaseUrl)
	cookie := &http.Cookie{
		Name:     "sid",
		Value:    id,
		Path:     "/",
		Domain:   o.BaseUrl.Host,
		Secure:   true,
		HttpOnly: true,
	}
	cookies = append(cookies, cookie)
	o.HttpClient.Jar.SetCookies(o.BaseUrl, cookies)
}

func removeExpiredOktaSessions(sourceCaches []file.OktaSessionCache) []file.OktaSessionCache {
	var returnCache []file.OktaSessionCache
	curTime := time.Now()
	for _, sourceCache := range sourceCaches {
		expireTime, err := time.Parse(time.RFC3339, sourceCache.ExpiresAt)
		if err == nil && expireTime.After(curTime) {
			returnCache = append(returnCache, sourceCache)
		}
	}
	return returnCache
}

func (o *OktaClient) verifyPushMfa(oktaAuthResponse *OktaAuthResponse, selectedFactor OktaAuthFactors) error {
	verified := false
	for retry := 0; retry < o.Retries; retry++ {
		body := fmt.Sprintf(`{"stateToken":"%s"}`, oktaAuthResponse.StateToken)
		authResponse, err := o.HttpClient.Post(selectedFactor.Links.Verify.Url, "application/json", strings.NewReader(body))
		if err != nil {
			return err
		}

		z, _ := ioutil.ReadAll(authResponse.Body)
		if err := json.Unmarshal(z, &oktaAuthResponse); err != nil {
			return err
		}

		if oktaAuthResponse.Status == "SUCCESS" {
			verified = true
			break
		} else if oktaAuthResponse.Status == "MFA_CHALLENGE" || oktaAuthResponse.Status == "WAITING" {
			time.Sleep(o.Timeout)
		}
	}

	if !verified {
		return fmt.Errorf("Failed to verify challenge within timeout window.")
	}

	return nil
}

func (o *OktaClient) verifyTotpMfa(oktaAuthResponse *OktaAuthResponse, selectedFactor OktaAuthFactors) error {
	code, err := o.ConsoleReader.ReadLine("Code: ")
	if err != nil {
		return err
	}
	body := fmt.Sprintf(`{"stateToken":"%s","passCode":"%s"}`, oktaAuthResponse.StateToken, code)
	authResponse, err := o.HttpClient.Post(selectedFactor.Links.Verify.Url, "application/json", strings.NewReader(body))
	if err != nil {
		return err
	}

	z, _ := ioutil.ReadAll(authResponse.Body)
	if err := json.Unmarshal(z, &oktaAuthResponse); err != nil {
		return err
	}

	return nil
}

func (o *OktaClient) doMfa(oktaAuthResponse *OktaAuthResponse) error {
	if oktaAuthResponse.Status == "MFA_REQUIRED" {
		o.ConsoleReader.Println("MFA Required")
		for idx, factor := range oktaAuthResponse.Embedded.Factors {
			o.ConsoleReader.Println(fmt.Sprintf("%d - %s", idx, factor.FactorType))
		}

		var mfaIdx int
		var err error
		if mfaIdx, err = o.ConsoleReader.ReadInt("Select an available MFA option: "); err != nil {
			log.Fatal(err)
		}
		selectedFactor := oktaAuthResponse.Embedded.Factors[mfaIdx]
		vurl := oktaAuthResponse.Embedded.Factors[mfaIdx].Links.Verify.Url

		body := fmt.Sprintf(`{"stateToken":"%s"}`, oktaAuthResponse.StateToken)
		authResponse, err := o.HttpClient.Post(vurl, "application/json", strings.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}

		z, _ := ioutil.ReadAll(authResponse.Body)
		err = json.Unmarshal(z, &oktaAuthResponse)
		if err != nil {
			log.Fatal(err)
		}

		// This is a rough outline and can be better organized. For now
		// I'm comfortable with adding in this kind of handling for
		// multiple MFA factors. I'd like for this to be done in a
		// mapped action form (e.g. actions[factortype] => perform action)
		if selectedFactor.FactorType == "token:software:totp" {
			err = o.verifyTotpMfa(oktaAuthResponse, selectedFactor)
			if err != nil {
				log.Fatal(err)
			}
		} else if selectedFactor.FactorType == "push" {
			err = o.verifyPushMfa(oktaAuthResponse, selectedFactor)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

type OktaAuthResponse struct {
	ExpiresAt    time.Time                `json:"expiresAt"`
	SessionToken string                   `json:"sessionToken"`
	StateToken   string                   `json:"stateToken"`
	Status       string                   `json:"status"`
	Embedded     OktaAuthResponseEmbedded `json:"_embedded"`
}

type errorResponse struct {
	Code    string         `json:"errorCode"`
	Summary string         `json:"errorSummary"`
	Link    string         `json:"errorLink"`
	ErrorId string         `json:"errorId"`
	Causes  []errorSummary `json:"errorCauses"`
}

type errorSummary struct {
	Summary string `json:"errorSummary"`
}

type OktaAuthResponseEmbedded struct {
	Factors []OktaAuthFactors `json:"factors"`
}

type OktaAuthFactors struct {
	Id         string    `json:"id"`
	FactorType string    `json:"factorType"`
	Links      OktaLinks `json:"_links"`
}

type OktaLinks struct {
	Verify OktaVerifyLink `json:"verify"`
}

type OktaVerifyLink struct {
	Url string `json:"href"`
}

type OktaSessionsRequest struct {
	SessionToken string `json:"sessionToken"`
}

type OktaSessionResponse struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ExpiresAt string `json:"expiresAt"`
}

type OktaMeResponse struct {
	Id string `json:"id"`
}

type OktaAppLink struct {
	Id            string `json:"id"`
	Label         string `json:"label"`
	LinkUrl       string `json:"linkUrl"`
	AppName       string `json:"appName"`
	AppInstanceId string `json:"appInstanceId"`
}

func GetSaml(r io.Reader) (string, error) {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return "", z.Err()
		case html.SelfClosingTagToken:
			tn, hasAttr := z.TagName()

			if string(tn) == "input" {
				attr := make(map[string]string)
				for hasAttr {
					key, val, moreAttr := z.TagAttr()
					attr[string(key)] = string(val)
					if !moreAttr {
						break
					}
				}

				if attr["name"] == "SAMLResponse" {
					return string(attr["value"]), nil
				}
			}
		}
	}
}
