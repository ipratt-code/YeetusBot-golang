package moderation

import (
	"fmt"
	"main/internal/commands"
	"main/internal/utils"
	"github.com/bwmarrin/discordgo"
	"main/internal/cmdErrors"
	"strings"
)

type Kick struct{}

func (c *Kick) Invokes() []string {
	return []string{"kick", "k", "ki"}
}

func (c *Kick) Description() string {
	return "Kick a user from the discord server. Required arguments: user (mention a user). Optional arguments: reason"
}

func (c *Kick) AdminRequired() bool {
	return true
}

func (c *Kick) PermissionsRequired() (bool, uint) {
	return true, discordgo.PermissionKickMembers
}

func (c *Kick) Exec(ctx *commands.Context) (err error) {
	reasonList := []string{}
	var reason string
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User"})
	}else if ctx.Args[0] == "@everyone" {
		return fmt.Errorf("You can't run this command on everyone!")
	}else if len(ctx.Args) < 2 {
		reason = "No reason supplied"
	}else if len(ctx.Args) >= 2 {
		reasonList = append([]string{reason}, ctx.Args[1:]...)
		reason = strings.Join(reasonList, " ")
	}else if ctx.Args[0] == "@everyone" {
		return fmt.Errorf("You can't run this command on everyone!")
	}
	
	id := utils.ParseIDFromMention(ctx.Args[0])
	usr := utils.GetUserByID(ctx, id)
	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = fmt.Sprintf("%v#%v was kicked!", usr.Username, usr.Discriminator)
	msgEmb.Description = fmt.Sprintf("%v was kicked from the server!\nReason: %v", usr.Mention(), reason)
	msgEmb.Color = 0xbad8eb

	err = ctx.Session.GuildMemberDeleteWithReason(ctx.Message.GuildID, usr.ID, reason)
	_, _ = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)
	return err
	
}