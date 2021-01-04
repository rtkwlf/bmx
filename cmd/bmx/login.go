package main

import (
	"fmt"
	"log"

	"github.com/rtkwlf/bmx/config"

	"github.com/rtkwlf/bmx/saml/identityProviders/okta"
	"github.com/rtkwlf/bmx/saml/serviceProviders/aws"

	"github.com/rtkwlf/bmx"
	"github.com/spf13/cobra"
)

var loginOptions = bmx.LoginCmdOptions{}

func init() {
	loginCmd.Flags().StringVar(&loginOptions.Org, "org", "", "the okta org api to target")
	loginCmd.Flags().StringVar(&loginOptions.User, "user", "", "the user to authenticate with")
	loginCmd.Flags().StringVar(&loginOptions.Account, "account", "", "the account name to auth against")
	loginCmd.Flags().StringVar(&loginOptions.Role, "role", "", "the desired role to assume")
	loginCmd.Flags().BoolVar(&loginOptions.NoMask, "nomask", false, "set to not mask the password. this helps with debugging.")
	loginCmd.Flags().StringVar(&loginOptions.Output, "output", "", "the output format [bash|powershell]")

	if userConfig.Org == "" {
		loginCmd.MarkFlagRequired("org")
	}

	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create a session",
	Long:  `If logged out, create a new session`,
	Run: func(cmd *cobra.Command, args []string) {
		mergedOptions := mergePrintOptions(userConfig, loginOptions)

		oktaClient, err := okta.NewOktaClient(mergedOptions.Org, consolerw)
		if err != nil {
			log.Fatal(err)
		}

		awsProvider := aws.NewAwsServiceProvider(consolerw)
		command := bmx.Print(oktaClient, awsProvider, consolerw, mergedOptions)
		fmt.Println(command)
	},
}

func mergeLoginOptions(uc config.UserConfig, pc bmx.LoginCmdOptions) bmx.LoginCmdOptions {
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

	return pc
}
