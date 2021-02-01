# Configuration Files

Many of the command line parameters for BMX can be specified in a configuration files that will be loaded automatically by BMX.

## Configuration directory

The BMX configuration directory contains user-defined configuration settings, such as assume roles, input methods, and so on.

```bash
~/.bmx/config
```

You cannot change the location of the BMX configuration directory at this time.

To use per-project BMX settings, you can make use of project configurations by creating the file `.bmx` in any project directory. When BMX is running from the directory or a subdirectory, it can make use of the configuration options set in that configuration file.

```ini
# projectABC/.bmx
account = 'AWS ACME - ABC' # The  Okta AWS App name for projectABC
```

## Project Scoped Configuration

A project scoped configuration can be defined by creating a `.bmx` file anywhere in your project's directory structure. When running BMX, it will traverse up the directory structure to find any project scoped configurations. The properties defined in these files will overwrite the user scoped configuration settings.

To enable project-scoped configurations, you need to enable the user scoped configuration setting `allow_project_configs` like so:

```ini
# ~/.bmx/config
org                   = acmecorp   
user                  = john.doe

# Enabled to allow project-scope configurations
allow_project_configs = true
```

## Configuration Settings

|name|type|default|description|
| --- | --- | --- | --- |
| allow_project_configs | bool | `false` | If `true` enables project scoped configurations, otherwise disabled |
| org | string | `-` | The Okta organization name. See your okta url https://{org}.okta.com/ |
| user | string | `-` | Your Okta username. Typically `first.lastname` |
| account | string | `-` | The Okta AWS App name for the AWS Account. This is not the same as the AWS account name |
| role | string | `-` | The AWS SSO SAML role to use when connecting to the AWS Account |
| assume_role | string | `-` | An AWS role arn that will be assumed as the last step in `bmx print` |
| profile | string | `-` | The name of the write profile for `bmx write` |
| factor | `(push,token:software:totp)` | `-` | The desired multi-factor authentication factor-type to use |
| input | `(console,applescript,always_applescript)` | `console` | The intended behaviour for requesting user input. See [input](./input.md) documentation. |
