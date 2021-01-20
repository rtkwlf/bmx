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

package aws

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/rtkwlf/bmx/console"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

func NewAwsServiceProvider(consolerw console.ConsoleReader) *AwsServiceProvider {
	awsSession, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	stsClient := sts.New(awsSession)

	serviceProvider := &AwsServiceProvider{
		StsClient:   stsClient,
		InputReader: consolerw,
		UserOutput:  os.Stderr,
	}
	return serviceProvider
}

type AwsServiceProvider struct {
	StsClient   stsiface.STSAPI
	InputReader console.ConsoleReader
	UserOutput  *os.File
}

func (a AwsServiceProvider) ListRoles(saml string) (roles []AwsRole, err error) {
	decodedSaml, err := base64.StdEncoding.DecodeString(saml)
	if err != nil {
		return roles, err
	}

	samlResponse := &Saml2pResponse{}
	err = xml.Unmarshal(decodedSaml, samlResponse)
	if err != nil {
		return roles, err
	}

	roles = listRoles(samlResponse)
	return roles, err
}

func (a AwsServiceProvider) GetCredentials(saml string, desiredRole string) *sts.Credentials {
	decodedSaml, err := base64.StdEncoding.DecodeString(saml)
	if err != nil {
		log.Fatal(err)
	}

	samlResponse := &Saml2pResponse{}
	err = xml.Unmarshal(decodedSaml, samlResponse)
	if err != nil {
		log.Fatal(err)
	}

	var role AwsRole
	roles := listRoles(samlResponse)

	if desiredRole == "" {
		role = a.pickRole(roles)
	} else {
		role = findRole(roles, desiredRole)
	}

	samlInput := &sts.AssumeRoleWithSAMLInput{
		PrincipalArn:  aws.String(role.Principal),
		RoleArn:       aws.String(role.ARN),
		SAMLAssertion: aws.String(saml),
	}

	out, err := a.StsClient.AssumeRoleWithSAML(samlInput)
	if err != nil {
		log.Fatal(err)
	}

	return out.Credentials
}

func (a AwsServiceProvider) AssumeRole(creds sts.Credentials, targetRole string, sessionName string) (*sts.Credentials, error) {
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     aws.StringValue(creds.AccessKeyId),
			SecretAccessKey: aws.StringValue(creds.SecretAccessKey),
			SessionToken:    aws.StringValue(creds.SessionToken),
		})},
	})
	if err != nil {
		log.Fatal(err)
	}

	stsClient := sts.New(awsSession)
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(targetRole),
		RoleSessionName: aws.String(sessionName),
	}
	out, err := stsClient.AssumeRole(input)
	if err != nil {
		return nil, err
	}

	return &sts.Credentials{
		AccessKeyId:     out.Credentials.AccessKeyId,
		SecretAccessKey: out.Credentials.SecretAccessKey,
		SessionToken:    out.Credentials.SessionToken,
		Expiration:      out.Credentials.Expiration,
	}, nil

}

func findRole(roles []AwsRole, desiredRole string) AwsRole {
	desiredRole = strings.ToLower(desiredRole)
	for _, role := range roles {
		if strings.Compare(strings.ToLower(role.Name), desiredRole) == 0 {
			return role
		}
	}

	log.Fatalf("Unable to find desired role [%s]", desiredRole)
	return AwsRole{}
}

func (a AwsServiceProvider) pickRole(roles []AwsRole) AwsRole {
	if len(roles) == 1 {
		return roles[0]
	}

	for i, role := range roles {
		fmt.Fprintf(a.UserOutput, "[%d] %s\n", i, role.Name)
	}
	j, _ := a.InputReader.ReadInt("Select a role: ")

	return roles[j]
}

func listRoles(samlResponse *Saml2pResponse) []AwsRole {
	var roles []AwsRole
	for _, v := range samlResponse.Assertion.AttributeStatement.Attributes {
		if v.Name == "https://aws.amazon.com/SAML/Attributes/Role" {
			for _, w := range v.Values {
				splitRole := strings.Split(w, ",")
				role := AwsRole{}
				role.Principal = splitRole[0]
				role.ARN = splitRole[1]
				role.Name = strings.SplitAfter(role.ARN, "role/")[1]

				roles = append(roles, role)
			}
		}
	}
	return roles
}

type Saml2pResponse struct {
	XMLName   xml.Name       `xml:"Response"`
	Assertion Saml2Assertion `xml:"Assertion"`
}

type Saml2Assertion struct {
	XMLName            xml.Name                `xml:"Assertion"`
	AttributeStatement Saml2AttributeStatement `xml:"AttributeStatement"`
}

type Saml2AttributeStatement struct {
	XMLName    xml.Name         `xml:"AttributeStatement"`
	Attributes []Saml2Attribute `xml:"Attribute"`
}

type Saml2Attribute struct {
	XMLName xml.Name `xml:"Attribute"`
	Name    string   `xml:"Name,attr"`
	Values  []string `xml:"AttributeValue"`
}

type AwsRole struct {
	Name      string
	ARN       string
	Principal string
}
