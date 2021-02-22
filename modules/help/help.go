package help

import (
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
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

func strlistfmt(l []string) string {
	fmtstr := ""
	for _, item := range l {
		fmtstr = fmtstr + item + " | "
	}
	return fmtstr[0:len(fmtstr)-2]
}

// This is the help function to display the help message. Change it as you wish.
func (c *Help) Exec(ctx *commands.Context) error{
	msgEmb := &discordgo.MessageEmbed{}
	msgEmb.Title = "YeetusBot Help Panel"
	msgEmb.Color = 0xbad8eb
	for _, cmd := range ctx.Handler.CmdInstances {
		msgEmb.Fields = append(msgEmb.Fields, &discordgo.MessageEmbedField{
			Name: strlistfmt(cmd.Invokes()),
			Value: cmd.Description(),
		})
	}

	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, msgEmb)
	return err
}