package grug

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aquilax/truncate"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	// config for pretty prints
	spew.Config.MaxDepth = 2
	spew.Config.Indent = "  "

	AllActions = append(AllActions, []Action{
		{
			// Returns the command message's ID
			Name: "GetCommandMessageID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.ID, nil
			},
		},
		{
			// Takes the channel ID of the command message and returns it
			Name: "GetCommandChannelID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.ChannelID, nil
			},
		},
		{
			// Takes the guild ID of the command message and returns it
			Name: "GetCommandGuildID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.GuildID, nil
			},
		},
		{
			// Takes the user of the command message and returns it
			Name: "GetCommandUser",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand.Member.User, nil
			},
		},
		{
			// Returns the command message
			Name: "GetCommandMessage",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return g.CurrentCommand, nil
			},
		},
		{
			// Pretty formats the 0th argument and returns it
			Name: "PrettyFormat",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return spew.Sdump(args[0]), nil
			},
		},
		{
			// Truncates the 0th argument to the amount of characters in arg 1
			Name: "Truncate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				val := atostr(args[0])
				truncTo, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}

				truncated := truncate.Truncate(val, truncTo, "...", truncate.PositionEnd)
				return truncated, nil
			},
		},
		{
			// Performs a string of accesses
			// E.g. FieldA.0.B would access args[0].FieldA[0].B
			Name: "FieldAccess",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				if args[1] == "." {
					return args[0], nil
				}
				indexPlan := strings.Split(atostr(args[1]), ".")

				var value interface{}
				for n, field := range indexPlan {
					if n == 0 {
						value = args[0]
					}

					idx, err := strconv.Atoi(field)
					if err == nil {
						valAsSlice, ok := value.([]interface{})
						if !ok {
							return nil, fmt.Errorf("%v is not accessible by index", value)
						}
						value = valAsSlice[idx]
					} else {
						valAsMap, ok := value.(map[string]interface{})
						if !ok {
							return nil, fmt.Errorf("%v is not accessible by name %s", value, field)
						}
						value = valAsMap[field]
					}
				}

				return value, nil
			},
		},
		{
			// Converts the input to a JSON string
			Name: "ToJSON",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				json, err := json.MarshalIndent(args[0], "", "  ")
				if err != nil {
					return nil, err
				}
				return string(json), nil
			},
		},
		{
			// Concatenates all args as strings
			Name: "Concatenate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				argsAsStr := make([]string, 0)
				for _, arg := range args {
					if argStr, ok := arg.(string); ok {
						argsAsStr = append(argsAsStr, argStr)
					}
				}
				return strings.Join(argsAsStr, ""), nil
			},
		},
		{
			// Passes arg[0] back directly
			Name: "Passthrough",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return args[0], nil
			},
		},
	}...)
}
