package handlers

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

const ownerID = "337266897458429956"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == ownerID {
		if m.Content == "superstop" {
			fmt.Println(s.State.User.Username, "is shutting down...")
			s.Close()
			os.Exit(0)
			return
		}

		var allowedMentions = discordgo.MessageAllowedMentions{}
		var reference = discordgo.MessageReference{MessageID: m.ID, ChannelID: m.ChannelID, GuildID: m.GuildID}
		var messageSend = discordgo.MessageSend{Content: m.Content, Reference: &reference, AllowedMentions: &allowedMentions}
		s.ChannelMessageSendComplex(m.ChannelID, &messageSend)
	}
}
