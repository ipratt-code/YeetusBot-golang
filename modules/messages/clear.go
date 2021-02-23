package messages

import (
	"main/internal/commands"
	"main/internal/utils"
	"main/internal/cmdErrors"
	"strconv"
	"fmt"
)

type Clear struct {}


func (c *Clear) Invokes() []string {
	return []string{"clear", "c", "cl"}
}

func (c *Clear) Description() string {
	return "Clears last n messages. Required arguments: n (number of messages to clear. Maximum number of messages: 100)"
}

func (c *Clear) AdminRequired() bool {
	return true
}

func (c *Clear) Exec(ctx *commands.Context) (err error) {
	var msgIDStringList []string
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"Number of messsages to get"})
	}
	n, err := strconv.Atoi(ctx.Args[0])
	if err == nil && n == 0 {
		return fmt.Errorf("You cant clear 0 messages!")
	}else if n > 100 {
		return fmt.Errorf("You cant clear more than 100 messages!")
	}
	/*if err != nil {
		if ctx.Args[0] == "all" {
			n = 0
		}else {
			return err
		}
	}*/
	msgs, err := utils.GetChannelMessages(ctx, n)
	if err != nil {
		return err
	}
	for _, msg := range msgs {
		msgIDStringList = append(msgIDStringList, msg.ID)
	}
	err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, msgIDStringList)
	return err
}

