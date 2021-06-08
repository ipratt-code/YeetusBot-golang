package moderation

import (
	"fmt"
	"main/internal/commands"
	"main/internal/utils"
	"github.com/bwmarrin/discordgo"
	"main/internal/cmdErrors"
	"strings"
)

type Unmute struct{}

func (c *Unmute) Invokes() []string {
	return []string{"Unmute", "um", "umu"}
}

func (c *Unmute) Description() string {
	return "Unmutes somone in text channels."
}

func (c *Unmute) AdminRequired() bool {
	return true
}

func (c *Unmute) PermissionsRequired() (bool, uint) {
	return true, discordgo.PermissionVoiceMuteMembers
}


func (c *Unmute) Exec(ctx *commands.Context) (err error) {
	defer utils.CatchGoroutinePanic()
	//fmt.Printf("%v\n", ctx.Args)
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
	id := utils.ParseIDFromMention(ctx.Args[0])
	err = ctx.Session.GuildMemberRoleRemove(ctx.Message.GuildID, id, mutedRole.ID)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	return err
}