package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, event *discordgo.Ready) {

	fmt.Println("ToshibaMusicMachine is now running. Press CTRL-C to exit.")

}

func messageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Message.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(e.Message.Content, "/") {

		commandAndRest := strings.SplitN(e.Message.Content, " ", 2)
		command := commandAndRest[0]
		commandName := command[1:]

		if commandName == "play" {
			play(s, e, commandAndRest[1])
		}

	}

}

func findVoiceChannel(s *discordgo.Session, e *discordgo.MessageCreate) (string, error) {

	c, err := s.State.Channel(e.ChannelID)
	if err != nil {
		return "", err
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return "", err
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == e.Author.ID {
			return vs.ChannelID, nil
		}
	}

	return "", errors.New("Channel not found.")

}

func loadSound(path string) ([][]byte, error) {

	// encodeSession, err := dca.EncodeFile(path, dca.StdEncodeOptions)

	// if err != nil {
	// 	return nil, err
	// }

	// defer encodeSession.Cleanup()

	// file, err := os.Create("./output.dca")

	// defer file.Close()

	// if err != nil {
	// 	return nil, err
	// }

	// io.Copy(file, encodeSession)

	file, err := os.Open("./test.dca")
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return nil, err
	}

	var buffer = make([][]byte, 0)

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return buffer, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}

}

func play(s *discordgo.Session, e *discordgo.MessageCreate, params string) {

	buf, err := loadSound("./airhorn.dca")

	if err != nil {
		log.Fatal(err)
	}

	cid, err := findVoiceChannel(s, e)

	vc, err := s.ChannelVoiceJoin(e.Message.GuildID, cid, false, true)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(250 * time.Millisecond)

	vc.Speaking(true)

	for _, buff := range buf {
		vc.OpusSend <- buff
	}

	vc.Speaking(false)

	time.Sleep(250 * time.Millisecond)

	vc.Disconnect()

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
