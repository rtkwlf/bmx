package bmx

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rtkwlf/bmx/saml/serviceProviders/aws"

	"github.com/rtkwlf/bmx/console"
	"github.com/rtkwlf/bmx/saml/identityProviders"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/rtkwlf/bmx/saml/serviceProviders"
)

type CredentialProcessCmdOptions struct {
	Org      string
	User     string
	Account  string
	NoMask   bool
	Password string
	Role     string
	Output   string
	Factor   string
}

type CredentialProcessResult struct {
	Version         int
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

func GetUserInfoFromCredentialProcessCmdOptions(printOptions CredentialProcessCmdOptions) serviceProviders.UserInfo {
	user := serviceProviders.UserInfo{
		Org:      printOptions.Org,
		User:     printOptions.User,
		Account:  printOptions.Account,
		NoMask:   printOptions.NoMask,
		Password: printOptions.Password,
		Role:     printOptions.Role,
		Factor:   printOptions.Factor,
	}
	return user
}

func selectRoleFromSaml(saml string, desiredRole string, awsProvider serviceProviders.ServiceProvider, consolerw console.ConsoleReader) (role aws.AwsRole, err error) {
	roles, err := awsProvider.ListRoles(saml)
	if err != nil {
		return role, err
	}

	if len(roles) == 0 {
		return role, fmt.Errorf("No roles available from SAML")
	}

	if desiredRole == "" {
		roleLabels := []string{}
		for _, role := range roles {
			roleLabels = append(roleLabels, role.Name)
		}

		index, err := consolerw.Option("AWS Account Role", "Select a role: ", roleLabels)
		if err != nil {
			return role, err
		}
		desiredRole = roleLabels[index]
	}
	role, err = aws.FindAwsRoleByName(desiredRole, roles)
	if err != nil {
		return role, err
	}
	return role, nil
}

func CredentialProcess(idProvider identityProviders.IdentityProvider, awsProvider serviceProviders.ServiceProvider, consolerw console.ConsoleReader, printOptions CredentialProcessCmdOptions) string {
	printOptions.User = getUserIfEmpty(consolerw, printOptions.User)
	user := GetUserInfoFromCredentialProcessCmdOptions(printOptions)

	saml, err := authenticate(user, idProvider, consolerw)
	if err != nil {
		log.Fatal(err)
	}

	role, err := selectRoleFromSaml(saml, printOptions.Role, awsProvider, consolerw)
	creds := awsProvider.GetCredentials(saml, role)
	command := credentialProcessCommand(printOptions, creds)
	return command
}

func credentialProcessCommand(printOptions CredentialProcessCmdOptions, creds *sts.Credentials) string {
	result := &CredentialProcessResult{
		Version:         1,
		AccessKeyId:     *creds.AccessKeyId,
		SecretAccessKey: *creds.SecretAccessKey,
		SessionToken:    *creds.SessionToken,
		Expiration:      *creds.Expiration,
	}
	b, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(b)
}
