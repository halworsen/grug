package grug

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
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
			// Puts together a list of all loaded commands
			Name: "GetCommandList",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				listMsg := "**Loaded commands**\n"
				for _, cmd := range g.Commands {
					listMsg += fmt.Sprintf("%s: %s\n", cmd.Name, strings.Join(cmd.Activators, ", "))
				}
				return listMsg, nil
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
			// super specific but whatever
			// gets the ID of the latest message in range of a given message ID
			Name: "GetLastMediaMessageIDAroundID",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				mID := atostr(args[0])
				messages, err := g.DiscordSession.ChannelMessages(g.CurrentCommand.ChannelID, 100, mID, "", mID)

				var mediaMsg *discordgo.Message
				for _, m := range messages {
					// any attachment or any embed of the correct type
					if len(m.Attachments) > 0 || (len(m.Embeds) > 0 &&
						(m.Embeds[0].Type == discordgo.EmbedTypeImage ||
							m.Embeds[0].Type == discordgo.EmbedTypeGifv ||
							m.Embeds[0].Type == discordgo.EmbedTypeVideo)) {
						mediaMsg = m
						break
					}
				}

				if mediaMsg == nil {
					return nil, errors.New("unable to find any relevant media messages")
				}
				return mediaMsg.ID, err
			},
		},
		{
			// Returns the content of a message given its message ID
			Name: "GetMessageContent",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				mID := atostr(args[0])
				msg, err := g.DiscordSession.ChannelMessage(g.CurrentCommand.ChannelID, mID)
				return msg.Content, err
			},
		},
		{
			// Extracts the URL of the first media content in a message
			Name: "GetMediaURL",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				mID := atostr(args[0])
				msg, err := g.DiscordSession.ChannelMessage(g.CurrentCommand.ChannelID, mID)
				if err != nil {
					return nil, err
				}

				mediaURL := ""
				if len(msg.Embeds) > 0 {
					mediaURL = msg.Embeds[0].URL
				} else if len(msg.Attachments) > 0 {
					mediaURL = msg.Attachments[0].URL
				}

				if mediaURL == "" {
					return nil, errors.New("no embeds in given message")
				}

				return mediaURL, nil
			},
		},
	}...)
}
