# Getting Started: The Basics

## Installing Bmx

BMX is available on GitHub in the [releases](https://github.com/rtkwlf/bmx/releases) page. This lists all available versions of BMX that can be installed. Download and unzip the file, copying the `bmx` binary into a folder in your `$PATH`.

> :warning: **If you are a member of Arctic Wolf**: Do not install from here. Consult internal docs.

### Installing from Source

If you'd like to work off the latest changes, you will need to install `golang`. If you don't have it installed, visit [golang's install page](https://golang.org/doc/install).

Once installed, install bmx from GitHub by running the following command:

```bash
go get -v -u github.com/rtkwlf/bmx/...
```

> :warning: **If you are a member of Arctic Wolf**: Do not install from here. Consult internal docs.

## Using BMX

The following is a quick introduction to BMX for authentication to AWS Accounts. For background, BMX uses SAML resources configured in an AWS Account to retrieve IAM STS credentials. These can then be used to interact with the AWS API.

Setting up an Okta AWS App in an AWS Account requires creating IAM roles in an AWS Account that will be assumed with SAML. Typically this account is a 'IAM/SSO'-only account. This means that no services are deployed in the account, and the IAM roles in the account are restricted to `AssumeRole` permissions. From this account, you can `AssumeRole` into other accounts in the organization where you will have development permissions (as needed).

Sometimes this may be different, resulting in multiple Okta AWS Apps for a single AWS organization. BMX supports this through the `account` parameter.

When you have successfully installed BMX into your environment, you can confirm it by running the command `bmx version`. If you installed from source, you will see `bmx/nostamp`.

### Logging into Okta

To authenticate and obtain a session via the command line, enter the following via the command line:

```bash
bmx login
```

This will prompt you for your Okta organization and credentials. When complete, you will be told the session duration. This can be useful when starting for the day, as depending on the duration granted by the SAML IAM roles it may exceed the traditional workday.

You can consult [configuration](./config) to reduce the amount of inputs needed.

### Temporary Credentials from IAM

To retrieve temporary credentials for communicating with the AWS API, you can run the following via the command line:

```bash
bmx print --output bash
```

If you'd like to source these credentials in your shell, you can run `print` as such:

```bash
`bmx print --output bash`
```

### User configuration

To avoid specifying the organization each time you run BMX, you can use the starter configuration available through `ini-config`. You can start by creating the BMX configuration directory via the command line:

```bash
mkdir -p ~/.bmx
touch ~/.bmx/config
```

You can then create the starter configuration via the command line:

```bash
# Fill in the correct username and organization for Okta
bmx ini-config --user john.doe --org acmeorg > ~/.bmx/config

# Read out the contents of the configuration file
cat ~/.bmx/config
```

The configuration file should look like the following:

```ini
org                   = acmeorg   # Your okta organization
user                  = john.doe  # Your okta username
allow_project_configs = true      # Enable project-scoped configuration
```

You can learn more about configuration options in the [configuration documentation](./config). It would also be advisable to setup the convenient defaults as seen in the [sample configurations](./config/sample.md).

## Next steps

If you want to learn different features of BMX, you can explore the available command using the help option, or explore the [configuration options](./config). If you'd like to leverage BMX with AWSCLI, you can learn more about [how BMX handles credential_process](./credential_process.md).

> :warning: **If you are a member of Arctic Wolf**: Consult internal documentation on Okta & AWS

-----
