package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"strings"
	"time"
)

func (m *Mux) Price(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	query := strings.Join(ctx.Fields[1:], "+")

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Finding prices for `%s`...", strings.Replace(query, "+", " ", -1)),
		Color:       0xFFFF00,
	})

	results := util.SearchItem(query)

	if len(results.Items) == 0 {
		noItemFoundResponse := &discordgo.MessageEmbed{
			Timestamp: time.Now().Format(time.RFC3339),
			Footer: &discordgo.MessageEmbedFooter{
				Text: "osrs.cx",
			},
			Description: fmt.Sprintf("Couldn't find prices for `%s`!", strings.Replace(query, "+", " ", -1)),
			Color:       0xFF0000,
		}

		_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, noItemFoundResponse)
		return
	}

	item := results.Items[0]
	itemPrice := item.GetOSBPrice()

	resp := &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    item.Name,
			IconURL: item.GetIconUrl(),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: item.GetLargeIconUrl(),
		},
		Description: "_" + item.Description + "_",
		Color:       0x00FF00,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx | RuneLite/OSBuddy API",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Buy Price",
				Value:  util.RenderInteger("#,###.", itemPrice.BuyAverage) + " gp",
				Inline: true,
			},
			{
				Name:   "Overall Price",
				Value:  util.RenderInteger("#,###.", itemPrice.OverallAverage) + " gp",
				Inline: true,
			},
			{
				Name:   "Sell Price",
				Value:  util.RenderInteger("#,###.", itemPrice.SellAverage) + " gp",
				Inline: false,
			},
		},
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, resp)

	return
}
