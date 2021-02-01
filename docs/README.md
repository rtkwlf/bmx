# Getting Started: The Basics

## Installing Bmx

BMX is available on GitHub in the [releases](https://github.com/rtkwlf/bmx/releases) page. This lists all available versions of BMX that can be installed. Download and unzip the file

Copy

### Installing from Source

If you'd like to work off the latest changes, you will need to install `golang`. If you don't have it installed, visit [golang's install page]().

Once installed, install bmx from GitHub by running the following command:

```bash
go get -v -u github.com/rtkwlf/bmx/...
```

## Using BMX

The following is a quick introduction to Component through building a simple
static site.  It demonstrates the basic use of component for compiling local
javascript and css files and remote css files.

You can confirm you have successfully configured bmx by running `bmx version`. If you installed from source, then the you should receive: `bmx/nostamp`

### Logging into Okta


```bash
bmx login
```

### Temporary Credentials from IAM

```bash
bmx print
```

### User configuration

```bash
mkdir -p ~/.bmx
touch ~/.bmx/config
```

```ini
org                   = acmeorg   # Okta organization
user                  = john.doe  # Your okta username
allow_project_configs = true      # Enable project-scoped configuration
```

Advanced configuration is covered by [advanced configuration](./config).

### Setting defaults

```ini
factor                = push                # Default to using 'push' multi-factor if available
account               = 'AWS ACME'          # The name of the Okta AWS App to use for SSO
role                  = 'ACME-SSODeveloper' # Use the 'Developer' SSO role for the account
```

### Project Level Configurations

## Next steps

If you want to learn different features of BMX, you can explore the [available commands](./commands) or the [configuration options](./config).

There include examples of configurations and intended use-cases.

-----
