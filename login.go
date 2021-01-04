package bmx

import (
	"fmt"
	"log"

	"github.com/rtkwlf/bmx/console"

	"github.com/rtkwlf/bmx/saml/identityProviders/okta"
	"github.com/rtkwlf/bmx/saml/serviceProviders"
)

type LoginCmdOptions struct {
	Org      string
	User     string
	Account  string
	NoMask   bool
	Password string
	Role     string
	Output   string
}

func GetUserInfoFromLoginCmdOptions(loginOptions LoginCmdOptions) serviceProviders.UserInfo {
	user := serviceProviders.UserInfo{
		Org:      loginOptions.Org,
		User:     loginOptions.User,
		Account:  loginOptions.Account,
		NoMask:   loginOptions.NoMask,
		Password: loginOptions.Password,
		Role:     loginOptions.Role,
	}
	return user
}

func Login(idProvider *okta.OktaClient, consolerw console.ConsoleReader, loginOptions LoginCmdOptions) string {
	loginOptions.User = getUserIfEmpty(consolerw, loginOptions.User)
	user := GetUserInfoFromLoginCmdOptions(loginOptions)

	_, err := authenticate(user, idProvider, consolerw)
	if err != nil {
		log.Fatal(err)
	}

	session, ok := idProvider.GetCachedOktaSession(loginOptions.User, loginOptions.Org)
	if !ok {
		return fmt.Sprintf("Failed to create session for %s", loginOptions.User)
	}
	return fmt.Sprintf("Session for %s expires at %s", session.Userid, session.ExpiresAt)
}
