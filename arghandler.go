package grug

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/halworsen/grug/util"
)

var argRegexp *regexp.Regexp
var storedRegexp *regexp.Regexp
var storedNameRegexp *regexp.Regexp

func init() {
	// ! alone is not valid but golang regex has no lookahead so we'll validate manually
	argRegexp = regexp.MustCompile(`!(-?[0-9]+)?(:)?(-?[0-9]+)?(?:\.\.\.)?`)
	storedRegexp = regexp.MustCompile(`!([a-zA-Z_]+)(?:\.\.\.)?`)
	storedNameRegexp = regexp.MustCompile(`^[a-zA-Z_]+$`)
}

// Check if a store name is valid or not
func validateStoredName(name string) bool {
	return storedNameRegexp.Match([]byte(name))
}

// note: alters the valueMap directly
// little bit of "replace" abuse. we really just need to run the replace function for every match
func (g *GrugSession) populateTemplateValueMap(valueMap *map[string]interface{}, arg string, usrArgs []string) {
	// !stored_name retrieves a stored name from the arg store
	storedRegexp.ReplaceAllStringFunc(arg, func(s string) string {
		if _, ok := (*valueMap)[s]; !ok {
			// 1st char is the !
			(*valueMap)[s] = g.ArgStore[s[1:]]
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
				lower, upper := util.GetSliceBounds(submatches[1:], len(usrArgs))
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

// StoreArg stores the given value in the given field in an arg store
func (g *GrugSession) storeArg(name string, val interface{}) error {
	if !validateStoredName(name) {
		return errors.New(fmt.Sprint(name, " is not a valid store name"))
	}

	g.ArgStore[name] = val
	return nil
}

// PurgeArgStore clears the arg store by recreating it
func (g *GrugSession) purgeArgStore() {
	g.ArgStore = make(map[string]interface{})
}

// ParseArgs parses templated values and inserts their respective actual values
func (g *GrugSession) parseArgs(cfgArgs []interface{}, usrArgs []string) ([]interface{}, error) {
	finalArgs := make([]interface{}, 0)
	templateValueMap := make(map[string]interface{})
	for _, arg := range cfgArgs {
		// if it's not a string there's no templating to do
		argAsStr, ok := arg.(string)
		if !ok {
			finalArgs = append(finalArgs, arg)
			continue
		}

		// update the templating value map so each template maps to its value
		g.populateTemplateValueMap(&templateValueMap, argAsStr, usrArgs)

		// directly pass args when the template appears alone
		argAdder := func(match string) {
			if len(match) > 3 && match[len(match)-3:] == "..." {
				util.FMapSlice(templateValueMap[argAsStr], func(arg interface{}) {
					finalArgs = append(finalArgs, arg)
				})
			} else {
				finalArgs = append(finalArgs, templateValueMap[argAsStr])
			}
		}

		if match := storedRegexp.FindString(argAsStr); match == argAsStr {
			argAdder(match)
			continue
		}
		if match := argRegexp.FindString(argAsStr); argAsStr != "!" && match == argAsStr {
			argAdder(match)
			continue
		}

		// Replaces a string match s with the string representation of that arg
		// If suffixed with "...", s is assumed to be the template for a slice, and every slice element's string representation is joined using " "
		stringReplacer := func(s string) string {
			if len(s) > 3 && s[len(s)-3:] == "..." {
				asStrSlice := make([]string, 0)
				util.FMapSlice(templateValueMap[s], func(arg interface{}) {
					asStrSlice = append(asStrSlice, arg.(string))
				})
				return strings.Join(asStrSlice, " ")
			}
			return util.Atostr(templateValueMap[s])
		}

		argAsStr = storedRegexp.ReplaceAllStringFunc(argAsStr, stringReplacer)
		argAsStr = argRegexp.ReplaceAllStringFunc(argAsStr, func(s string) string {
			if s == "!" {
				return s
			}
			return stringReplacer(s)
		})

		finalArgs = append(finalArgs, argAsStr)
	}
	return finalArgs, nil
}
