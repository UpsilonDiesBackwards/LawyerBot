package core

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)



func Blacklist(s *discordgo.Session, m *discordgo.MessageCreate) {
	blacklistedPhrases := map[string]bool {
		"pedo": false,
		"paedo": false,
		"nonce": false,
		"kiddie fiddler": false,
		"loli lover": false,
		"reddit mod": false,
		"discord mod": false,
		"kitten": false,
		"little girl": false,
		"loli": false,
		"kierbunny": false,
		"anime": false,
		"pervert": false,
		"creep": false,
		"daddy": false,
		"uwu": false,
		"owo": false,
		"child toucher": false,
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
	)

	for key, _ := range blacklistedPhrases {
		fmt.Println("Key: ", key)

		rand.Seed(time.Now().Unix())
		if strings.Contains(m.Content, key) {
			s.ChannelMessageSend(m.ChannelID, responses[rand.Intn(len(responses))])

		}
	}

	if strings.Contains(m.Content, "*test") {
		s.ChannelMessageSend(m.ChannelID, "The Lawyer is in court")
	}
}