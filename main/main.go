package main

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/legaldiscordbot/core"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

func main() {
	core.AppContext = &core.Context{}

	// Load Config file
	err := core.LoadConfigFile("config.json")
	if err != nil {
		fmt.Printf("Error occured whilst loading configuration file: %s\n\n", err.Error())
		return
	}

	core.AppContext.Discord, err = discordgo.New("Bot " + core.AppContext.Config.Token)
	if err != nil {
		fmt.Printf("Error! Failed to create Discord session: %s\n", err.Error())
		return
	}

	core.AppContext.Discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Successfully initalised connection to Discord.")
	})

	err = core.AppContext.Discord.Open()
	if err != nil {
		fmt.Printf("Error! Failed to connect to Discord: %s\n", err.Error())
		return
	}

	err = core.AppContext.Discord.UpdateGameStatus(0, "My clients are dumb as hell")
	if err != nil {
		log.Printf("Error setting bot status: %s\n", err.Error())
	}

	core.AppContext.Discord.AddHandler(core.MessageCreateHandler)
	core.AppContext.Discord.AddHandler(core.Blacklist)
	core.AppContext.Discord.AddHandler(core.Joinvoicechat)

	defer func(Discord *discordgo.Session) {
		_ = Discord.Close()
	}(core.AppContext.Discord)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("\n\nShutting down the Lawyer...")

}
