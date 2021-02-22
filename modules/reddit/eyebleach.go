package reddit

import (
	"main/internal/commands"
	//"fmt"
)

type Eyebleach struct{}

func (c *Eyebleach) Invokes() []string {
	return []string{"eyebleach", "e", "eb"}
}

func (c *Eyebleach) Description() string {
	return "Bleaches your eyes with photos from r/eyebleach and r/aww"
}

func (c *Eyebleach) AdminRequired() bool {
	return false
}

func (c *Eyebleach) Exec(ctx *commands.Context) (err error) {
	post, err := redditRandomRetrieve([]string{"eyebleach", "aww"}, 20)
	if err != nil {
		return err
	}
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, post.URL)
	return err
}
