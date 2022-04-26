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

package bmx_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/rtkwlf/bmx"
	"github.com/rtkwlf/bmx/mocks"
	awsmocks "github.com/rtkwlf/bmx/saml/serviceProviders/aws/mocks"
)

func assertAwsTokenEnv(t *testing.T, output string) {
	awsStsVars := [3]string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	}

	for _, envVar := range awsStsVars {
		if !strings.Contains(output, envVar) {
			t.Errorf("Environment variable %s missing, got: %s", envVar, output)
		}
	}
}

func TestMonkey(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org: "myorg",
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
}

func TestPShellPrint(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org:    "myorg",
		Output: bmx.Powershell,
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
	if !strings.Contains(output, "$env:") {
		t.Errorf("Shell command was incorrect, got: %s, expected powershell", output)
	}
}

func TestBashPrint(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org:    "myorg",
		Output: bmx.Bash,
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	assertAwsTokenEnv(t, output)
	if !strings.Contains(output, "export ") {
		t.Errorf("Shell command was incorrect, got: %s, expected bash", output)
	}
}

func TestJsonPrint(t *testing.T) {
	options := bmx.PrintCmdOptions{
		Org:    "myorg",
		Output: bmx.Json,
	}

	oktaClient := &mocks.Mokta{}

	consolerw := mocks.ConsoleReaderMock{}
	awsProvider := awsmocks.AwsServiceProviderMock{}

	output := bmx.Print(oktaClient, awsProvider, consolerw, options)

	if strings.Contains(output, "export ") {
		t.Errorf("Shell command was incorrect, got: %s, expected json", output)
	}
	if strings.Contains(output, "$env:") {
		t.Errorf("Shell command was incorrect, got: %s, expected json", output)
	}
	var creds *sts.Credentials
	err := json.Unmarshal([]byte(output), &creds)
	if err != nil {
		t.Errorf("Did not receive valid JSON")
	}

	if *creds.AccessKeyId == "" || *creds.SecretAccessKey == "" || *creds.SessionToken == "" {
		t.Errorf("what")
	}
}
