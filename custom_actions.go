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
				messages, err := g.DiscordSession.ChannelMessages(g.CurrentCommand.ChannelID, 50, mID, "", mID)

				var mediaMsg *discordgo.Message
				for _, m := range messages {
					if len(m.Embeds) > 0 && (m.Embeds[0].Type == discordgo.EmbedTypeImage ||
						m.Embeds[0].Type == discordgo.EmbedTypeGifv ||
						m.Embeds[0].Type == discordgo.EmbedTypeVideo) {
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
	}...)
}
