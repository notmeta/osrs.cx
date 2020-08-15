package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

func (m *Mux) Stats(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	username := strings.Join(ctx.Fields[1:], "+")

	if len(username) == 0 {
		username = util.GetRsn(dm.Author)

		if len(username) == 0 {
			_, _ = ds.ChannelMessageSend(dm.ChannelID, "Please specify a username, or set one with `::setrsn [username]`")
			return
		}
	}

	reg := regexp.MustCompile("<@!([0-9]*?)>")
	matches := reg.FindStringSubmatch(username)
	if len(matches) > 0 {
		mem, err := ds.GuildMember(dm.GuildID, matches[1])

		if err != nil {
			log.Println("error finding user '"+matches[1]+"':", err)
		}

		if mem == nil {
			_, _ = ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("User %s not found!", username))
			return
		}

		rsn := util.GetRsn(mem.User)
		if len(rsn) == 0 {
			_, _ = ds.ChannelMessageSend(dm.ChannelID, fmt.Sprintf("%s has no username set!", username))
			return
		}

		username = rsn
	}

	apiUrl := util.GetHiscoresApiUrl(&username)
	friendlyUrl := util.GetFriendlyHiscoresUrl(&username)

	msg, _ := ds.ChannelMessageSendEmbed(dm.ChannelID, &discordgo.MessageEmbed{
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "osrs.cx",
		},
		Description: fmt.Sprintf("Getting stats for `%s`...", strings.Replace(username, "+", " ", -1)),
		Color:       0xFFFF00,
	})

	resp, err := util.Get(apiUrl)

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
		embed := util.ParseHiscore(&username, &body).GenerateHiscoresEmbed()

		statsResponse = embed
	}

	_, _ = ds.ChannelMessageEditEmbed(dm.ChannelID, msg.ID, statsResponse)

	return
}
