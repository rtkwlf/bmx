# Setting up BMX with credential_process

BMX supports acting as an external process for sourcing credentials for the AWSCLI. This feature is known as [credential_process](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html). 

When BMX is running as a `credential_process` it may not always be able to request password input from the console. In these cases, BMX supports AppleScript fallbacks. You can review how [bmx handles input](./config/input.md) in the documentation.

## Getting Started

You'll need to open your AWS configuration file. If you already have entries, it's recommend to make a backup before fiddling with the entries.

```bash
cat ~/.aws/config
```

When you are ready, you can create a profile named `bmx` (or anything for that matter) and set the `credential_process` parameter to be equal to `bmx credential-process`. When running without parameters, bmx will look to the configuration file (`~/.bmx/config`) to see if any of the properties are set. If they are not, then BMX will request input from the user.

```ini
# ~/.aws/config
[profile bmx]
credential_process = bmx credential-process
output = json

[profile default]
role_arn = arn:aws:iam::123456789011:role/ACMESSO-Developer
source_profile = bmx
output = json
```

When running the command `aws sts get-caller-identity`, this will use the source profile `bmx` to retrieve credentials from the BMX command line tool. It will then use the `default` profile, which assumes the role of `ACMESSO-Developer`. The result will be 

You can specify additional inputs to `bmx credential-process` if you'd like to avoid being prompted for anything but a password. If you login ahead of time using `bmx login`, you can avoid being prompted at all.