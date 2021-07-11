# grug

Grug is a customizable Discord bot with composable and pluggable commands that are loaded as configurations at runtime.

## Adding commands

Grug commands are sequential steps of *actions*. Once a command is invoked, Grug will sequentially perform the steps according to the command configuration. If an action returns a value, it may be stored on a stack for later use. User arguments to the command can also be supplied as action arguments. For example:

```yaml
name: "Calculator" # Command name
desc: "Perform some simple calculation" # Command description
activators: # List of ways to invoke the command
  - "calc"
steps: # Actions are executed sequentially
  - action: Plus # Add arg 1 and arg 2 together
    args:
      - "!1" # Use the first user supplied argument as an action argument for Plus
      - "!2" # Use the second user supplied argument as an action argument for Plus
    push: true # Push the result of the Plus action onto the result stack
  - action: Reply # Reply in the same channel that the message was sent from
    args:
      - "!1 + !2 = !pop" # !pop pops one action result off the stack
```

Look in [commands](./commands) for more examples.
