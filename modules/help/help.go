package help

import (
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
	//"fmt"
)

type Help struct{}

func (c *Help) Invokes() []string {
	return []string{"help", "h", "he"}
}

func (c *Help) Description() string {
	return "Returns the YeetusBot help panel"
}

func (c *Help) AdminRequired() bool {
	return false
}

func (c *Help) PermissionsRequired() (bool, uint) {
	return false, 0
}


func strlistfmt(l []string) string {
	fmtstr := ""
	for _, item := range l {
		fmtstr = fmtstr + item + " | "
	}
	return fmtstr[0:len(fmtstr)-2]
}

// This is the help function to display the help message. Change it as you wish.
func (c *Help) Exec(ctx *commands.Context) error{
	newDM, _ := ctx.Session.UserChannelCreate(ctx.Message.Author.ID)
	
	helpEmb := &discordgo.MessageEmbed{}
	helpEmb.Title = "YeetusBot Help Panel"
	helpEmb.Color = 0xbad8eb
	for _, cmd := range ctx.Handler.CmdInstances {
		helpEmb.Fields = append(helpEmb.Fields, &discordgo.MessageEmbedField{
			Name: strlistfmt(cmd.Invokes()),
			Value: cmd.Description(),
		})
	}
	_, err := ctx.Session.ChannelMessageSendEmbed(newDM.ID, helpEmb)
	
	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = "Don't worry, help is on the way!"
	msgEmb.Color = 0xbad8eb
	msgEmb.Description = "Check your DMs, " +  ctx.Message.Author.Mention() + "!"
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)
	return err
}
