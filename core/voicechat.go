package core

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

var joinVC = false

func Joinvoicechat(s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceChannels := make([]*discordgo.Channel, 0)

	for _, guild := range s.State.Guilds {
		channels, _ := s.GuildChannels(guild.ID)

		for _, c := range channels {
			if c.Type != discordgo.ChannelTypeGuildVoice {
				continue
			}

			voiceChannels = append(voiceChannels, c)
		}
	}

	if joinVC == true {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(5) // n = max wait time in minutes
		fmt.Printf("\nSleeping %d minute(s)...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Println("Done")

		vc := voiceChannels[rand.Intn(len(voiceChannels))]
		v, _ := s.ChannelVoiceJoin(vc.GuildID, vc.ID, true, false)

		time.Sleep(time.Duration(2) * time.Hour)
		v, _ = s.ChannelVoiceJoin(vc.GuildID, vc.ID, false, false)

		dgvoice.PlayAudioFile(v, "media/laugh.wav", make(chan bool))

		time.Sleep(time.Duration(20) * time.Second)
		v, _ = s.ChannelVoiceJoin(vc.GuildID, vc.ID, true, false)


		//fmt.Println("Server: ", vc.GuildID, "Channel: ", vc.Name)
		//
		////fmt.Println("\n My Balls:", string(voiceChannels[rand.Intn(len(voiceChannels))])
		////fmt.Println("\nServer Name: ", guild.Name,
		////	"\nChannel Name: ", c.Name,
		////	"\nChannel ID: ", c.ID)
		////
		////if err != nil {
		////	fmt.Printf("\nConnection Status: Failed to connect to voice channel: %s\n", err)
		////	return
		////} else {
		////	fmt.Println("\n\n\nConnected! to Server Name: ", guild.Name,
		////		"\nChannel Name: ", c.Name,
		////		"\nChannel ID: ", c.ID )
		////	return
		////}
	}
}
