# Input Modes

Bmx supports the user configuration parameter `input` to control behaviour around requesting user input. There exist cases where Bmx is run with no guarantee of terminal input. An example of this would be in the AWSCLI (credential_process)[https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html] that sources credentials with an external process. To accomodate edge cases, bmx supports alternative input behaviours to ensure workflows continue without interruption.

The configuration `input` has been made available to change the behaviour of input device selection. This is intended to address scenarios where the determination of an available tty is incorrect, and the interface option needs to be overwritten.

The available options are [`console`, `applescript`, `always_applescript`], with the options described as such:

- `console` tries to use the console first, then falls back on other input methods.
- `applescript` tries to use the console first, and uses AppleScript first when in limited.
- `always_applescript` tries to use AppleScript for input always.

The default input mode is `console`.