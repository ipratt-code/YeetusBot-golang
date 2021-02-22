package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
)

func GetUserByID(ctx *commands.Context, id string) *discordgo.User{
	usrList, err := ctx.Session.GuildMembers(ctx.Message.GuildID, "", 1000)
	for _, usr := range usrList {
		if usr.User.ID == id {
			return usr.User
		}
	}
	//fmt.Printf("%v\n", usrList)
	if err != nil{
		fmt.Printf("%v\n", err)
	}
	return nil
	
}	