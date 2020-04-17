package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"strconv"
	"strings"
	"time"
)

const maxSearchResult = 10

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *Mux) Search(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	query := strings.Join(ctx.Fields[1:], " ")

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Searching for `%s`...", query),
		Color:       0xFFFF00,
	})

	itemIds := util.FindItemIds(query)
	maxAmount := min(len(itemIds), maxSearchResult)

	if len(itemIds) == 0 {
		noItemFoundResponse := &discordgo.MessageEmbed{
			Timestamp: time.Now().Format(time.RFC3339),
			Footer: &discordgo.MessageEmbedFooter{
				Text: "osrs.cx",
			},
			Description: fmt.Sprintf("No items found for `%s`!", query),
			Color:       0xFF0000,
		}

		_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, noItemFoundResponse)
		return
	}

	description := fmt.Sprintf("%d results for `%s`", len(itemIds), query)

	if len(itemIds) > maxSearchResult {
		description += ", showing top " + strconv.Itoa(maxAmount)
	}

	resp := &discordgo.MessageEmbed{
		Timestamp:   time.Now().Format(time.RFC3339),
		Description: description,
		Color:       0x00FF00,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx | RuneLite/RuneScape API",
		},
	}

	for _, id := range itemIds[:maxAmount] {
		item := util.RunescapeItemForId(id).Item

		resp.Fields = append(resp.Fields,
			&discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("%s - %d", item.Name, item.Id),
				Value:  item.Description,
				Inline: false,
			})
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, resp)

	return
}
