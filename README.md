# BMX

BMX grants you API access to your AWS accounts, based on Okta credentials that you already own.  
It uses your Okta identity to create short-term AWS STS tokens, as an alternative to long-term IAM access keys.
BMX manages your STS tokens with the following commands:

1. `bmx print` writes your short-term tokens to `stdout` as AWS environment variables.  You can execute `bmx print`'s output to make the environment variables available to your shell.
1. `bmx write` writes your short-term tokens to `~/.aws/credentials`.

BMX prints detailed usage information when you run `bmx -h` or `bmx <cmd> -h`.

BMX was developed by D2L ([Brightspace/bmx](https://github.com/Brightspace/bmx/)), and modifications have been made to the project by Arctic Wolf.

## Installation

Available versions of BMX are available on the [releases](https://github.com/rtkwlf/bmx/releases) page. 

## Features

1. BMX is multi-platform: it runs on Linux, Windows, and Mac.
1. BMX maintains your Okta session for 12 hours: you enter your Okta password once a day, and BMX takes care of the rest.
1. Project scoped configurations
1. BMX supports Web and SMS MFA.

## Versioning

BMX is maintained under the [Semantic Versioning guidelines](http://semver.org/).

### Getting Involved

BMX has [issues](https://github.com/rtkwlf/bmx/issues).

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.
