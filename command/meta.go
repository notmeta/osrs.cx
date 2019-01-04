package command

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func (m *Mux) Meta(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	uniqueMembers := make(map[string]bool)
	totalChannels := 0

	for _, guild := range ds.State.Guilds {
		totalChannels += len(guild.Channels)

		for _, member := range guild.Members {
			uniqueMembers[member.User.ID] = true
		}

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
				Name:   "Guilds",
				Value:  strconv.Itoa(len(ds.State.Guilds)),
				Inline: true,
			},
			{
				Name:   "Unique Users",
				Value:  strconv.Itoa(len(uniqueMembers)),
				Inline: true,
			},
			{
				Name:   "Channels",
				Value:  strconv.Itoa(totalChannels),
				Inline: true,
			},
		},
	}

	_, _ = ds.ChannelMessageSendEmbed(dm.ChannelID, resp)

	return
}
