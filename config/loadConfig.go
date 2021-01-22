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

package config

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"

	"github.com/mitchellh/go-homedir"
)

const (
	configFileName  = "config"
	projectFileName = ".bmx"
)

const (
	// UseConsole tries to use the console first, then falls back on
	// other input methods.
	UseConsole = "console"

	// UseAppleScriptLimited tries to use the console first, and uses
	// AppleScript first when in limited.
	UseAppleScriptLimited = "applescript"

	// AlwaysUseAppleScript tries to use AppleScript for input always.
	AlwaysUseAppleScript = "always_applescript"
)

type UserConfig struct {
	AllowProjectConfigs bool
	Org                 string
	User                string
	Account             string
	Role                string
	Profile             string
	AssumeRole          string
	Factor              string
	Input               string
}

func NewUserConfig() UserConfig {
	config := UserConfig{
		AllowProjectConfigs: false,
	}

	return config
}

type ConfigLoader struct {
	UserDirectory    string
	WorkingDirectory string
}

func (d ConfigLoader) LoadConfigs() UserConfig {
	config := NewUserConfig()
	userConfigPath := d.getUserConfig()
	if userConfigPath != "" {
		if err := ini.MapToWithMapper(&config, ini.TitleUnderscore, userConfigPath); err != nil {
			log.Fatalf("Error reflecting config from [%s]\n%s", userConfigPath, err)
		}
	}
	if config.AllowProjectConfigs {
		projectConfigFile := d.FindProjectConfigFile(d.WorkingDir())
		if projectConfigFile != "" {
			if err := ini.MapToWithMapper(&config, ini.TitleUnderscore, projectConfigFile); err != nil {
				log.Fatalf("Error reflecting config from [%s]\n%s", projectConfigFile, err)
			}
		}
	}

	return config
}

func (d ConfigLoader) getUserConfig() string {
	userConfig := filepath.ToSlash(path.Join(d.UserDir(), ".bmx", configFileName))
	if _, err := os.Stat(userConfig); err == nil {
		return userConfig
	}
	return ""
}

func (d ConfigLoader) FindProjectConfigFile(startDir string) string {
	return findFile(path.Join(startDir, projectFileName))
}

func (d ConfigLoader) UserDir() string {
	if d.UserDirectory == "" {
		d.UserDirectory, _ = homedir.Dir()
	}
	return d.UserDirectory
}

func (d ConfigLoader) WorkingDir() string {
	if d.WorkingDirectory == "" {
		d.WorkingDirectory, _ = os.Getwd()
	}
	return d.WorkingDirectory
}

// findFile takes a full path to a file and will recursively search up the directory structure until it finds a file of the
// desired name or it reaches the root directory. If it cannot find the file, it will return an empty string.
func findFile(configPath string) string {
	if info, err := os.Stat(configPath); os.IsNotExist(err) || info.IsDir() {
		configDir := filepath.Dir(configPath)
		parentDir := filepath.Dir(configDir)
		if parentDir == "." || configDir == string(filepath.Separator) {
			return ""
		}

		fileName := filepath.Base(configPath)

		return findFile(filepath.Join(parentDir, fileName))
	}

	return configPath
}
