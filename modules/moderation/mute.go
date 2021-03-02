package moderation

import (
	"fmt"
	"main/internal/commands"
	//"main/internal/utils"
	"github.com/bwmarrin/discordgo"
	"main/internal/cmdErrors"
	"strings"
)

type Mute struct{}

func (c *Mute) Invokes() []string {
	return []string{"mute", "mu"}
}

func (c *Mute) Description() string {
	return "Mutes somone in text channels."
}

func (c *Mute) AdminRequired() bool {
	return true
}

func (c *Mute) Exec(ctx *commands.Context) (err error) {
	roles, err := ctx.Session.GuildRoles(ctx.Message.GuildID)
	if err != nil {
		return err
	}
	var mutedRole *discordgo.Role
	for _, role := range roles {
		if strings.Contains(strings.ToLower(role.Name), "muted"){
			mutedRole = role
			//fmt.Printf("%v\n", mutedRole)
		}
	}
	if mutedRole == nil {
		return fmt.Errorf("You need to create a @muted role! (make sure the role is lower than the highest role this bot has)")
	}
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User (mention a user)"})
	}
	id := ctx.Args[0][3:len(ctx.Args[0])-1]
	err = ctx.Session.GuildMemberRoleAdd(ctx.Message.GuildID, id, mutedRole.ID)
	//fmt.Printf("%v\n", err.Error())
	return err
}