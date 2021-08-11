package discordgrug

import (
	"context"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/halworsen/grug"
	"github.com/halworsen/grug/util"
)

func init() {
	grug.AllActions = append(grug.AllActions, []grug.Action{
		{
			// Replies in the same channel as the command was executed in with whatever is in arg 0
			Name: "Reply",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				msg := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				_, err := discordSession.ChannelMessageSend(curMsg.ChannelID, msg)
				return "", err
			},
		},
		{
			// Returns the content of a message given its message ID
			Name: "GetMessageContent",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				mID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				msg, err := discordSession.ChannelMessage(curMsg.ChannelID, mID)
				return msg.Content, err
			},
		},
		{
			// super specific but whatever
			// gets the ID of the latest message in range of a given message ID
			Name: "GetLastMediaMessageIDAroundID",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				mID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				messages, err := discordSession.ChannelMessages(curMsg.ChannelID, 100, mID, "", mID)

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
			// Extracts the URL of the first media content in a message
			Name: "GetMediaURL",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				mID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				msg, err := discordSession.ChannelMessage(curMsg.ChannelID, mID)
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
