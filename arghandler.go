package grug

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var argStore map[string]interface{}
var argRegexp *regexp.Regexp
var storedRegexp *regexp.Regexp
var storedNameRegexp *regexp.Regexp

func init() {
	argStore = make(map[string]interface{})
	// ! alone is not valid but golang regex has no lookahead so we'll validate manually
	argRegexp = regexp.MustCompile(`!(-?[0-9]+)?(:)?(-?[0-9]+)?`)
	storedRegexp = regexp.MustCompile(`!([a-zA-Z_]+)`)
	storedNameRegexp = regexp.MustCompile(`^[a-zA-Z_]+$`)
}

// Check if a store name is valid or not
func validateStoredName(name string) bool {
	return storedNameRegexp.Match([]byte(name))
}

// Given an array of the form ["lower" ":" "upper"], returns the lower/upper bounds as ints
func getSliceBounds(slice []string, max int) (int, int) {
	lower, upper := 0, max

	if len(slice) < 3 {
		return lower, upper
	}

	if slice[0] != "" {
		lowerIn, err := strconv.Atoi(slice[0])
		if lowerIn < 0 {
			lowerIn += max
		}
		if err == nil {
			lower = lowerIn
		}
	}

	if slice[1] != "" {
		upperIn, err := strconv.Atoi(slice[2])
		if upperIn < 0 {
			upperIn += max
		}
		if err == nil {
			upper = upperIn
		}
	}

	return lower, upper
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

		// special case: user arg template alone can pass the slice preserving the split
		matches = argRegexp.FindStringSubmatch(arg)
		if len(matches) > 0 && matches[0] == arg && arg != "!" {
			// its a slice
			if matches[2] == ":" {
				// get the bounds and add all of the args according to the slice bounds
				lower, upper := getSliceBounds(matches[1:], len(usrArgs))
				for i := lower; i < upper; i++ {
					finalArgs = append(finalArgs, usrArgs[i])
				}
				continue
			}

			argIdx, _ := strconv.Atoi(matches[1])
			if argIdx < 0 {
				argIdx += len(usrArgs)
			} else {
				argIdx -= 1
			}

			if argIdx < 0 || argIdx > len(usrArgs)-1 {
				panic(fmt.Sprint("arg index out of bounds for configured user arg ", matches[1]))
			}

			finalArgs = append(finalArgs, usrArgs[argIdx])
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

		// Try to find arguments of the form !1, !2:, !1:5, etc. and replace them with their respective user args
		newArg = argRegexp.ReplaceAllStringFunc(newArg, func(s string) string {
			submatches := argRegexp.FindStringSubmatch(s)
			if submatches[0] == "!" {
				return s
			}

			if submatches[2] == ":" {
				lower, upper := getSliceBounds(submatches[1:], len(usrArgs))
				return strings.Join(usrArgs[lower:upper], " ")
			}

			argIdx, _ := strconv.Atoi(submatches[1])
			if argIdx < 0 {
				argIdx += len(usrArgs)
			} else {
				argIdx -= 1
			}

			if argIdx < 0 || argIdx > len(usrArgs)-1 {
				panic(fmt.Sprint("arg index out of bounds for configured user arg ", submatches[1]))
			}

			return usrArgs[argIdx]
		})

		finalArgs = append(finalArgs, newArg)
	}
	return finalArgs, nil
}
