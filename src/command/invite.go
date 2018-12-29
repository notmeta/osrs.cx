package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Invite(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	resp := fmt.Sprintf("Invite me to your server using this link: https://discordapp.com/api/oauth2/authorize?client_id=%s&scope=bot&permissions=%d",
		ds.State.User.ID, 67464256)

	channel, errCreate := ds.UserChannelCreate(dm.Author.ID)
	_, errSend := ds.ChannelMessageSend(channel.ID, resp)

	if errCreate != nil || errSend != nil {
		resp = "Failed to send pm!"
	} else {
		resp = "PM sent!"
	}

	_, _ = ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}
