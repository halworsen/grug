package grug

import (
	"context"
	"encoding/csv"
	"fmt"
	"strings"
)

// GrugSession holds all information relevant to the operation of Grug
type GrugSession struct {
	Config       *GrugConfig            // The Grug master config
	Commands     []Command              // All loaded commands
	Actions      []Action               // All actions used by Grug
	ActivatorMap map[string]Command     // Uniquely maps command name to Command struct
	ActionMap    map[string]Action      // Uniquely maps action name to Action struct
	ArgStore     map[string]interface{} // Args stored after action executions go in here
}

// HandleMessage invokes commands based on the message. msgCtx is ultimately passed to all actions that are executed
func (g *GrugSession) HandleMessage(msgCtx context.Context, msg string) {
	r := csv.NewReader(strings.NewReader(msg))
	r.Comma = ' '
	parts, err := r.Read()
	if len(parts) == 0 || err != nil {
		return
	}

	// empty invoker is allowed and means commands are invoked directly
	if g.Config.Invoker == "" {
		// lil hack, we prepend an empty string to match the (empty) bot invoker
		paddedParts := make([]string, 0)
		paddedParts = append(paddedParts, "")
		parts = append(paddedParts, parts...)
	}

	// Check if the bot is being invoked
	if parts[0] != g.Config.Invoker {
		return
	}

	// Check which (if any valid) command is being invoked
	cmd, present := g.ActivatorMap[parts[1]]
	if !present {
		return
	}

	userArgs := parts[2:]
	// Execute all steps in the command
	g.Log(logInfo, "Executing", cmd.Name)
	for step, activator := range cmd.Plan {
		// Recover from any unexpected/unhandled failures
		if !g.Config.HardError {
			defer func() {
				r := recover()
				if r != nil {
					g.Log(logError, fmt.Sprint("Panicked out of step ", step, " (action: ", activator.ActionName, "), command execution was forcefully aborted! - ", r))
				}
			}()
		}

		err := g.performAction(msgCtx, activator, userArgs)
		if err != nil {
			g.Log(logError, fmt.Sprint("Failed to execute step ", step, " - ", err))
			if activator.HaltOnFail {
				g.Log(logInfo, fmt.Sprint("Halting execution of command ", cmd.Name))
				break
			}
		}
	}

	g.purgeArgStore()
}

// New sets up a Grug session by parsing its master config, loading commands and establishing a Discord session
func (g *GrugSession) New(cfgPath string) {
	g.ArgStore = make(map[string]interface{})

	g.Log(logInfo, "Loading Grug master config from", cfgPath)
	err := g.loadMasterConfig(cfgPath)
	if err != nil {
		g.Log(logFatal, "Failed to load master configuration -", err)
	}

	g.Log(logInfo, "Loading Grug commands")
	err = g.loadCommands()
	if err != nil {
		g.Log(logFatal, "Failed to load commands -", err)
	}

	err = g.constructActivatorMap()
	if err != nil {
		g.Log(logFatal, "Failed to construct commands -", err)
	}
	g.Actions = AllActions
	g.constructActionMap()

	g.Log(logInfo, "Grug session ready")
}
