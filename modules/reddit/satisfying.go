package reddit

import (
	"main/internal/commands"
	//"fmt"
)

type Satisfying struct{}

func (c *Satisfying) Invokes() []string {
	return []string{"satisfying", "sa"}
}

func (c *Satisfying) Description() string {
	return "Satisfies you with stuff from r/satisfying and r/oddlysatisfying"
}

func (c *Satisfying) AdminRequired() bool {
	return false
}

func (c *Satisfying) PermissionsRequired() (bool, uint) {
	return false, 0
}


func (c *Satisfying) Exec(ctx *commands.Context) (err error) {
	post, err := redditRandomRetrieve([]string{"satisfying", "oddlysatisfying"}, 20)
	if err != nil {
		return err
	}
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, post.URL)
	return err
}
