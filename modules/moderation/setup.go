package moderation

import (
	"fmt"
	"main/internal/commands"
	"main/internal/utils"
	"github.com/bwmarrin/discordgo"
	//"main/internal/cmdErrors"
	//"strings"
)



type Setup struct{}



func (c *Setup) Invokes() []string {
	return []string{"setup", "s", "se"}
}

func (c *Setup) Description() string {
	return "Sets up the discord server with muted roles and standard permissions"
}

func (c *Setup) AdminRequired() bool {
	return false
}

func (c *Setup) PermissionsRequired() (bool, uint) {
	return true, discordgo.PermissionAdministrator
}

func (c *Setup) Exec(ctx *commands.Context) (err error) {
	botID := "746507703635148901"
	botAsGuildMember := utils.GetMemberByID(ctx, botID)
	fmt.Printf("%+v\n", botAsGuildMember.Roles)
	//role perms for nothing is == 1024
	//fmt.Printf("%+v\n", mutedRole)
	return err
}