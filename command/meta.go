package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"time"
)

func (m *Mux) Meta(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	totalMembers := 0
	totalChannels := 0

	for _, guild := range ds.State.Guilds {
		totalChannels += len(guild.Channels)
		totalMembers += guild.MemberCount
	}

	resp := &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ds.State.User.Username,
			IconURL: ds.State.User.AvatarURL(""),
		},
		Color: 0x00FF00,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Servers",
				Value:  util.RenderInteger("#,###.", len(ds.State.Guilds)),
				Inline: true,
			},
			{
				Name:   "Users",
				Value:  util.RenderInteger("#,###.", totalMembers),
				Inline: true,
			},
			{
				Name:   "Channels",
				Value:  util.RenderInteger("#,###.", totalChannels),
				Inline: true,
			},
		},
	}

	_, _ = ds.ChannelMessageSendEmbed(dm.ChannelID, resp)

	return
}
