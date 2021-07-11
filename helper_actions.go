package grug

import (
	"reflect"
	"strconv"
	"unsafe"

	"github.com/aquilax/truncate"
	"github.com/davecgh/go-spew/spew"
)

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

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
				g.Log(logDebug, len(truncated))
				return truncated, nil
			},
		},
	}...)
}
