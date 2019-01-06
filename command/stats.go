package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/model"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func (m *Mux) Stats(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	username := strings.Join(ctx.Fields[1:], "+")

	apiUrl := model.GetHiscoresApiUrl(&username)
	friendlyUrl := model.GetFriendlyHiscoresUrl(&username)

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Getting stats for `%s`...", strings.Replace(username, "+", " ", -1)),
		Color:       0xFFFF00,
	})

	client := &http.Client{}

	req, _ := http.NewRequest("GET", *apiUrl, nil)
	req.Header.Add("User-Agent", fmt.Sprintf("osrs.cx/%s (+https://github.com/notmeta/osrs.cx)", runtime.Version()))

	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)

	statsResponse := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: username,
			URL:  *friendlyUrl,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
	}

	if resp.StatusCode != 200 {
		statsResponse.Description = fmt.Sprintf("Failed to find stats for `%s`!\n(%s)", username, resp.Status)
		statsResponse.Color = 0xFF0000
	} else {
		body := string(body)
		embed := model.ParseHiscore(&username, &body).GenerateHiscoresEmbed()

		statsResponse = embed
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, statsResponse)

	return
}
