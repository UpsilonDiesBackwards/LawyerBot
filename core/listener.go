package core

import (
	"github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore self
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "https://tenor.com/view/police-miss-kobayashi-dragon-maid-kanna-kamui-sit-here-gif-12872035" {
		return
	}

	if strings.Contains(m.Content, ">help") {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbedAdvanced("Lawyer Bot Help Page",
			">test :: Test to see if the bot is working\n" +
			">vc join on/off :: Allow the Lawyer to randomly join a random voice channel in the server", 2))
	}

	if strings.Contains(m.Content, ">test") {
		s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/police-miss-kobayashi-dragon-maid-kanna-kamui-sit-here-gif-12872035")
	}

	if strings.Contains(m.Content, ">vc join on") {
		s.ChannelMessageSend(m.ChannelID, "Lawyer bot can now randomly join Voice Channels")
		joinVC = true
	}

	if strings.Contains(m.Content, ">vc join off") {
		s.ChannelMessageSend(m.ChannelID, "Lawyer bot can no longer randomly join Voice Channels")
		joinVC = false
	}
}
