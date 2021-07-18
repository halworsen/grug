package grug

import (
	"encoding/json"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			Name: "ChannelDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				channel, err := g.DiscordSession.ChannelDelete(cID)
				return channel, err
			},
		},
		{
			Name: "ChannelEdit",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, newName := atostr(args[0]), atostr(args[1])
				channel, err := g.DiscordSession.ChannelEdit(cID, newName)
				return channel, err
			},
		},
		{
			Name: "ChannelEditComplex",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				var cEdit discordgo.ChannelEdit
				json.Unmarshal([]byte(atostr(args[1])), &cEdit)
				channel, err := g.DiscordSession.ChannelEditComplex(cID, &cEdit)
				return channel, err
			},
		},
		// ChannelFileSend not supported
		// ChannelFileSendWithMessage not supported
		{
			Name: "ChannelInviteCreate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				var cInvite discordgo.Invite
				json.Unmarshal([]byte(atostr(args[1])), &cInvite)
				invite, err := g.DiscordSession.ChannelInviteCreate(cID, cInvite)
				return invite, err
			},
		},
		{
			Name: "ChannelInvites",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				invites, err := g.DiscordSession.ChannelInvites(cID)
				return invites, err
			},
		},
		{
			Name: "ChannelMessage",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				msg, err := g.DiscordSession.ChannelMessage(cID, mID)
				return msg, err
			},
		},
		// ChannelMessageAck not supported
		{
			Name: "ChannelMessageCrosspost",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				msg, err := g.DiscordSession.ChannelMessageCrosspost(cID, mID)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.ChannelMessageDelete(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessageEdit",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID, content := atostr(args[0]), atostr(args[1]), atostr(args[2])
				msg, err := g.DiscordSession.ChannelMessageEdit(cID, mID, content)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageEditComplex",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				var msgEdit discordgo.MessageEdit
				json.Unmarshal([]byte(atostr(args[0])), &msgEdit)
				msg, err := g.DiscordSession.ChannelMessageEditComplex(&msgEdit)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageEditEmbed",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				var msgEmbed discordgo.MessageEmbed
				json.Unmarshal([]byte(atostr(args[2])), &msgEmbed)
				msg, err := g.DiscordSession.ChannelMessageEditEmbed(cID, mID, &msgEmbed)
				return msg, err
			},
		},
		{
			Name: "ChannelMessagePin",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.ChannelMessagePin(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessageSend",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, msgContent := atostr(args[0]), atostr(args[1])
				msg, err := g.DiscordSession.ChannelMessageSend(cID, msgContent)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendComplex",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				var msgSend discordgo.MessageSend
				json.Unmarshal([]byte(atostr(args[1])), &msgSend)
				msg, err := g.DiscordSession.ChannelMessageSendComplex(cID, &msgSend)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendEmbed",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				var msgEmbed discordgo.MessageEmbed
				json.Unmarshal([]byte(atostr(args[1])), &msgEmbed)
				msg, err := g.DiscordSession.ChannelMessageSendEmbed(cID, &msgEmbed)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendReply",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, msgContent := atostr(args[0]), atostr(args[1])
				var msgRef discordgo.MessageReference
				json.Unmarshal([]byte(atostr(args[1])), &msgRef)
				msg, err := g.DiscordSession.ChannelMessageSendReply(cID, msgContent, &msgRef)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageSendTTS",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, msgContent := atostr(args[0]), atostr(args[1])
				msg, err := g.DiscordSession.ChannelMessageSendTTS(cID, msgContent)
				return msg, err
			},
		},
		{
			Name: "ChannelMessageUnpin",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.ChannelMessageUnpin(cID, mID)
				return nil, err
			},
		},
		{
			Name: "ChannelMessages",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, beforeID, afterID, aroundID := atostr(args[0]), atostr(args[2]), atostr(args[3]), atostr(args[4])
				limit, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				msgs, err := g.DiscordSession.ChannelMessages(cID, limit, beforeID, afterID, aroundID)
				return msgs, err
			},
		},
		// ChannelMessagesBulkDelete not supported
		{
			Name: "ChannelMessagesPinned",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				msgs, err := g.DiscordSession.ChannelMessagesPinned(cID)
				return msgs, err
			},
		},
		// ChannelNewsFollow not supported
		// ChannelPermissionDelete not supported
		// ChannelPermissionSet not supported
		{
			Name: "ChannelTyping",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				err := g.DiscordSession.ChannelTyping(cID)
				return nil, err
			},
		},
		// ChannelVoiceJoin not supported
		// ChannelVoiceJoinManual not supported
		{
			Name: "ChannelWebhooks",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				webhooks, err := g.DiscordSession.ChannelWebhooks(cID)
				return webhooks, err
			},
		},
		{
			Name: "ChannelWebhooks",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID := atostr(args[0])
				webhooks, err := g.DiscordSession.ChannelWebhooks(cID)
				return webhooks, err
			},
		},
		{
			Name: "GuildInfo",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				guild, err := g.DiscordSession.Guild(gID)
				return guild, err
			},
		},
		{
			Name: "GuildAuditLog",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uId, beforeID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				actionType, err := strconv.Atoi(atostr(args[3]))
				if err != nil {
					return nil, err
				}
				limit, err := strconv.Atoi(atostr(args[4]))
				if err != nil {
					return nil, err
				}
				auditLog, err := g.DiscordSession.GuildAuditLog(gID, uId, beforeID, actionType, limit)
				return auditLog, err
			},
		},
		{
			Name: "GuildBan",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uId := atostr(args[0]), atostr(args[1])
				ban, err := g.DiscordSession.GuildBan(gID, uId)
				return ban, err
			},
		},
		{
			Name: "GuildBanCreate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uId := atostr(args[0]), atostr(args[1])
				days, err := strconv.Atoi(atostr(args[2]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.GuildBanCreate(gID, uId, days)
				return nil, err
			},
		},
		{
			Name: "GuildBanCreateWithReason",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uId, reason := atostr(args[0]), atostr(args[1]), atostr(args[2])
				days, err := strconv.Atoi(atostr(args[3]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.GuildBanCreateWithReason(gID, uId, reason, days)
				return nil, err
			},
		},
		{
			Name: "GuildBanDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uId := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.GuildBanDelete(gID, uId)
				return nil, err
			},
		},
		{
			Name: "GuildBans",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				bans, err := g.DiscordSession.GuildBans(gID)
				return bans, err
			},
		},
		// GuildChannelCreate not supported
		{
			Name: "GuildChannelCreateComplex",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				var channelCreateData discordgo.GuildChannelCreateData
				json.Unmarshal([]byte(atostr(args[1])), &channelCreateData)
				channel, err := g.DiscordSession.GuildChannelCreateComplex(gID, channelCreateData)
				return channel, err
			},
		},
		{
			Name: "GuildChannels",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				channels, err := g.DiscordSession.GuildChannels(gID)
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
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, eID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.GuildEmojiDelete(gID, eID)
				return nil, err
			},
		},
		// GuildEmojiEdit not supported
		{
			Name: "GuildEmojis",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				emojis, err := g.DiscordSession.GuildEmojis(gID)
				return emojis, err
			},
		},
		{
			Name: "GuildIcon",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				icon, err := g.DiscordSession.GuildIcon(gID)
				return icon, err
			},
		},
		// GuildIntegrationCreate not supported
		// GuildIntegrationDelete not supported
		// GuildIntegrationEdit not supported
		// GuildIntegrationSync not supported
		{
			Name: "GuildIntegrations",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				integrations, err := g.DiscordSession.GuildIntegrations(gID)
				return integrations, err
			},
		},
		{
			Name: "GuildInvites",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				invites, err := g.DiscordSession.GuildInvites(gID)
				return invites, err
			},
		},
		{
			Name: "GuildMember",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID := atostr(args[0]), atostr(args[1])
				member, err := g.DiscordSession.GuildMember(gID, uID)
				return member, err
			},
		},
		// GuildMemberAdd not supported
		{
			Name: "GuildMemberDeafen",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID := atostr(args[0]), atostr(args[1])
				deaf, err := strconv.ParseBool(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.GuildMemberDeafen(gID, uID, deaf)
				return nil, err
			},
		},
		{
			Name: "GuildMemberDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.GuildMemberDelete(gID, uID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberDeleteWithReason",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID, reason := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.GuildMemberDeleteWithReason(gID, uID, reason)
				return nil, err
			},
		},
		// GuildMemberEdit not supported
		{
			Name: "GuildMemberMove",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID, cID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.GuildMemberMove(gID, uID, &cID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberMute",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID := atostr(args[0]), atostr(args[1])
				mute, err := strconv.ParseBool(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.GuildMemberMute(gID, uID, mute)
				return nil, err
			},
		},
		{
			Name: "GuildMemberNickname",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID, nick := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.GuildMemberNickname(gID, uID, nick)
				return nil, err
			},
		},
		{
			Name: "GuildMemberRoleAdd",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID, rID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.GuildMemberRoleAdd(gID, uID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildMemberRoleRemove",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, uID, rID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.GuildMemberRoleRemove(gID, uID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildMembers",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, after := atostr(args[0]), atostr(args[1])
				limit, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				members, err := g.DiscordSession.GuildMembers(gID, after, limit)
				return members, err
			},
		},
		{
			Name: "GuildPrune",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				days, err := strconv.ParseUint(atostr(args[1]), 10, 32)
				if err != nil {
					return nil, err
				}
				count, err := g.DiscordSession.GuildPrune(gID, uint32(days))
				return count, err
			},
		},
		{
			Name: "GuildPruneCount",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				days, err := strconv.ParseUint(atostr(args[1]), 10, 32)
				if err != nil {
					return nil, err
				}
				count, err := g.DiscordSession.GuildPruneCount(gID, uint32(days))
				return count, err
			},
		},
		{
			Name: "GuildRoleCreate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				role, err := g.DiscordSession.GuildRoleCreate(gID)
				return role, err
			},
		},
		{
			Name: "GuildRoleDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, rID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.GuildRoleDelete(gID, rID)
				return nil, err
			},
		},
		{
			Name: "GuildRoleEdit",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID, rID, name := atostr(args[0]), atostr(args[1]), atostr(args[2])
				color, err := strconv.Atoi(atostr(args[3]))
				if err != nil {
					return nil, err
				}
				hoist, err := strconv.ParseBool(atostr(args[4]))
				if err != nil {
					return nil, err
				}
				perm, err := strconv.ParseInt(atostr(args[5]), 10, 64)
				if err != nil {
					return nil, err
				}
				mention, err := strconv.ParseBool(atostr(args[6]))
				if err != nil {
					return nil, err
				}
				role, err := g.DiscordSession.GuildRoleEdit(gID, rID, name, color, hoist, perm, mention)
				return role, err
			},
		},
		// GuildRoleReorder not supported
		{
			Name: "GuildRoles",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				roles, err := g.DiscordSession.GuildRoles(gID)
				return roles, err
			},
		},
		// GuildSplash not supported
		{
			Name: "GuildWebhooks",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				gID := atostr(args[0])
				webhooks, err := g.DiscordSession.GuildWebhooks(gID)
				return webhooks, err
			},
		},
		{
			Name: "HeartbeatLatency",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				latency := g.DiscordSession.HeartbeatLatency()
				return latency, nil
			},
		},
		{
			Name: "InviteInfo",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				iID := atostr(args[0])
				invite, err := g.DiscordSession.Invite(iID)
				return invite, err
			},
		},
		{
			Name: "InviteAccept",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				iID := atostr(args[0])
				invite, err := g.DiscordSession.InviteAccept(iID)
				return invite, err
			},
		},
		{
			Name: "InviteDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				iID := atostr(args[0])
				invite, err := g.DiscordSession.InviteDelete(iID)
				return invite, err
			},
		},
		{
			Name: "InviteWithCounts",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				iID := atostr(args[0])
				invite, err := g.DiscordSession.InviteWithCounts(iID)
				return invite, err
			},
		},
		{
			Name: "MessageReactionAdd",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID, eID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.MessageReactionAdd(cID, mID, eID)
				return nil, err
			},
		},
		{
			Name: "MessageReactionRemove",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID, eID, uID := atostr(args[0]), atostr(args[1]), atostr(args[2]), atostr(args[3])
				err := g.DiscordSession.MessageReactionRemove(cID, mID, eID, uID)
				return nil, err
			},
		},
		{
			Name: "MessageReactions",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID, eID, beforeID, afterID := atostr(args[0]), atostr(args[1]), atostr(args[2]), atostr(args[4]), atostr(args[5])
				limit, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				users, err := g.DiscordSession.MessageReactions(cID, mID, eID, limit, beforeID, afterID)
				return users, err
			},
		},
		{
			Name: "MessageReactionsRemoveAll",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID := atostr(args[0]), atostr(args[1])
				err := g.DiscordSession.MessageReactionsRemoveAll(cID, mID)
				return nil, err
			},
		},
		{
			Name: "MessageReactionsRemoveEmoji",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, mID, eID := atostr(args[0]), atostr(args[1]), atostr(args[2])
				err := g.DiscordSession.MessageReactionsRemoveEmoji(cID, mID, eID)
				return nil, err
			},
		},
		// Relationship* is not supported
		// Request* is not supported
		{
			Name: "UpdateGameStatus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				name := atostr(args[1])
				idle, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.UpdateGameStatus(idle, name)
				return nil, err
			},
		},
		{
			Name: "UpdateListeningStatus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				name := atostr(args[0])
				err := g.DiscordSession.UpdateListeningStatus(name)
				return nil, err
			},
		},
		{
			Name: "UpdateStatusComplex",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				var usd discordgo.UpdateStatusData
				json.Unmarshal([]byte(atostr(args[0])), &usd)
				err := g.DiscordSession.UpdateStatusComplex(usd)
				return nil, err
			},
		},
		{
			Name: "UpdateStreamingStatus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				name, url := atostr(args[1]), atostr(args[2])
				idle, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				err = g.DiscordSession.UpdateStreamingStatus(idle, name, url)
				return nil, err
			},
		},
		{
			Name: "UserInfo",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				uID := atostr(args[0])
				user, err := g.DiscordSession.User(uID)
				return user, err
			},
		},
		// UserAvatar not supported
		{
			Name: "UserChannelCreate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				rID := atostr(args[0])
				channel, err := g.DiscordSession.UserChannelCreate(rID)
				return channel, err
			},
		},
		{
			Name: "UserChannelPermissions",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				rID, cID := atostr(args[0]), atostr(args[1])
				perms, err := g.DiscordSession.UserChannelPermissions(rID, cID)
				return perms, err
			},
		},
		{
			Name: "UserChannels",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				channels, err := g.DiscordSession.UserChannels()
				return channels, err
			},
		},
		{
			Name: "UserConnections",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				connections, err := g.DiscordSession.UserConnections()
				return connections, err
			},
		},
		{
			Name: "UserGuilds",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				beforeID, afterID := atostr(args[1]), atostr(args[2])
				limit, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				guilds, err := g.DiscordSession.UserGuilds(limit, beforeID, afterID)
				return guilds, err
			},
		},
		// UserNoteSet not supported
		// UserSettings not supported
		// UserUpdate not supported
		// UserUpdateStatus not supported
		{
			Name: "VoiceICE",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				voiceICE, err := g.DiscordSession.VoiceICE()
				return voiceICE, err
			},
		},
		{
			Name: "VocieRegions",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				vocieRegions, err := g.DiscordSession.VoiceRegions()
				return vocieRegions, err
			},
		},
		{
			Name: "WebhookInfo",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID := atostr(args[0])
				webhook, err := g.DiscordSession.Webhook(wID)
				return webhook, err
			},
		},
		{
			Name: "WebhookCreate",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				cID, name, avatar := atostr(args[0]), atostr(args[1]), atostr(args[2])
				webhook, err := g.DiscordSession.WebhookCreate(cID, name, avatar)
				return webhook, err
			},
		},
		{
			Name: "WebhookDelete",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID := atostr(args[0])
				err := g.DiscordSession.WebhookDelete(wID)
				return nil, err
			},
		},
		{
			Name: "WebhookDeleteWithToken",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID, token := atostr(args[0]), atostr(args[1])
				webhook, err := g.DiscordSession.WebhookDeleteWithToken(wID, token)
				return webhook, err
			},
		},
		{
			Name: "WebhookEdit",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID, name, avatar, cID := atostr(args[0]), atostr(args[1]), atostr(args[2]), atostr(args[3])
				role, err := g.DiscordSession.WebhookEdit(wID, name, avatar, cID)
				return role, err
			},
		},
		{
			Name: "WebhookEditWithToken",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID, token, name, avatar := atostr(args[0]), atostr(args[1]), atostr(args[2]), atostr(args[3])
				role, err := g.DiscordSession.WebhookEditWithToken(wID, token, name, avatar)
				return role, err
			},
		},
		{
			Name: "WebhookExecute",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID, token := atostr(args[0]), atostr(args[1])
				wait, err := strconv.ParseBool(atostr(args[2]))
				if err != nil {
					return nil, err
				}
				var wParams discordgo.WebhookParams
				json.Unmarshal([]byte(atostr(args[1])), &wParams)
				msg, err := g.DiscordSession.WebhookExecute(wID, token, wait, &wParams)
				return msg, err
			},
		},
		{
			Name: "WebhookWithToken",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				wID, token := atostr(args[0]), atostr(args[1])
				webhook, err := g.DiscordSession.WebhookWithToken(wID, token)
				return webhook, err
			},
		},
	}...)
}
