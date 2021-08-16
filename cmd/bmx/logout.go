package main

import (
	"log"

	"github.com/rtkwlf/bmx/config"

	"github.com/rtkwlf/bmx/saml/identityProviders/okta"

	"github.com/rtkwlf/bmx"
	"github.com/spf13/cobra"
)

var logoutOptions = bmx.LogoutCmdOptions{}

func init() {
	logoutCmd.Flags().StringVar(&logoutOptions.Org, "org", "", "the okta org api to target")
	if userConfig.Org == "" {
		logoutCmd.MarkFlagRequired("org")
	}

	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Revokes all Okta sessions",
	Long:  `Forces revocation of the Okta sessions`,
	Run: func(cmd *cobra.Command, args []string) {
		mergedOptions := mergeLogoutOptions(userConfig, logoutOptions)
		oktaClient, err := okta.NewOktaClient(mergedOptions.Org, consolerw)
		if err != nil {
			log.Fatal(err)
		}

		bmx.Logout(oktaClient)
	},
}

func mergeLogoutOptions(uc config.UserConfig, pc bmx.LogoutCmdOptions) bmx.LogoutCmdOptions {
	if pc.Org == "" {
		pc.Org = uc.Org
	}
	return pc
}
