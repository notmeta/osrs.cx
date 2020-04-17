package main

import (
	"github.com/notmeta/osrs.cx/command"
	"os"
)

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.
var Router = command.New()

const DefaultPrefix = "::"

func init() {
	if os.Getenv("PREFIX") != "" {
		Router.Prefix = os.Getenv("PREFIX")
	} else {
		Router.Prefix = DefaultPrefix
	}

	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	Session.AddHandler(Router.OnMessageCreate)

	// Register the build-in help command.
	_, _ = Router.Route("help", "Display this message.", Router.Help, "commands")
	_, _ = Router.Route("invite", "PMs a link to invite this bot to your server.", Router.Invite)
	_, _ = Router.Route("ping", "", Router.Ping)
	_, _ = Router.Route("stats", "Get hiscores for a given account.", Router.Stats, "hiscores", "hs")
	_, _ = Router.Route("github", "", Router.Github)
	_, _ = Router.Route("meta", "Meta statistics for the bot.", Router.Meta)
	_, _ = Router.Route("price", "Find the exchange price of an item.", Router.Price)
	_, _ = Router.Route("setrsn", "Remember your username for future commands.", Router.SetRsn)
}
