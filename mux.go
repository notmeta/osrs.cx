package main

import "github.com/notmeta/osrs.cx/command"

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.
var Router = command.New()

const DefaultPrefix = "::"

func init() {
	Router.Prefix = DefaultPrefix

	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	Session.AddHandler(Router.OnMessageCreate)

	// Register the build-in help command.
	_, _ = Router.Route("help", "Display this message.", Router.Help, "commands")
	_, _ = Router.Route("invite", "PMs a link to invite this bot to your server.", Router.Invite)
	_, _ = Router.Route("ping", "Pong!", Router.Ping)
	_, _ = Router.Route("stats", "Get hiscores for a given account.", Router.Stats, "hiscores", "hs")
}
