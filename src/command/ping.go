package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func (m *Mux) Ping(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	now := time.Now()

	timeSent, _ := dm.Timestamp.Parse()
	diff := (now.UnixNano() / 1000000) - (timeSent.UnixNano() / 1000000)

	resp := fmt.Sprintf(":ping_pong: Pong! `%dms`", diff)

	_, _ = ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}
