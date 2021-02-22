package reddit

import (
	"main/internal/commands"
	//"fmt"
)

type Blursedimages struct{}

func (c *Blursedimages) Invokes() []string {
	return []string{"eyebleach", "e", "eb"}
}

func (c *Blursedimages) Description() string {
	return "Gets a blursed image from r/blursedimages"
}

func (c *Blursedimages) AdminRequired() bool {
	return false
}

func (c *Blursedimages) Exec(ctx *commands.Context) (err error) {
	post, err := redditRandomRetrieve([]string{"blursedimages"}, 20)
	if err != nil {
		return err
	}
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, post.URL)
	return err
}
