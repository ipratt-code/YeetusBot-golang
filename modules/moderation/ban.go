package moderation

import (
	"main/internal/commands"
	"main/internal/utils"
	"main/internal/cmdErrors"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strconv"
	"strings"
)

type Ban struct{}

func (c *Ban) Invokes() []string {
	return []string{"ban", "ba"}
}

func (c *Ban) Description() string {
	return "Bans a user from the discord server. Required arguments: User (mention a user), Length of time banned (pass `forever` to ban someone indefinitely. Time is in days. If over 7 days, gets corrected to 7 days)"
}

func (c *Ban) AdminRequired() bool {
	return true
}

func (c *Ban) Exec(ctx *commands.Context) (err error) {
	days := 7
	reasonList := []string{}
	var reason string
	print()
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User"})
	}else if ctx.Args[0] == "@everyone" {
		return fmt.Errorf("You can't run this command on everyone!")
	}else if len(ctx.Args) < 2 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User", "Length of time banned (in days. If over 7 days, gets corrected to 7 days)"})
	}else if len(ctx.Args) < 3 {
		reason = "No reason supplied"
	}else if len(ctx.Args) >= 2 {
		reasonList = append([]string{reason}, ctx.Args[2:]...)
		reason = strings.Join(reasonList, " ")
	}else if !(ctx.Args[1] == "forever"){
		days, err = strconv.Atoi(ctx.Args[1])
		if err != nil {
			return cmdErrors.BadArgumentsError([]string{"Length of time banned (in days. If over 7 days, gets corrected to 7 days)"})
		}
	}
	id := ctx.Args[0][3:len(ctx.Args[0])-1]
	usr := utils.GetUserByID(ctx, id)
	err = ctx.Session.GuildBanCreateWithReason(ctx.Message.GuildID, usr.ID, reason, days)
	
	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = fmt.Sprintf("%v#%v was banned!", usr.Username, usr.Discriminator)
	strdays := strconv.Itoa(days)
	if strdays == "7"{
		strdays = "i n f i n i t e"
	}
	msgEmb.Description = fmt.Sprintf("%v was banned for %s days from the server!\nReason: %s", usr.Mention(), strdays, reason)
	msgEmb.Color = 0xbad8eb
	_, _ = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)
	return err
}