package pingpong

import (
	"main/internal/commands"
	"github.com/bwmarrin/discordgo"
	"strings"
	//"fmt"
)

type PingPong struct{}

func (c *PingPong) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *PingPong) Description() string {
	return "Pong!"
}

func (c *PingPong) AdminRequired() bool {
	return false
}

func (c *PingPong) Exec(ctx *commands.Context) (err error) {
	print()
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	return err
}


type PingPongMessageHandler struct{}

func NewMessageHandler() *PingPongMessageHandler {
	return &PingPongMessageHandler{}
}

func (h *PingPongMessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if strings.Contains(strings.ToLower(e.Message.Content), "ping") {
		_, _ = s.ChannelMessageSend(e.Message.ChannelID, "Pong!")
	}
}
