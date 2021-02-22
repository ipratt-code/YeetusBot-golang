package pingpong

import (
	"main/internal/commands"
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
