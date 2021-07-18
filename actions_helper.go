package grug

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aquilax/truncate"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
				asStr, ok := args[1].(string)
				if !ok {
					return nil, fmt.Errorf("unable to interpret %v as string", args[0])
				}
				return fieldAccess(args[0], asStr)
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
			// Joins all args as a string with a given delimiter
			Name: "ConcatenateWith",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				argsAsStr := make([]string, 0)
				delimiter, ok := args[len(args)-1].(string)
				if !ok {
					return nil, fmt.Errorf("unable to use %v as concatenation delimiter", args[len(args)-1])
				}
				for i := 0; i < len(args)-1; i++ {
					if argStr, ok := args[i].(string); ok {
						argsAsStr = append(argsAsStr, argStr)
					}
				}
				return strings.Join(argsAsStr, delimiter), nil
			},
		},
		{
			// Passes arg[0] back directly
			Name: "Passthrough",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return args[0], nil
			},
		},
		{
			// Attempts to unmarshal the 0th arg as JSON
			Name: "UnmarshalJSON",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				asByteSlice, ok := args[0].([]byte)
				if !ok {
					return nil, fmt.Errorf("unable to interpret %v as byte slice", args[0])
				}

				data := make(map[string]interface{})
				json.Unmarshal(asByteSlice, &data)

				return data, nil
			},
		},
		{
			// Attempts to unmarshal the 0th arg as JSON
			Name: "FormatIntCommas",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				var val int
				switch t := args[0].(type) {
				case float32:
					val = int(t)
				case float64:
					val = int(t)
				case int:
					val = t
				default:
					return nil, fmt.Errorf("unable to interpret %v as a number", args[0])
				}

				p := message.NewPrinter(language.English)
				formatted := p.Sprintf("%d", val)

				return formatted, nil
			},
		},
	}...)
}
