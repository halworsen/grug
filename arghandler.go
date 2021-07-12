package grug

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var argStore map[string]interface{}
var argRegexp *regexp.Regexp
var storedRegexp *regexp.Regexp
var storedNameRegexp *regexp.Regexp

func init() {
	argStore = make(map[string]interface{})
	argRegexp = regexp.MustCompile("!([0-9]+)")
	storedRegexp = regexp.MustCompile("!([a-zA-Z_]+)")
	storedNameRegexp = regexp.MustCompile("^[a-zA-Z_]+$")
}

// Check if a store name is valid or not
func validateStoredName(name string) bool {
	return storedNameRegexp.Match([]byte(name))
}

// StoreArg stores the given value in the given field in an arg store
func StoreArg(name string, val interface{}) error {
	if !validateStoredName(name) {
		return errors.New(fmt.Sprint(name, " is not a valid store name"))
	}

	argStore[name] = val
	return nil
}

// PurgeArgStore clears the arg store by recreating it
func PurgeArgStore() {
	argStore = make(map[string]interface{})
}

// ParseArgs parses templated values and inserts their respective actual values
func ParseArgs(cfgArgs []interface{}, usrArgs []string) ([]interface{}, error) {
	finalArgs := make([]interface{}, 0)
	for _, arg := range cfgArgs {
		arg := atostr(arg)

		// special case: !stored_name alone can pass arbitrary values
		matches := storedRegexp.FindStringSubmatch(arg)
		if len(matches) > 0 && matches[0] == arg {
			val, ok := argStore[matches[1]]
			if ok {
				finalArgs = append(finalArgs, val)
				continue
			}
		}

		// !stored_name retrieves a stored name from the arg store
		newArg := storedRegexp.ReplaceAllStringFunc(arg, func(s string) string {
			submatches := storedRegexp.FindStringSubmatch(s)
			val, ok := argStore[submatches[1]]
			if ok {
				return atostr(val)
			}
			return s
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
