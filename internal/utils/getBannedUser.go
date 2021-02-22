package utils

import (
	//"fmt"
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
)

func GetBannedUser(ctx *commands.Context, username, discriminator string) *discordgo.User {
	bans, _:= ctx.Session.GuildBans(ctx.Message.GuildID)

	for _, ban := range bans {
		if ban.User.Username == username {
			if ban.User.Discriminator == discriminator{
				return ban.User
			}
		}
	}

	return nil
}