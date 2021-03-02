package moderation

import (
	"main/internal/commands"
	"main/internal/utils"
	"github.com/bwmarrin/discordgo"
	"main/internal/cmdErrors"
	"strings"
	"fmt"
)

type Unban struct{}

func (c *Unban) Invokes() []string {
	return []string{"unban", "ub", "uba"}
}

func (c *Unban) Description() string {
	return "Bans a user from the discord server. Required arguments: User (format <username>#<discriminator>), Length of time banned (pass `forever` to ban someone indefinitely)"
}

func (c *Unban) AdminRequired() bool {
	return true
}

func (c *Unban) Exec(ctx *commands.Context) (err error) {
	usrFull := ""
	reasonList := []string{}
	var reason string
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User"})
	}else if len(ctx.Args) >= 1{
		indx := 0
		for idx, str := range ctx.Args {
			if strings.ContainsAny(str, "#") {
				usrFull = strings.Join(ctx.Args[:idx + 1], " ")
				indx = idx
			}

		}
		reasonList := append(reasonList, ctx.Args[indx + 1:]...)
		reason = strings.Join(reasonList, " ")
	}
	userSplit := strings.Split(usrFull, "#")
	//fmt.Printf("%v\n%v\n%v\n", ctx.Args, usrFull, userSplit)
	if len(userSplit) < 2 || len(userSplit) > 3 {
		return cmdErrors.BadArgumentsError([]string{"User (<username>#<discriminator>)"})
	}
	usr := utils.GetBannedUser(ctx, userSplit[0], userSplit[1])
	if usr == nil {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "That user is not banned!")
		return
	}
	err = ctx.Session.GuildBanDelete(ctx.Message.GuildID, usr.ID)

	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = fmt.Sprintf("%v#%v was unbanned!", usr.Username, usr.Discriminator)
	msgEmb.Description = fmt.Sprintf("%v was unbanned from the server!\nReason: %s", usr.Mention(), reason)
	msgEmb.Color = 0xbad8eb
	_, _ = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)


	return err
}