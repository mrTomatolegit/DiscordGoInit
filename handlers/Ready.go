package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	fmt.Println(s.State.User.Username, "is ready!")
}
