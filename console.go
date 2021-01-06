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

package bmx

import (
	"fmt"
	"log"
	"strings"

	"github.com/rtkwlf/bmx/console"
	"github.com/rtkwlf/bmx/saml/identityProviders"
	"github.com/rtkwlf/bmx/saml/identityProviders/okta"
	"github.com/rtkwlf/bmx/saml/serviceProviders"
)

const (
	usernamePrompt = "Okta Username: "
	passwordPrompt = "Okta Password: "
)

func getUserIfEmpty(consolerw console.ConsoleReader, usernameFlag string) string {
	var username string
	if len(usernameFlag) == 0 {
		var err error
		username, err = consolerw.ReadLine(usernamePrompt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		username = usernameFlag
	}
	return username
}

func getPassword(consolerw console.ConsoleReader, noMask bool) string {
	var pass string
	if noMask {
		var err error
		pass, err = consolerw.ReadLine(passwordPrompt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var err error
		pass, err = consolerw.ReadPassword(passwordPrompt)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Fprintln(os.Stderr)
	}
	return pass
}

func selectPrompt(prompt string, options []string, consolerw console.ConsoleReader) (int, error) {
	var selection int
	for idx, option := range options {
		consolerw.Println(fmt.Sprintf("[%d] %s", idx, option))
	}
	selection, err := consolerw.ReadInt(prompt)
	if err != nil {
		return -1, err
	}
	return selection, nil
}

func authenticate(user serviceProviders.UserInfo, oktaClient identityProviders.IdentityProvider, consolerw console.ConsoleReader) (string, error) {
	var userID string
	var ok bool
	userID, ok = oktaClient.AuthenticateFromCache(user.User, user.Org)
	if !ok {
		user.Password = getPassword(consolerw, user.NoMask)
		var err error
		userID, err = oktaClient.Authenticate(user.User, user.Password, user.Org)
		if err != nil {
			log.Fatal(err)
		}
	}

	oktaApplications, err := oktaClient.ListApplications(userID)
	if err != nil {
		log.Fatal(err)
	}

	app, found := findApp(user.Account, oktaApplications)
	if !found {
		// format account labels
		appLabels := []string{}
		for _, app := range oktaApplications {
			appLabels = append(appLabels, app.Label)
		}

		consolerw.Println("Available accounts:")
		accountID, err := selectPrompt("Select an account: ", appLabels, consolerw)
		if err != nil {
			log.Fatal(err)
		}

		app = &oktaApplications[accountID]
	}

	return oktaClient.GetSaml(*app)
}

func findApp(app string, apps []okta.OktaAppLink) (foundApp *okta.OktaAppLink, found bool) {
	for _, v := range apps {
		if strings.ToLower(v.Label) == strings.ToLower(app) {
			return &v, true
		}
	}

	return nil, false
}
