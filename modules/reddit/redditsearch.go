package reddit

import (
	"main/internal/commands"
	"strings"
	//"fmt"
)

type Redditsearch struct{}

func (c *Redditsearch) Invokes() []string {
	return []string{"redditsearch", "rs"}
}

func (c *Redditsearch) Description() string {
	return "Looks up a subreddit and takes a random hot post"
}

func (c *Redditsearch) AdminRequired() bool {
	return false
}

func (c *Redditsearch) PermissionsRequired() (bool, uint) {
	return false, 0
}


func (c *Redditsearch) Exec(ctx *commands.Context) (err error) {
	arg := strings.Join(ctx.Args, "+")
	post, err := redditRandomRetrieve([]string{arg}, 20)
	if err != nil {
		return err
	}
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, post.URL)
	return err
}
