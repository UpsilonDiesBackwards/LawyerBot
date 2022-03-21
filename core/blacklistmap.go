package core

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)

func Blacklist(s *discordgo.Session, m *discordgo.MessageCreate) {
	blacklistedPhrases := map[string]bool{
		"pedo":           false,
		"paedo":          false,
		"nonce":          false,
		"kiddie fiddler": false,
		"loli lover":     false,
		"reddit mod":     false,
		"discord mod":    false,
		"kitten":         false,
		"little girl":    false,
		"loli":           false,
		"kierbunny":      false,
		"anime":          false,
		"pervert":        false,
		"creep":          false,
		"daddy":          false,
		"child toucher":  false,
		"child":          false,
	}

	responses := make([]string, 0)
	responses = append(responses,
		"https://tenor.com/view/police-miss-kobayashi-dragon-maid-kanna-kamui-sit-here-gif-12872035",
		"Wait what?",
		"This doesn't look good for the court case",
		"This one here officer",
		"Yes officer, I would like to report a crime.",
		"What type of Jeffrey Epstein shit is that?",
		"The fuck?",
		"You're definitely going to jail",
		"I give up on you, you do ***not*** pay me enough for this",
	)

	for key, _ := range blacklistedPhrases {
		//fmt.Println("Key: ", key)

		rand.Seed(time.Now().Unix())
		if strings.Contains(m.Content, key) {
			_, err := s.ChannelMessageSend(m.ChannelID, responses[rand.Intn(len(responses))])
			if err != nil {
				println("\nFailed to send message response!\n")
				return
			}

		}
	}

	if strings.Contains(m.Content, "*test") {
		_, err := s.ChannelMessageSend(m.ChannelID, "The Lawyer is in court")
		if err != nil {
			println("\nFailed to send test message!\n")
			return
		}
	}
}
