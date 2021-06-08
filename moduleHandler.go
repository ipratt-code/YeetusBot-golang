package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"main/internal/cmdErrors"
	"main/internal/commands"
	"main/internal/config"
	"main/internal/events"

	//these are the modules
	"main/modules/chatbot"
	"main/modules/help"
	"main/modules/messages"
	"main/modules/moderation"
	"main/modules/pingpong"
	"main/modules/nou"
	"main/modules/reddit"
)

type ExportedCmdHandler struct {
	cmdHandler *commands.CommandHandler
}

func registerEvents(s *discordgo.Session) {

	s.AddHandler(events.NewReadyHandler().Handler)

	// add command module event listeners
	s.AddHandler(chatbot.NewMessageHandler().Handler)
	s.AddHandler(pingpong.NewMessageHandler().Handler)
	s.AddHandler(nou.NewMessageHandler().Handler)
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)

	// func to register the errors
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		errEmb := &discordgo.MessageEmbed{}
		errEmb.Title = "An Error Occurred!"
		errEmb.Description = fmt.Sprintf("%v", cmdErrors.FriendlyErrors(err).Error())
		errEmb.Color = 0xbad8eb

		_, _ = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID,
			errEmb)

	}

	// registering the commands so they are active
	// preferably register them in order of priority
	cmdHandler.RegisterCommand(&help.Help{})
	cmdHandler.RegisterCommand(&moderation.Kick{})
	cmdHandler.RegisterCommand(&moderation.Ban{})
	cmdHandler.RegisterCommand(&moderation.Unban{})
	cmdHandler.RegisterCommand(&moderation.Mute{})
	cmdHandler.RegisterCommand(&moderation.Unmute{})
	//cmdHandler.RegisterCommand(&moderation.Setup{})
	cmdHandler.RegisterCommand(&messages.Clear{})
	cmdHandler.RegisterCommand(&chatbot.Chatbot{})
	cmdHandler.RegisterCommand(&reddit.Meme{})
	cmdHandler.RegisterCommand(&reddit.Blursedimages{})
	cmdHandler.RegisterCommand(&reddit.Redditsearch{})
	cmdHandler.RegisterCommand(&reddit.Eyebleach{})
	cmdHandler.RegisterCommand(&reddit.Satisfying{})
	cmdHandler.RegisterCommand(&nou.Nou{})
	cmdHandler.RegisterCommand(&pingpong.PingPong{})

	// register the middleware to execute commands
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
