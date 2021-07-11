package grug

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			// Compiles a help message for the command named by arg 0
			Name: "GetCommandHelp",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cmdName := atostr(args[0])
				cmd, exists := g.ActivatorMap[cmdName]
				if !exists {
					return "That command doesn't exist :(", nil
				}

				helpText := fmt.Sprintf("%s\n%s\nActivators: %s", cmd.Name, cmd.Description, strings.Join(cmd.Activators, ", "))
				return helpText, nil
			},
		},
		{
			// Replies in the same channel as the command was executed in with whatever is in arg 0
			Name: "Reply",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				msg := atostr(args[0])
				_, err := g.DiscordSession.ChannelMessageSend(g.CurrentCommand.ChannelID, msg)
				return "", err
			},
		},
		{
			// Computes the result of arg 0 + arg 1
			Name: "Plus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return strconv.Itoa(a + b), nil
			},
		},
	}...)
}
