package grug

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

var argStack *stack.Stack
var argRegexp *regexp.Regexp
var popRegexp *regexp.Regexp

func init() {
	argStack = stack.New()
	argRegexp = regexp.MustCompile("!([0-9]+)")
	popRegexp = regexp.MustCompile("!pop")
}

// Pushes the given result onto the arg stack
func PushStepResult(result interface{}) {
	argStack.Push(result)
}

// Parses special runtime dependent arguments from configuration-time defined arguments and returns a slice of usable args
func ParseArgs(cfgArgs []interface{}, usrArgs []string) ([]interface{}, error) {
	finalArgs := make([]interface{}, 0)
	for _, arg := range cfgArgs {
		arg := atostr(arg)

		// special case: !pop *alone* can pass arbitrary values
		if arg == "!pop" {
			finalArgs = append(finalArgs, argStack.Pop())
			continue
		}

		// !pop is a special arg that uses the topmost value from the arg stack as an argument
		newArg := popRegexp.ReplaceAllStringFunc(arg, func(s string) string {
			return atostr(argStack.Pop())
		})

		// Try to find arguments of the form !1, !2, etc. and replace them with their respective user args
		newArg = argRegexp.ReplaceAllStringFunc(newArg, func(s string) string {
			submatches := argRegexp.FindStringSubmatch(s)
			argIdx, _ := strconv.Atoi(submatches[1])

			argIdx -= 1
			if argIdx < 0 || argIdx > len(usrArgs)-1 {
				panic(fmt.Sprint("arg index out of bounds for configured user arg ", submatches[1]))
			}

			return usrArgs[argIdx]
		})

		finalArgs = append(finalArgs, newArg)
	}
	return finalArgs, nil
}
