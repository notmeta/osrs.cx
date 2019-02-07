package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const apiUrl = "https://crystalmathlabs.com/tracker/api.php?type=update"

func (m *Mux) Update(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	username := strings.Join(ctx.Fields[1:], "+")

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Updating `%s`...", strings.Replace(username, "+", " ", -1)),
		Color:       0xFFFF00,
	})

	client := &http.Client{}

	req, _ := http.NewRequest("GET", apiUrl+"&player="+username, nil)
	req.Header.Add("User-Agent", fmt.Sprintf("osrs.cx/%s (+https://github.com/notmeta/osrs.cx)", runtime.Version()))

	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)

	response := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: username,
			//URL:  *friendlyUrl,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
	}

	if resp.StatusCode != 200 {
		response.Description = fmt.Sprintf("Failed to find stats for `%s`!\n(%s)", username, resp.Status)
		response.Color = 0xFF0000
	} else {
		body := string(body)
		cmlResponse, colour := getCMLReponse(body)
		response.Description = cmlResponse
		response.Color = colour
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, response)

}

func getCMLReponse(returnCode string) (response string, colour int) {

	switch strings.TrimSpace(returnCode) {
	case "1":
		return "Success!", 0x00FF00
	case "2":
		return "Player not on RuneScape hiscores.", 0xFF0000
	case "3":
		return "Negative XP gain detected.", 0xFF0000
	case "4":
		return "Unknown error.", 0xFF0000
	case "5":
		return "This player has been updated within the last 30 seconds.", 0xFFFF00
	case "6":
		return "The player name was invalid.", 0xFF0000
	default:
		does := returnCode == "1"
		return "Unknown return code: '" + returnCode + "'" + strconv.FormatBool(does), 0xFF0000

	}
}
