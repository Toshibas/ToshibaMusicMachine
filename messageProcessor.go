package main

import (
	// "fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageProcessor struct {
	defaultGuildStatus bool
	dbConnector        *DBConnector
}

func NewMessageProcessor(config *Config, dbConnector *DBConnector) *MessageProcessor {

	return &MessageProcessor{
		defaultGuildStatus: config.DefGuildStatus,
		dbConnector:        dbConnector,
	}

}

func (mp *MessageProcessor) ProcessMessage(s *discordgo.Session, e *discordgo.MessageCreate) {

	if e.Message.Author.ID == s.State.User.ID {
		return
	}

	exist := mp.dbConnector.GuildExists(e.Message.GuildID)

	if !exist {
		mp.dbConnector.AddGuild(e.Message.GuildID, mp.defaultGuildStatus)
	}

	guildInfo := mp.dbConnector.GetGuildByID(e.Message.GuildID)

	if guildInfo.Allowed == false {

		s.ChannelMessageSend(e.Message.ChannelID, "You shall not pass!!!")

		return

	}

	handleCommand(s, e)

}

func handleCommand(s *discordgo.Session, e *discordgo.MessageCreate) {

}
