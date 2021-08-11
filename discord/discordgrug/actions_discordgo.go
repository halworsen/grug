package discordgrug

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/halworsen/grug"
	"github.com/halworsen/grug/util"
)

func init() {
	grug.AllActions = append(grug.AllActions, []grug.Action{
		{
			Name: "ChannelDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channel, err := discordSession.ChannelDelete(cID)
				return channel, err
			},
		},
		{
			Name: "ChannelEdit",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, newName := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channel, err := discordSession.ChannelEdit(cID, newName)
				return channel, err
			},
		},
		{
			Name: "ChannelEditComplex",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				var cEdit discordgo.ChannelEdit
				json.Unmarshal([]byte(util.Atostr(args[1])), &cEdit)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channel, err := discordSession.ChannelEditComplex(cID, &cEdit)
				return channel, err
			},
		},
		// ChannelFileSend not supported
		// ChannelFileSendWithMessage not supported
		{
			Name: "ChannelInviteCreate",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				var cInvite discordgo.Invite
				json.Unmarshal([]byte(util.Atostr(args[1])), &cInvite)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invite, err := discordSession.ChannelInviteCreate(cID, cInvite)
				return invite, err
			},
		},
		{
			Name: "ChannelInvites",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invites, err := discordSession.ChannelInvites(cID)
				return invites, err
			},
		},
		{
			Name: "ChannelMessage",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessage(cID, mID)
				return msg, err
			},
		},
		// ChannelMessageAck not supported
		{
			Name: "ChannelMessageCrosspost",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageCrosspost(cID, mID)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.ChannelMessageDelete(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessageEdit",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID, content := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageEdit(cID, mID, content)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageEditComplex",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				var msgEdit discordgo.MessageEdit
				json.Unmarshal([]byte(util.Atostr(args[0])), &msgEdit)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageEditComplex(&msgEdit)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageEditEmbed",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				var msgEmbed discordgo.MessageEmbed
				json.Unmarshal([]byte(util.Atostr(args[2])), &msgEmbed)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageEditEmbed(cID, mID, &msgEmbed)
				return msg, err
			},
		},
		{
			Name: "ChannelMessagePin",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.ChannelMessagePin(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessageSend",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, msgContent := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageSend(cID, msgContent)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendComplex",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				var msgSend discordgo.MessageSend
				json.Unmarshal([]byte(util.Atostr(args[1])), &msgSend)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageSendComplex(cID, &msgSend)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendEmbed",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				var msgEmbed discordgo.MessageEmbed
				json.Unmarshal([]byte(util.Atostr(args[1])), &msgEmbed)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageSendEmbed(cID, &msgEmbed)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendReply",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, msgContent := util.Atostr(args[0]), util.Atostr(args[1])
				var msgRef discordgo.MessageReference
				json.Unmarshal([]byte(util.Atostr(args[1])), &msgRef)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageSendReply(cID, msgContent, &msgRef)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendTTS",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, msgContent := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.ChannelMessageSendTTS(cID, msgContent)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageUnpin",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.ChannelMessageUnpin(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessages",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, beforeID, afterID, aroundID := util.Atostr(args[0]), util.Atostr(args[2]), util.Atostr(args[3]), util.Atostr(args[4])
				limit, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msgs, err := discordSession.ChannelMessages(cID, limit, beforeID, afterID, aroundID)
				return msgs, err
			},
		},
		// ChannelMessagesBulkDelete not supported
		{
			Name: "ChannelMessagesPinned",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msgs, err := discordSession.ChannelMessagesPinned(cID)
				return msgs, err
			},
		},
		// ChannelNewsFollow not supported
		// ChannelPermissionDelete not supported
		// ChannelPermissionSet not supported
		{
			Name: "ChannelTyping",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.ChannelTyping(cID)
				return nil, err
			},
		},
		// ChannelVoiceJoin not supported
		// ChannelVoiceJoinManual not supported
		{
			Name: "ChannelWebhooks",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhooks, err := discordSession.ChannelWebhooks(cID)
				return webhooks, err
			},
		},
		{
			Name: "ChannelWebhooks",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhooks, err := discordSession.ChannelWebhooks(cID)
				return webhooks, err
			},
		},
		{
			Name: "GuildInfo",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				guild, err := discordSession.Guild(gID)
				return guild, err
			},
		},
		{
			Name: "GuildAuditLog",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uId, beforeID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				actionType, err := strconv.Atoi(util.Atostr(args[3]))
				if err != nil {
					return nil, err
				}
				limit, err := strconv.Atoi(util.Atostr(args[4]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				auditLog, err := discordSession.GuildAuditLog(gID, uId, beforeID, actionType, limit)
				return auditLog, err
			},
		},
		{
			Name: "GuildBan",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uId := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				ban, err := discordSession.GuildBan(gID, uId)
				return ban, err
			},
		},
		{
			Name: "GuildBanCreate",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uId := util.Atostr(args[0]), util.Atostr(args[1])
				days, err := strconv.Atoi(util.Atostr(args[2]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.GuildBanCreate(gID, uId, days)
				return nil, err
			},
		},
		{
			Name: "GuildBanCreateWithReason",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uId, reason := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				days, err := strconv.Atoi(util.Atostr(args[3]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.GuildBanCreateWithReason(gID, uId, reason, days)
				return nil, err
			},
		},
		{
			Name: "GuildBanDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uId := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildBanDelete(gID, uId)
				return nil, err
			},
		},
		{
			Name: "GuildBans",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				bans, err := discordSession.GuildBans(gID)
				return bans, err
			},
		},
		// GuildChannelCreate not supported
		{
			Name: "GuildChannelCreateComplex",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				var channelCreateData discordgo.GuildChannelCreateData
				json.Unmarshal([]byte(util.Atostr(args[1])), &channelCreateData)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channel, err := discordSession.GuildChannelCreateComplex(gID, channelCreateData)
				return channel, err
			},
		},
		{
			Name: "GuildChannels",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channels, err := discordSession.GuildChannels(gID)
				return channels, err
			},
		},
		// GuildChannelsReorder not supported
		// GuildCreate not supported
		// GuildDelete not supported
		// GuildEdit not supported
		// GuildEmbed not supported
		// GuildEmbedEdit not supported
		// GuildEmojiCreate not supported
		{
			Name: "GuildEmojiDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, eID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildEmojiDelete(gID, eID)
				return nil, err
			},
		},
		// GuildEmojiEdit not supported
		{
			Name: "GuildEmojis",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				emojis, err := discordSession.GuildEmojis(gID)
				return emojis, err
			},
		},
		{
			Name: "GuildIcon",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				icon, err := discordSession.GuildIcon(gID)
				return icon, err
			},
		},
		// GuildIntegrationCreate not supported
		// GuildIntegrationDelete not supported
		// GuildIntegrationEdit not supported
		// GuildIntegrationSync not supported
		{
			Name: "GuildIntegrations",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				integrations, err := discordSession.GuildIntegrations(gID)
				return integrations, err
			},
		},
		{
			Name: "GuildInvites",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invites, err := discordSession.GuildInvites(gID)
				return invites, err
			},
		},
		{
			Name: "GuildMember",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				member, err := discordSession.GuildMember(gID, uID)
				return member, err
			},
		},
		// GuildMemberAdd not supported
		{
			Name: "GuildMemberDeafen",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID := util.Atostr(args[0]), util.Atostr(args[1])
				deaf, err := strconv.ParseBool(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.GuildMemberDeafen(gID, uID, deaf)
				return nil, err
			},
		},
		{
			Name: "GuildMemberDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberDelete(gID, uID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberDeleteWithReason",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID, reason := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberDeleteWithReason(gID, uID, reason)
				return nil, err
			},
		},
		// GuildMemberEdit not supported
		{
			Name: "GuildMemberMove",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID, cID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberMove(gID, uID, &cID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberMute",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID := util.Atostr(args[0]), util.Atostr(args[1])
				mute, err := strconv.ParseBool(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.GuildMemberMute(gID, uID, mute)
				return nil, err
			},
		},
		{
			Name: "GuildMemberNickname",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID, nick := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberNickname(gID, uID, nick)
				return nil, err
			},
		},
		{
			Name: "GuildMemberRoleAdd",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID, rID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberRoleAdd(gID, uID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberRoleRemove",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, uID, rID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildMemberRoleRemove(gID, uID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildMembers",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, after := util.Atostr(args[0]), util.Atostr(args[1])
				limit, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				members, err := discordSession.GuildMembers(gID, after, limit)
				return members, err
			},
		},
		{
			Name: "GuildPrune",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				days, err := strconv.ParseUint(util.Atostr(args[1]), 10, 32)
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				count, err := discordSession.GuildPrune(gID, uint32(days))
				return count, err
			},
		},
		{
			Name: "GuildPruneCount",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				days, err := strconv.ParseUint(util.Atostr(args[1]), 10, 32)
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				count, err := discordSession.GuildPruneCount(gID, uint32(days))
				return count, err
			},
		},
		{
			Name: "GuildRoleCreate",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				role, err := discordSession.GuildRoleCreate(gID)
				return role, err
			},
		},
		{
			Name: "GuildRoleDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, rID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.GuildRoleDelete(gID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildRoleEdit",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID, rID, name := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				color, err := strconv.Atoi(util.Atostr(args[3]))
				if err != nil {
					return nil, err
				}
				hoist, err := strconv.ParseBool(util.Atostr(args[4]))
				if err != nil {
					return nil, err
				}
				perm, err := strconv.ParseInt(util.Atostr(args[5]), 10, 64)
				if err != nil {
					return nil, err
				}
				mention, err := strconv.ParseBool(util.Atostr(args[6]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				role, err := discordSession.GuildRoleEdit(gID, rID, name, color, hoist, perm, mention)
				return role, err
			},
		},
		// GuildRoleReorder not supported
		{
			Name: "GuildRoles",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				roles, err := discordSession.GuildRoles(gID)
				return roles, err
			},
		},
		// GuildSplash not supported
		{
			Name: "GuildWebhooks",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				gID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhooks, err := discordSession.GuildWebhooks(gID)
				return webhooks, err
			},
		},
		{
			Name: "HeartbeatLatency",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				latency := discordSession.HeartbeatLatency()
				return latency, nil
			},
		},
		{
			Name: "InviteInfo",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				iID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invite, err := discordSession.Invite(iID)
				return invite, err
			},
		},
		{
			Name: "InviteAccept",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				iID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invite, err := discordSession.InviteAccept(iID)
				return invite, err
			},
		},
		{
			Name: "InviteDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				iID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invite, err := discordSession.InviteDelete(iID)
				return invite, err
			},
		},
		{
			Name: "InviteWithCounts",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				iID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				invite, err := discordSession.InviteWithCounts(iID)
				return invite, err
			},
		},
		{
			Name: "MessageReactionAdd",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID, eID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.MessageReactionAdd(cID, mID, eID)
				return nil, err
			},
		},
		{
			Name: "MessageReactionRemove",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID, eID, uID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2]), util.Atostr(args[3])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.MessageReactionRemove(cID, mID, eID, uID)
				return nil, err
			},
		},
		{
			Name: "MessageReactions",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID, eID, beforeID, afterID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2]), util.Atostr(args[4]), util.Atostr(args[5])
				limit, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				users, err := discordSession.MessageReactions(cID, mID, eID, limit, beforeID, afterID)
				return users, err
			},
		},
		{
			Name: "MessageReactionsRemoveAll",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.MessageReactionsRemoveAll(cID, mID)
				return nil, err
			},
		},
		{
			Name: "MessageReactionsRemoveEmoji",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, mID, eID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.MessageReactionsRemoveEmoji(cID, mID, eID)
				return nil, err
			},
		},
		// Relationship* is not supported
		// Request* is not supported
		{
			Name: "UpdateGameStatus",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				name := util.Atostr(args[1])
				idle, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.UpdateGameStatus(idle, name)
				return nil, err
			},
		},
		{
			Name: "UpdateListeningStatus",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				name := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.UpdateListeningStatus(name)
				return nil, err
			},
		},
		{
			Name: "UpdateStatusComplex",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				var usd discordgo.UpdateStatusData
				json.Unmarshal([]byte(util.Atostr(args[0])), &usd)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.UpdateStatusComplex(usd)
				return nil, err
			},
		},
		{
			Name: "UpdateStreamingStatus",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				name, url := util.Atostr(args[1]), util.Atostr(args[2])
				idle, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err = discordSession.UpdateStreamingStatus(idle, name, url)
				return nil, err
			},
		},
		{
			Name: "UserInfo",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				uID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				user, err := discordSession.User(uID)
				return user, err
			},
		},
		// UserAvatar not supported
		{
			Name: "UserChannelCreate",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				rID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channel, err := discordSession.UserChannelCreate(rID)
				return channel, err
			},
		},
		{
			Name: "UserChannelPermissions",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				rID, cID := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				perms, err := discordSession.UserChannelPermissions(rID, cID)
				return perms, err
			},
		},
		{
			Name: "UserChannels",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				channels, err := discordSession.UserChannels()
				return channels, err
			},
		},
		{
			Name: "UserConnections",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				connections, err := discordSession.UserConnections()
				return connections, err
			},
		},
		{
			Name: "UserGuilds",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				beforeID, afterID := util.Atostr(args[1]), util.Atostr(args[2])
				limit, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				guilds, err := discordSession.UserGuilds(limit, beforeID, afterID)
				return guilds, err
			},
		},
		// UserNoteSet not supported
		// UserSettings not supported
		// UserUpdate not supported
		// UserUpdateStatus not supported
		{
			Name: "VoiceICE",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				voiceICE, err := discordSession.VoiceICE()
				return voiceICE, err
			},
		},
		{
			Name: "VocieRegions",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				vocieRegions, err := discordSession.VoiceRegions()
				return vocieRegions, err
			},
		},
		{
			Name: "WebhookInfo",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhook, err := discordSession.Webhook(wID)
				return webhook, err
			},
		},
		{
			Name: "WebhookCreate",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				cID, name, avatar := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhook, err := discordSession.WebhookCreate(cID, name, avatar)
				return webhook, err
			},
		},
		{
			Name: "WebhookDelete",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID := util.Atostr(args[0])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				err := discordSession.WebhookDelete(wID)
				return nil, err
			},
		},
		{
			Name: "WebhookDeleteWithToken",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID, token := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhook, err := discordSession.WebhookDeleteWithToken(wID, token)
				return webhook, err
			},
		},
		{
			Name: "WebhookEdit",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID, name, avatar, cID := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2]), util.Atostr(args[3])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				role, err := discordSession.WebhookEdit(wID, name, avatar, cID)
				return role, err
			},
		},
		{
			Name: "WebhookEditWithToken",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID, token, name, avatar := util.Atostr(args[0]), util.Atostr(args[1]), util.Atostr(args[2]), util.Atostr(args[3])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				role, err := discordSession.WebhookEditWithToken(wID, token, name, avatar)
				return role, err
			},
		},
		{
			Name: "WebhookExecute",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID, token := util.Atostr(args[0]), util.Atostr(args[1])
				wait, err := strconv.ParseBool(util.Atostr(args[2]))
				if err != nil {
					return nil, err
				}
				var wParams discordgo.WebhookParams
				json.Unmarshal([]byte(util.Atostr(args[1])), &wParams)
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				msg, err := discordSession.WebhookExecute(wID, token, wait, &wParams)
				return msg, err
			},
		},
		{
			Name: "WebhookWithToken",
			Exec: func(g *grug.GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				wID, token := util.Atostr(args[0]), util.Atostr(args[1])
				discordSession := ctx.Value(DISCORDSESSION).(*discordgo.Session)
				webhook, err := discordSession.WebhookWithToken(wID, token)
				return webhook, err
			},
		},
	}...)
}
