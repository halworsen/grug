package discordgrug

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/halworsen/grug"
)

func init() {
	grug.AllActions = append(grug.AllActions, []grug.Action{
		{
			// Returns the command message's ID
			Name: "GetCommandMessageID",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				return curMsg.ID, nil
			},
		},
		{
			// Takes the channel ID of the command message and returns it
			Name: "GetCommandChannelID",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				return curMsg.ChannelID, nil
			},
		},
		{
			// Takes the guild ID of the command message and returns it
			Name: "GetCommandGuildID",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				return curMsg.GuildID, nil
			},
		},
		{
			// Takes the user of the command message and returns it
			Name: "GetCommandUser",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				return curMsg.Member.User, nil
			},
		},
		{
			// Returns the command message
			Name: "GetCommandMessage",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				curMsg := ctx.Value(DISCORDMSG).(*discordgo.MessageCreate)
				return curMsg, nil
			},
		},
	}...)
}
