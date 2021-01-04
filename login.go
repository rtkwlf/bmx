package bmx

import (
	"log"

	"github.com/rtkwlf/bmx/console"
	"github.com/rtkwlf/bmx/saml/identityProviders"

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

func Login(idProvider identityProviders.IdentityProvider, consolerw console.ConsoleReader, loginOptions LoginCmdOptions) string {
	loginOptions.User = getUserIfEmpty(consolerw, loginOptions.User)
	user := GetUserInfoFromLoginCmdOptions(loginOptions)

	_, err := authenticate(user, idProvider, consolerw)
	if err != nil {
		log.Fatal(err)
	}
	return "Session created."
}
