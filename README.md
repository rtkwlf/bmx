> This repository is now deprecated, and no longer maintained. It is recommended to look at the upstream project: (https://github.com/Brightspace/bmx/).


# BMX

BMX grants you API access to your AWS accounts, based on Okta credentials that you already own.  
It uses your Okta identity to create short-term AWS STS tokens, as an alternative to long-term IAM access keys.
BMX manages your STS tokens with the following commands:

1. `bmx print` writes your short-term tokens to `stdout` as AWS environment variables.  You can execute `bmx print`'s output to make the environment variables available to your shell.
1. `bmx write` writes your short-term tokens to `~/.aws/credentials`.

BMX prints detailed usage information when you run `bmx -h` or `bmx <cmd> -h`.

BMX was developed by D2L ([Brightspace/bmx](https://github.com/Brightspace/bmx/)), and modifications have been made to the project by Arctic Wolf.

## Features

1. BMX is multi-platform: it runs on Linux, Windows, and Mac.
2. BMX maintains your Okta session for 12 hours: you enter your Okta password once a day, and BMX takes care of the rest.
3. Project scoped configurations
4. BMX supports Web and SMS MFA.

## Installation

Available versions of BMX are available on the [releases](https://github.com/rtkwlf/bmx/releases) page. 

## Getting Started

To authenticate and obtain a session via the command line, run the following:

```bash
bmx login
```

This will prompt you for your Okta organization and credentials. When you have successfully connected, you can run the following to get a set of IAM STS credentials for use with the AWS API:

```bash
bmx print
```

The command will print a series of environment set commands, that can be used to set the environment variables of the current shell session:

```bash
export AWS_SESSION_TOKEN=...
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...

# Run AWSCLI using environment variables for credentials
aws sts get-caller-identity
```

If you'd like to learn about the ways BMX assists with authenticating to AWS, you can review in the [getting started](./docs/) documentation.

## Versioning

BMX is maintained under the [Semantic Versioning guidelines](http://semver.org/).

### Getting Involved

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.
