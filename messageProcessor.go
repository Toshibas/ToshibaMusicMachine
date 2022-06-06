package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageProcessor struct {
	defaultGuildStatus bool
}

func NewMessageProcessor(config *Config) *MessageProcessor {

	return &MessageProcessor{
		defaultGuildStatus: config.DefGuildStatus,
	}

}

func (mp *MessageProcessor) ProcessMessage(s *discordgo.Session, e *discordgo.MessageCreate) {

	fmt.Println(e.Message.Content)

}
