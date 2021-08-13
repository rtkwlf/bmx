package main

import (
	"fmt"
	"log"

	"github.com/rtkwlf/bmx/config"

	"github.com/rtkwlf/bmx/saml/identityProviders/okta"

	"github.com/rtkwlf/bmx"
	"github.com/spf13/cobra"
)

var logoutOptions = bmx.LoginCmdOptions{}

func init() {
	logoutCmd.Flags().StringVar(&logoutOptions.Org, "org", "", "the okta org api to target")
	logoutCmd.Flags().StringVar(&logoutOptions.User, "user", "", "the user to authenticate with")
	logoutCmd.Flags().StringVar(&logoutOptions.Account, "account", "", "the account name to auth against")
	logoutCmd.Flags().StringVar(&logoutOptions.Role, "role", "", "the desired role to assume")
	logoutCmd.Flags().BoolVar(&logoutOptions.NoMask, "nomask", false, "set to not mask the password. this helps with debugging.")

	if userConfig.Org == "" {
		logoutCmd.MarkFlagRequired("org")
	}

	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Revokes the session",
	Long:  `Forces revocation of the session`,
	Run: func(cmd *cobra.Command, args []string) {
		mergedOptions := mergeLogoutOptions(userConfig, logoutOptions)

		oktaClient, err := okta.NewOktaClient(mergedOptions.Org, consolerw)
		if err != nil {
			log.Fatal(err)
		}

		response := bmx.Login(oktaClient, consolerw, mergedOptions)
		fmt.Println(response)
	},
}

func mergeLogoutOptions(uc config.UserConfig, pc bmx.LoginCmdOptions) bmx.LoginCmdOptions {
	if pc.Org == "" {
		pc.Org = uc.Org
	}
	if pc.User == "" {
		pc.User = uc.User
	}
	if pc.Account == "" {
		pc.Account = uc.Account
	}
	if pc.Role == "" {
		pc.Role = uc.Role
	}
	if pc.Factor == "" {
		pc.Factor = uc.Factor
	}

	return pc
}
