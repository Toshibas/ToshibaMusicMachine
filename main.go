package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	bot, err := discordgo.New("Bot " + "OTc4NjkxNDk2MTkyNjU5NDU2.GwS474.IVRh60zMfVZqWSeL87ziEg_kxRhxm2tTYkGg1I")

	if err != nil {
		log.Fatal(err)
	}

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("ToshibaMusicMachine.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()

}
