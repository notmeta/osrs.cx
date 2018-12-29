package main

import "command"

// This file adds the Disgord message route multiplexer, aka "command router".
// to the Disgord bot. This is an optional addition however it is included
// by default to demonstrate how to extend the Disgord bot.

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
	_, _ = Router.Route("help", "Display this message.", Router.Help)
	_, _ = Router.Route("invite", "PMs a link to invite this bot to your server.", Router.Invite)
	_, _ = Router.Route("ping", "Pong!", Router.Ping)
}
