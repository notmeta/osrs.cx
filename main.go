package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Version is a constant that stores the Disgord version information.
const Version = "v0.0.1"

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	Session.Token = os.Getenv("DG_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {

	// Declare any variables needed later.
	var err error

	// Print out a fancy logo!
	fmt.Printf(`

 ██████╗ ███████╗██████╗ ███████╗    ██████╗██╗  ██╗
██╔═══██╗██╔════╝██╔══██╗██╔════╝   ██╔════╝╚██╗██╔╝
██║   ██║███████╗██████╔╝███████╗   ██║      ╚███╔╝ 
██║   ██║╚════██║██╔══██╗╚════██║   ██║      ██╔██╗ 
╚██████╔╝███████║██║  ██║███████║██╗╚██████╗██╔╝ ██╗
 ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝ %-16s`+"\n\n", Version)

	// Parse command line arguments
	flag.Parse()

	// Verify a Token was provided
	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	_ = Session.UpdateStatus(0, "::help | osrs.cx")

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}
