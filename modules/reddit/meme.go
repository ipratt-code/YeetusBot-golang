package reddit

import (
	"main/internal/commands"
	//"fmt"
)

type Meme struct{}

func (c *Meme) Invokes() []string {
	return []string{"meme", "m", "me"}
}

func (c *Meme) Description() string {
	return "Returns a meme from r/dankmemes or r/memes"
}

func (c *Meme) AdminRequired() bool {
	return false
}

func (c *Meme) Exec(ctx *commands.Context) (err error) {
	post, err := redditRandomRetrieve([]string{"dankmemes", "memes"}, 20)
	if err != nil {
		return err
	}
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, post.URL)
	return err
}
