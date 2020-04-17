package command

import (
	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Github(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	resp := "Follow my development here: https://github.com/notmeta/osrs.cx"
	_, _ = ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}
