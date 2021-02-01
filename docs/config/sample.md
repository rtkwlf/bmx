# Setting up user configurations

BMX does not require any configuration properties for execution. If bmx is missing any information it will prompt for the value. The only exception to this is your password. These configuration features are included only as a time-saving measure, or assisting in configuring AWS Account conventions in source.

## Getting Started

A default user configuration is available by running the command `ini-config` like so:

```bash
bmx ini-config --user john.doe --org acmeorg
```

> :information_source: The configuration file follows the format of an ini file.

Which should yield a default user configuration like so (without the comments):

```ini
org                   = acmeorg   # Okta organization
user                  = john.doe  # Your okta username
allow_project_configs = true      # Enable project-scoped configuration
```

You can set this as your user configuration by running the following commands:

```bash
# Ensure the directory exists if it doesn't already
mkdir -p ~/.bmx 

# Write the output from the command into the configuration file
bmx ini-config --user john.doe --org acmeorg > ~/.bmx/config
```

> :warning: **This will overwrite the contents of ~/.bmx/config**: Be very careful here!

# Examples of Configurations

Below are some workflow scenarios and example configurations that help with the workflow.

## Convenient defaults

The following is a configuration that defaults the multi-factor authentication type to `push`, and sets the name of the intended AWS Account. This can be useful when the Okta AWS App you work with is relatively fixed.

```ini
org                   = acmeorg             # Your okta organization
user                  = john.doe            # Your okta username
factor                = push                # Default to using 'push' multi-factor if available
account               = 'AWS ACME'          # The name of the Okta AWS App to use for SSO
role                  = 'ACME-SSODeveloper' # Use the 'Developer' SSO role for the account

allow_project_configs = true                # Enable project-scoped configuration
```

## Limited Access Demo

The following is a configuration that defaults the multi-factor authentication type to `push`, and configures the AWS Account for a demo environment. This can be useful when you have members of the organization that work only in a certain AWS Account like a demo environment:

```ini
org                   = acmeorg                       # Your okta organization
user                  = john.doe                      # Your okta username
factor                = push                          # Default to using 'push' multi-factor if available
account               = 'AWS ACME'                    # The name of the Okta AWS App to use for SSO
role                  = 'ACME-SSODemoRole'            # Select the 'Demo' role in the organization
assume_role           = 'arn:aws:..:role/DemoRole'    # An AWS role arn in the demo environment

allow_project_configs = true                          # Enable project-scoped configuration
```

The `assume_role` setting is used when calling `print`. This will perform a `AssumeRole` action as the last step of `bmx print`. This can be useful to avoid confusion with a single Okta SAML account.