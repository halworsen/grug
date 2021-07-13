# grug [![Go Report Card](https://goreportcard.com/badge/github.com/halworsen/grug)](https://goreportcard.com/report/github.com/halworsen/grug)

___

Grug is a customizable Discord bot with composable and pluggable commands that are loaded as configurations at runtime.

## Grug commands

Grug commands consist of 4 parts:

* A name - A descriptive name of the command
* A description - A description of the command and how to use it
* Activators - A list of ways to invoke the command
* A command plan - A list of *actions* or *conditionals*

When a command is invoked using one of its activators, Grug executes the command plan sequentially.

### Templating

Grug features simple templating to access stored values and user arguments. All templated values start with `!` followed by either a name (for stored values) or a slice (for user arguments).

| Example template | User args | Store | Result |
|------------------|-----------|-------|--------|
| `Your first arg was !1`| hello world | | `Your first arg was hello` |
| `Your first arg was !1`| "hello world" | | `Your first arg was hello world` |
| `Your 2nd and 3rd args were: !2:4` | a b c d e | | `Your 2nd and 3rd args were: b c` |
| `All your args were: !:` | foo bar baz | | `All your args were: foo bar baz` |
| `Your last arg was: !-1` | hi there | | `Your last arg was: there` |
| `!food is !2` | foo good | food: "cake" | `cake is good` |
| `!:` | a b c d | | "a", "b", "c", "d" are passed as arguments |
| `!food` | | food: `[[1 2] [3 4]]` | `[[1 2] [3 4]]` is passed as an argument |

### Actions

Actions are the base unit of Grug. They are (preferrably simple) tasks that may or may not take arguments, and may or may not produce output. Action output may be stored in named fields, and a failure plan may be specified as a plan to be executed when the action fails.

Actions are implemented in code as reusable components for composing commands.

```yaml
action: Plus # Action for adding the two arguments together
args:
  - "!1" # Use the first user supplied argument as an action argument for Plus
  - "!2" # Use the second user supplied argument as an action argument for Plus
store: plus_result # Store the result of the Plus action in the field "plus_result"
```

### Conditionals

Conditionals are fancy wrappers for an action execution where the result determines which plan should be executed.

For a list of available conditionals see [conditional_actions.go](./conditional_actions.go)

```yaml
if:
  condition: int> # int> is just a normal action that returns a bool
  args:
    - !1
    - 2
  true: # The plan to execute if user arg #1 > 2
    ...
  false: # The plan to execute if user arg #1 <= 2
    ...
```

### Failure handling

If a plan fails to execute normally, indicated by the action returning an error from its `Exec` implementation, and a failure plan is configured, the failure plan is executed in its entirety before the next action in the plan is executed.

The `haltOnFailure` option may also be set to true to abort command execution if the action fails. If a failure plan is specified, it will be executed before the command halts.

```yaml
- action: GetCommandMessageID
  store: msgID
- action: GetLastMediaMessageIDAroundID
  args:
    - "!msgID"
  store: mediaMsgID
  haltOnFailure: true # Halt command execution if this action fails
  failurePlan: # The plan to execute if this action fails
    - action: Reply
      args:
        - "Couldn't find any messages with media :/"
...
```

### Example

For more examples, see [example/commands](./example/commands).

```yaml
name: "Calculator" # Command name
desc: "Perform some simple calculation" # Command description
activators: # List of ways to invoke the command
  - "calc"
plan: # Actions are executed sequentially according to the plan
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
