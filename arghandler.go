package grug

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var argStore map[string]interface{}
var argRegexp *regexp.Regexp
var storedRegexp *regexp.Regexp
var storedNameRegexp *regexp.Regexp

func init() {
	argStore = make(map[string]interface{})
	// ! alone is not valid but golang regex has no lookahead so we'll validate manually
	argRegexp = regexp.MustCompile(`!(-?[0-9]+)?(:)?(-?[0-9]+)?(?:\.\.\.)?`)
	storedRegexp = regexp.MustCompile(`!([a-zA-Z_]+)(?:\.\.\.)?`)
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
		} else {
			lowerIn -= 1
		}
		if err == nil {
			lower = lowerIn
		}
	}

	if slice[1] != "" {
		upperIn, err := strconv.Atoi(slice[2])
		if upperIn < 0 {
			upperIn += max
		} else {
			upperIn -= 1
		}
		if err == nil {
			upper = upperIn
		}
	}

	return lower, upper
}

// note: alters the valueMap directly
// little bit of "replace" abuse. we really just need to run the replace function for every match
func populateTemplateValueMap(valueMap *map[string]interface{}, arg string, usrArgs []string) {
	// !stored_name retrieves a stored name from the arg store
	storedRegexp.ReplaceAllStringFunc(arg, func(s string) string {
		if _, ok := (*valueMap)[s]; !ok {
			// 1st char is the !
			(*valueMap)[s] = argStore[s[1:]]
		}
		return s
	})

	// Try to find arguments of the form !1, !2:, !1:5, etc.
	argRegexp.ReplaceAllStringFunc(arg, func(s string) string {
		// the regex can match ! alone so we filter that out here
		if s == "!" {
			return s
		}

		if _, ok := (*valueMap)[s]; !ok {
			submatches := argRegexp.FindStringSubmatch(s)
			// if it's a slice, add the slice
			if submatches[2] == ":" {
				lower, upper := getSliceBounds(submatches[1:], len(usrArgs))
				(*valueMap)[s] = usrArgs[lower:upper]
				return s
			}

			// else, add a specific index
			argIdx, _ := strconv.Atoi(submatches[1])
			if argIdx < 0 {
				argIdx += len(usrArgs)
			} else {
				argIdx -= 1
			}

			if argIdx < 0 || argIdx > len(usrArgs)-1 {
				panic(fmt.Sprint("arg index out of bounds for configured user arg ", submatches[1]))
			}

			(*valueMap)[s] = usrArgs[argIdx]
		}

		return s
	})
}

// appends a value but "opens up" slices and arrays and adds every element instead of the slice itself
func appendExpandSlices(slice []interface{}, val interface{}) []interface{} {
	reflectionVal := reflect.ValueOf(val)
	if reflectionVal.Kind() == reflect.Slice || reflectionVal.Kind() == reflect.Array {
		for i := 0; i < reflectionVal.Len(); i++ {
			slice = append(slice, reflectionVal.Index(i).Interface())
		}
	} else {
		slice = append(slice, val)
	}
	return slice
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
	templateValueMap := make(map[string]interface{})
	for _, arg := range cfgArgs {
		// if it's not a string there's no templating to do
		argAsStr, ok := arg.(string)
		if !ok {
			finalArgs = append(finalArgs, arg)
			continue
		}

		// update the templating value map
		populateTemplateValueMap(&templateValueMap, argAsStr, usrArgs)

		// directly passing args when the template appears alone
		if match := storedRegexp.FindString(argAsStr); match == argAsStr {
			if match[len(match)-3:] == "..." {
				finalArgs = appendExpandSlices(finalArgs, templateValueMap[argAsStr])
			}
			finalArgs = append(finalArgs, templateValueMap[argAsStr])
			continue
		}
		if match := argRegexp.FindString(argAsStr); argAsStr != "!" && match == argAsStr {
			if match[len(match)-3:] == "..." {
				finalArgs = appendExpandSlices(finalArgs, templateValueMap[argAsStr])
			}
			finalArgs = append(finalArgs, templateValueMap[argAsStr])
			continue
		}

		argAsStr = storedRegexp.ReplaceAllStringFunc(argAsStr, func(s string) string {
			return atostr(templateValueMap[s])
		})

		argAsStr = argRegexp.ReplaceAllStringFunc(argAsStr, func(s string) string {
			if s == "!" {
				return s
			}
			return atostr(templateValueMap[s])
		})

		finalArgs = append(finalArgs, argAsStr)
	}
	return finalArgs, nil
}
