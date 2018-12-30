package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func (m *Mux) Stats(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	username := strings.Join(ctx.Fields[1:], "+")
	msg, _ := ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("Getting stats for `%s`", username))

	client := &http.Client{}
	url := fmt.Sprintf("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=%s", username)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", fmt.Sprintf("osrs.cx/%s (+https://github.com/notmeta/osrs.cx)", runtime.Version()))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	statsResponse := ""

	if resp.StatusCode != 200 {
		statsResponse = fmt.Sprintf("Failed to find stats for `%s`!", username)
	} else {
		statsResponse = string(body)
	}

	_, _ = ds.ChannelMessageEdit(dm.ChannelID, msg.ID, statsResponse)

	return
}
