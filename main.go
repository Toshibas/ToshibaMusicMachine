package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, e *discordgo.Ready) {

	fmt.Println("Hello honey! ToshibaMusicMachine is running. Press Ctrl-C to stop.")

}

func main() {

	config := LoadConfiguration("./config.json")

	dbConnector, err := NewDBConnector(
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
	)

	if err != nil {
		log.Fatal(err)
	}

	messageProcessor := NewMessageProcessor(config)

	bot, err := discordgo.New("Bot " + config.BotToken)

	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(ready)
	bot.AddHandler(messageProcessor.ProcessMessage)

	err = bot.Open()

	if err != nil {
		log.Fatal("Error opening session", err)
	}

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()

	dbConnector.CloseConnection()

}
