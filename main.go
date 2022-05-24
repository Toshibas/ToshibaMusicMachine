package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ready(s *discordgo.Session, event *discordgo.Ready) {

	fmt.Println("ToshibaMusicMachine is now running.  Press CTRL-C to exit.")

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	

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
