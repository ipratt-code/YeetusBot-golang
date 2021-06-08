package nou

import (
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
	"strings"
	//"fmt"
)

type Nou struct{}

func (c *Nou) Invokes() []string {
	return []string{"nou", "nou"}
}

func (c *Nou) Description() string {
	return "no u"
}

func (c *Nou) AdminRequired() bool {
	return false
}

func (c *Nou) PermissionsRequired() (bool, uint) {
	return false, 0
}

func (c *Nou) Exec(ctx *commands.Context) (err error) {
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	return err
}

type nouMessageHandler struct{}

func NewMessageHandler() *nouMessageHandler {
	return &nouMessageHandler{}
}

func (h *nouMessageHandler) Check(m *discordgo.Message) bool {
	if strings.Contains(strings.ToLower(m.Content), "nou") {
		return true
	}
	return false
}

func (h *nouMessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if (strings.Contains(strings.ToLower(e.Message.Content), "no u") || strings.Contains(strings.ToLower(e.Message.Content), "nou") || strings.Contains(e.Message.Content, "ňØ Ū")) && !e.Message.Author.Bot {
		_, _ = s.ChannelMessageSend(e.Message.ChannelID, "no u")
	}
}
