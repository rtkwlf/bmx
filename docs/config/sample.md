# Setting up user configurations

BMX is designed. Writing all of the configuration options everytime will be tiresome,. BMX is capable of being run without any configuration files set.
<!-- Better outline of how user configurations -->

You'll need to open the user configuration file located at `~/.bmx/config` to set the files. The configuration file follows the format of an ini file.

A default user configuration is available by running the command `ini-config` like so:

```bash
bmx ini-config --user john.doe --org acmeorg
```

Which should yield a default user configuration like so, without the comments:

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

## Convenient defaults

If bmx is missing any information it will prompt you for those. With exception to the password, all parameters can be supplied as configuration values.

The following is a configuration that defaults the multi-factor authentication type to `push`, and sets the name of the intended `AWS Account`. Both of these are useful when working in

```ini
org                   = acmeorg             # Your okta organization
user                  = john.doe            # Your okta username
factor                = push                # Default to using 'push' multi-factor if available
account               = 'AWS ACME'          # The name of the Okta AWS App to use for SSO
role                  = 'ACME-SSODeveloper' # Use the 'Developer' SSO role for the account

allow_project_configs = true                # Enable project-scoped configuration
```

## Limited Access Demo

For members of the organization that may need access to AWS buth only in a limited scope such as a demo environment, it can be useful to level the `assume_role` configuration proper:

```ini
org                   = acmeorg                       # Your okta organization
user                  = john.doe                      # Your okta username
factor                = push                          # Default to using 'push' multi-factor if available
account               = 'AWS ACME'                    # The name of the Okta AWS App to use for SSO
role                  = 'ACME-SSODemoRole'            # Select the 'Demo' role in the organization
assume_role           = 'arn:aws:..:role/DemoRole'    # An AWS role arn in the demo environment

allow_project_configs = true                          # Enable project-scoped configuration
```

When. The `assume_role` setting is used when calling `bmx print`. This will perform an extra 'assume_role' step when performing the role jump. This property can be useful when the Okta AWS App does not exist in the same account as the demo account.