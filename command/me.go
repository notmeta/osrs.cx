package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
)

func (m *Mux) Me(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	rsn := util.GetRsn(dm.Author)

	if len(rsn) == 0 {
		_, _ = ds.ChannelMessageSend(dm.ChannelID, "No username set! Set one using `::setrsn [username]`")
		return
	}

	_, _ = ds.ChannelMessageSend(dm.ChannelID, rsn)
}
