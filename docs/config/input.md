# Input Modes

BMX supports controls to modify how user input is requested. This is necessary as BMX can execute in contexts where there is no console input. An example of this would be in the AWSCLI [credential_process](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html) that sources credentials with an external process. This means that BMX can executed by the AWS tooling when no console is open. The alternative input models are built with the intention to ensure workflows like this can continue.

The user configuration `input` has been made available to change the behaviour of input device selection. This is intended to be flexible to address scenarios where a successful input device is difficult to create. The available options are described below:

- `console` tries to use the console first, then falls back on other input methods.
- `applescript` tries to use the console first, and uses AppleScript first when in limited.
- `always_applescript` tries to use AppleScript for input always.

The default input mode is `console`.

## Non-interactive Context

There may be a context where no interaction is possible. In the event that you encounter such a case, you can use `bmx login` to retrieve a session ahead of time in an interactive context. If all needed input are provided to BMX, commands can run in a non-interactive context.