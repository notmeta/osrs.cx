package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/notmeta/osrs.cx/util"
	"log"
	"strings"
)

func (m *Mux) SetRsn(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	rsn := strings.Join(ctx.Fields[1:], "+")

	if len(rsn) == 0 {
		_, _ = ds.ChannelMessageSend(dm.ChannelID, "No username given!")
		return
	}

	err := util.Store.Set(fmt.Sprintf(util.RsnKeyFormat, dm.Author.ID), rsn, 0).Err()

	if err != nil {
		log.Panic(err)
		_, _ = ds.ChannelMessageSend(dm.ChannelID, err.Error())
	} else {
		_ = ds.MessageReactionAdd(dm.ChannelID, dm.ID, "👌")
	}
}
