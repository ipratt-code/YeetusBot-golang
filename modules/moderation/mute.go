package moderation

import (
	"fmt"
	"main/internal/commands"
	"main/internal/utils"
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
	return false
}

func (c *Mute) PermissionsRequired() (bool, uint) {
	return true, discordgo.PermissionVoiceMuteMembers
}

func (c *Mute) Exec(ctx *commands.Context) (err error) {
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
	//print("1\n")
	if mutedRole == nil {
		return fmt.Errorf("You need to create a @muted role! (make sure the role is lower than the highest role this bot has)")
	}
	if len(ctx.Args) < 1 {
		return cmdErrors.NeedRequiredArgumentsError([]string{"User (mention a user)"})
	}
	//print("2\n")
	id := utils.ParseIDFromMention(ctx.Args[0])
	err = ctx.Session.GuildMemberRoleAdd(ctx.Message.GuildID, id, mutedRole.ID)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	return err
}