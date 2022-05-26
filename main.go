package main

import (
	"fmt"
	"os"

	"GoInitBot/handlers"
	"github.com/bwmarrin/discordgo"
)

var token = os.Getenv("DISCORD_TOKEN")

func main() {
	var session, initErr = discordgo.New("Bot " + token)
	if initErr != nil {
		fmt.Println(initErr.Error())
		return
	}

	session.AddHandler(handlers.Ready)
	session.AddHandler(handlers.MessageCreate)

	var conErr = session.Open()
	if conErr != nil {
		fmt.Println(conErr.Error())
		return
	}
	select {}
}
