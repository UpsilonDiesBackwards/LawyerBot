package core

import "github.com/bwmarrin/discordgo"

var AppContext *Context

type Context struct {
	Discord *discordgo.Session
	Config  *Config
}
