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

package mocks

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"

	bmxaws "github.com/rtkwlf/bmx/saml/serviceProviders/aws"
)

type AwsServiceProviderMock struct{}

func (a AwsServiceProviderMock) GetCredentials(saml string, role bmxaws.AwsRole) *sts.Credentials {
	return &sts.Credentials{
		AccessKeyId:     aws.String("access_key_id"),
		SecretAccessKey: aws.String("secret_access_key"),
		SessionToken:    aws.String("session_token"),
		Expiration:      aws.Time(time.Now().Add(time.Hour * 8)),
	}
}

func (a AwsServiceProviderMock) AssumeRole(creds sts.Credentials, targetRole string, sessionName string) (*sts.Credentials, error) {
	return nil, nil
}

func (a AwsServiceProviderMock) ListRoles(saml string) (roles []bmxaws.AwsRole, err error) {
	var role bmxaws.AwsRole
	role.Name = "test_role"
	return append(roles, role), err
}
