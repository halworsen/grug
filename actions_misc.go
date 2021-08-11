package grug

import (
	"context"
	"fmt"
	"strings"

	"github.com/halworsen/grug/util"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			// Compiles a help message for the command named by arg 0
			Name: "GetCommandHelp",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cmdName := util.Atostr(args[0])
				cmd, exists := g.ActivatorMap[cmdName]
				if !exists {
					return "That command doesn't exist :(", nil
				}

				helpText := fmt.Sprintf("%s\n%s\nActivators: %s", cmd.Name, cmd.Description, strings.Join(cmd.Activators, ", "))
				return helpText, nil
			},
		},
		{
			// Puts together a list of all loaded commands
			Name: "GetCommandList",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				listMsg := "**Loaded commands**\n"
				for _, cmd := range g.Commands {
					listMsg += fmt.Sprintf("%s: %s\n", cmd.Name, strings.Join(cmd.Activators, ", "))
				}
				return listMsg, nil
			},
		},
	}...)
}
