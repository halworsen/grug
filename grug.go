package grug

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type GrugSession struct {
	Config         *GrugConfig
	Commands       []Command
	Actions        []Action
	ActivatorMap   map[string]Command
	ActionMap      map[string]Action
	DiscordSession *discordgo.Session
	CurrentCommand *discordgo.MessageCreate
}

func (g *GrugSession) grugMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	r := csv.NewReader(strings.NewReader(m.Content))
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

	// Save as the command being executed for use by any actions
	g.CurrentCommand = m

	userArgs := parts[2:]
	// Execute all steps in the command
	log.Println("[INFO] executing", cmd.Name)
	for step, activator := range cmd.Plan {
		// Recover from any unexpected/unhandled failures
		defer func() {
			r := recover()
			if r != nil {
				g.Log(logError, fmt.Sprint("Panicked out of step ", step, " (action: ", activator.ActionName, "), command execution was forcefully aborted! - ", r))
			}
		}()

		err := g.PerformStep(activator, userArgs)
		if err != nil {
			g.Log(logError, fmt.Sprint("Failed to execute step ", step, " - ", err))
			if activator.HaltOnFail {
				g.Log(logInfo, fmt.Sprint("Halting execution of command ", cmd.Name))
				break
			}
		}
	}

	PurgeArgStore()
}

// Sets up a new Grug session by parsing its master config, loading commands and establishing a Discord session
func (g *GrugSession) New(cfgPath string) {
	g.Log(logInfo, "Loading master config from", cfgPath)
	err := g.LoadMasterConfig(cfgPath)
	if err != nil {
		g.Log(logFatal, "Failed to load master configuration -", err)
	}

	g.Log(logInfo, "Loading commands")
	err = g.LoadCommands()
	if err != nil {
		g.Log(logFatal, "Failed to load commands -", err)
	}

	err = g.ConstructActivatorMap()
	if err != nil {
		g.Log(logFatal, "Failed to construct commands -", err)
	}
	g.Actions = AllActions
	g.ConstructActionMap()

	g.Log(logInfo, "Establishing Discord session")
	session, err := discordgo.New("Bot " + g.Config.Token)
	if err != nil {
		g.Log(logFatal, "Failed to create Discord session -", err)
	}
	g.DiscordSession = session

	session.AddHandler(g.grugMessageHandler)

	// Start listening for commands
	err = session.Open()
	if err != nil {
		g.Log(logFatal, "Failed to open connection -", err)
	}
}

// Closes the Discord session associated with the Grug session
func (g *GrugSession) Close() {
	g.Log(logInfo, "Closing Discord session")
	g.DiscordSession.Close()
}
