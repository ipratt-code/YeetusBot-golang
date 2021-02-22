package utils

import (
	//"fmt"
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
)

func GetChannelMessages(ctx *commands.Context, n int) ([]*discordgo.Message, error) {
	var chanMsgs []*discordgo.Message

//	msgs, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, n, "", "", "")
//	return msgs, err
	/*if n == 0 {
		iter := 0
		l, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, n, "", "", "")
		for len(l) > 0{			
			iter += 1
			if err != nil {
				return chanMsgs, err
			}
			chanMsgs = append(chanMsgs, l...)
			l, err = ctx.Session.ChannelMessages(ctx.Message.ChannelID, n, "", l[len(l)-1].ID, "")
		}
	}*/
	/*for n > 100 {
		msgs, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, 100, "", "", "")
		if err != nil {
			return chanMsgs, err
		}
		chanMsgs = append(chanMsgs, msgs...)
		n = n - 100
	}*/
	msgs, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, n, "", "", "")
	chanMsgs = append(chanMsgs, msgs...)
	return chanMsgs, err
}/*
	for n > 100{
		msgs, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, 100, "", "", "")
		if err != nil{
			return chanMsgs, err
		}
		chanMsgs = append(chanMsgs, msgs...)
		
		n = n - 100 
		
		if n < 100 && n > 0 {
			msgs, err = ctx.Session.ChannelMessages(ctx.Message.ChannelID, n, "", "", "")
			if err != nil{
				return chanMsgs, err
			}
			chanMsgs = append(chanMsgs, msgs...)
		}
	}
	return chanMsgs, nil
}*/