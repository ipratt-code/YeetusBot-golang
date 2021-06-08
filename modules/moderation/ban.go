package moderation

import (
	"main/internal/commands"
	"main/internal/utils"
	"main/internal/cmdErrors"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
)

type Ban struct{}

func (c *Ban) Invokes() []string {
	return []string{"ban", "ba"}
}

func (c *Ban) Description() string {
	return "Bans a user from the discord server. Required arguments: User (mention a user)"
}

func (c *Ban) AdminRequired() bool {
	return true
}

func (c *Ban) PermissionsRequired() (bool, uint) {
	return true, discordgo.PermissionBanMembers
}


func (c *Ban) Exec(ctx *commands.Context) (err error) {
	reasonList := []string{}
	var reason string
	print()
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User"})
	}else if ctx.Args[0] == "@everyone" {
		return fmt.Errorf("You can't run this command on everyone!")
	}else if len(ctx.Args) < 2 {
		reason = "No reason supplied"
	}else if len(ctx.Args) >= 2 {
		reasonList = append([]string{reason}, ctx.Args[2:]...)
		reason = strings.Join(reasonList, " ")
	}
	id := utils.ParseIDFromMention(ctx.Args[0])
	usr := utils.GetUserByID(ctx, id)
	err = ctx.Session.GuildBanCreateWithReason(ctx.Message.GuildID, usr.ID, reason, 7)
	
	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = fmt.Sprintf("%v#%v was banned!", usr.Username, usr.Discriminator)
	msgEmb.Description = fmt.Sprintf("%v was banned from the server!\nReason: %s", usr.Mention(), reason)
	msgEmb.Color = 0xbad8eb
	_, _ = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)
	return err
}
