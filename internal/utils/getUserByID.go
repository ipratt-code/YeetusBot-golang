package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
)

func GetUserByID(ctx *commands.Context, id string) *discordgo.User{
	afterID := ""
	usrList, err := ctx.Session.GuildMembers(ctx.Message.GuildID, afterID, 1000)
	for idx, usr := range usrList {
		//fmt.Printf("%v\n", idx)
		if usr.User.ID == id {
			return usr.User
		}else if idx >= 999 {
			afterID = usr.User.ID
		}
	}
	//fmt.Printf("%v\n", usrList)
	if err != nil{
		fmt.Printf("%v\n", err)
	}

	for {
		usrList, err = ctx.Session.GuildMembers(ctx.Message.GuildID, afterID, 1000)
		for idx, usr := range usrList {
			//fmt.Printf("%v\n", idx)
			if usr.User.ID == id {
				return usr.User
			}else if idx >= 999 {
				afterID = usr.User.ID
			}
		}
		if len(usrList) < 1000 {
			return nil
		}
	}

	return nil
}	