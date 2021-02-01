# Configuration Files

Many of the command line parameters for BMX can be specified in a configuration file located at `~/.bmx/config`. BMX will load this file automatically and populate the parameters where appropriate.

## Configuration directory

The BMX configuration directory contains user-defined parameters settings, such as assume roles, input methods, and so on.

```
~/.bmx/config
```

You cannot change the location of the BMX configuration directory at this time.

To use per-project BMX settings, you can make use of project configurations by creating the file `.bmx` in the project directory.

```
touch .bmx
```

## Project Scoped Configuration

Project configuration scope can be defined by creating a `.bmx` file anywhere in your project's directory structure. When running BMX, it will traverse up the directory structure to . The configuration defined in these files will overwrite the user scoped configuration settings.

To enable project-scoped configurations, you need to enable the user scoped configuration setting `allow_project_configs` like so:

```ini
org                   = acmecorp   
user                  = john.doe

# Enabled to allow project-scope configurations
allow_project_configs = true
```

A project configuration scope can be defined by creating a `.bmx` file anywhere in your project's directory structure. 
When running BMX in the folder with a `.bmx` file or in any folder nested beneath a `.bmx` file, BMX will walk up the 
hierarchy until it finds a `.bmx` file and overlay the configuration with the user scoped configuration file `~/.bmx/config`. 
Note that you must enable this feature with `allow_project_configs=true` in the user configuration file.

## Configuration Settings

|name|type|default|description|
| --- | --- | --- | --- |
| allow_project_configs | bool | `false` | Setting this to true will enable the project scoped configuration feature described below. |
| org | string | `-` | Specify the Okta org to connect to here. This value sets the api base URL for Okta calls (https://{org}.okta.com/). |
| user | string | `-` | This is the username used when connecting to the identity provider. |
| account | string | `-` | The AWS account to retrieve credentials for. |
| role | string | `-` | The AWS SSO SAML role to assume. |
| assume_role | string | `-` | An AWS role arn that will be assumed when calling `print`. |
| profile | string | `-` | The profile to `write` in `~/.aws/credentials`. |
| factor | `(push,token:software:totp)` | `-` | The desired multi-factor authentication factor-type to use. |
| input | `(console,applescript,always_applescript)` | `console` | The intended behaviour for requesting user input. See [input](./input.md) documentation. |
