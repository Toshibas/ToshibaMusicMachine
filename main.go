package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, event *discordgo.Ready) {

	fmt.Println("ToshibaMusicMachine is now running.  Press CTRL-C to exit.")

}

func messageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Message.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(e.Message.Content, "/") {

		commandAndRest := strings.SplitN(e.Message.Content, " ", 2)
		command := commandAndRest[0]
		commandName := command[1:]

		var content = "ok " + commandName + " ^-^"

		s.ChannelMessageSend(e.Message.ChannelID, content)
		if len(commandAndRest) > 1 {
			s.ChannelMessageSend(e.Message.ChannelID, commandAndRest[1])
		}

	}

}

func main() {

	bot, err := discordgo.New("Bot " + "OTc4NjkxNDk2MTkyNjU5NDU2.GwS474.IVRh60zMfVZqWSeL87ziEg_kxRhxm2tTYkGg1I")

	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(ready)

	bot.AddHandler(messageCreate)

	err = bot.Open()

	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()

}
