package main

import (
	"log"
	"os"

	"github.com/rtkwlf/bmx/config"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var configTemplate = config.NewUserConfig()

func init() {
	configCmd.Flags().StringVar(&configTemplate.Org, "org", "", "the okta org api to target")
	configCmd.Flags().StringVar(&configTemplate.User, "user", "", "the user to authenticate with")

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "ini-config",
	Short: "Print a minimal configuration for use",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := ini.Empty()

		section := cfg.Section("")
		section.Key("org").SetValue(configTemplate.Org)
		section.Key("user").SetValue(configTemplate.User)
		section.Key("allow_project_configs").SetValue("true")

		_, err := cfg.WriteTo(os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	},
}
