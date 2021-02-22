package main
import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
	"main/internal/config"
	"main/internal/events"
	"main/internal/cmdErrors"
	
	//these are the modules
	"main/modules/help"
	"main/modules/pingpong"
	"main/modules/moderation"
	"main/modules/messages"
	"main/modules/reddit"
)

type ExportedCmdHandler struct {
	cmdHandler *commands.CommandHandler
}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
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
	cmdHandler.RegisterCommand(&messages.Clear{})
	cmdHandler.RegisterCommand(&reddit.Meme{})
	cmdHandler.RegisterCommand(&reddit.Blursedimages{})
	cmdHandler.RegisterCommand(&reddit.Redditsearch{})
	cmdHandler.RegisterCommand(&reddit.Eyebleach{})
	cmdHandler.RegisterCommand(&reddit.Satisfying{})
	cmdHandler.RegisterCommand(&pingpong.PingPong{})

	// register the middleware to execute commands
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}