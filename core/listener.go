package core

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore self
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "*test") {
		s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/police-miss-kobayashi-dragon-maid-kanna-kamui-sit-here-gif-12872035")
	}

	if strings.Contains(m.Content, "pedo") {
		s.ChannelMessageSend(m.ChannelID, "they're a WHAT?!")
	}
}
