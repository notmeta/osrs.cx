package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func (m *Mux) Stats(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	username := strings.Join(ctx.Fields[1:], "+")

	apiUrl := util.GetHiscoresApiUrl(&username)
	friendlyUrl := util.GetFriendlyHiscoresUrl(&username)

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Getting stats for `%s`...", username),
	})

	client := &http.Client{}

	req, _ := http.NewRequest("GET", *apiUrl, nil)
	req.Header.Add("User-Agent", fmt.Sprintf("osrs.cx/%s (+https://github.com/notmeta/osrs.cx)", runtime.Version()))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
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
		hiscore := util.ParseHiscore(&username, &body)
		embed := util.GenerateHiscoresEmbed(hiscore)

		statsResponse = embed
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, statsResponse)

	return
}
