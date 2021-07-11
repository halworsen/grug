# grug

Grug is a customizable Discord bot with composable and pluggable commands that are loaded as configurations at runtime.

## Adding commands

Grug commands are sequential steps of *actions* or *conditionals* called action sequences. Once a command is invoked, Grug will sequentially perform the steps in the action sequence defined by the command configuration. If an action returns a value, it may be named and stored for later use. User arguments to the command can also be supplied as action arguments. For conditionals, two action sequences can be specified to determine behavior. For example:

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
    store: plus_result # Store the result of the Plus action
  - action: Reply # Reply in the same channel that the message was sent from
    args:
      - "!1 + !2 = !plus_result"
  - if: # Conditionally perform one of two action sequences
      condition: int> # The name of the conditional action to use for evaluating the condition
      args: # Operands/arguments to the conditional action
        - "!plus_result"
        - 100
      true: # Action sequence to perform if the condition was true
        - action: Reply
          args:
            - "wow that number was really big"
      false: # Action sequence to perform if the condition was false
        - action: Reply
          args:
            - "that number was kinda small"
```

Look in [commands](./commands) for more examples.
