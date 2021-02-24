package chatbot

import (
	"github.com/bwmarrin/discordgo"
	//"main/internal/commands"
	//"main/internal/cmdErrors"
	"main/internal/config"
	"main/internal/utils"
	"fmt"
	"strings"
	"context"
)

var cfg, _ = config.ParseConfigFromJSONFile("./config/config.json")
var prefix = cfg.Prefix
func chathandler(schan chan *discordgo.Session, echan chan *discordgo.MessageCreate, pchan chan bool, timeoutctx context.Context) {
	defer utils.CatchGoroutinePanic()
	fmt.Println("Chatbot Goroutine started!")
	var s *discordgo.Session
	var e *discordgo.MessageCreate
	flg := false
	for {
		if flg == true {
			break
		}
		select{
		case <-timeoutctx.Done():
			fmt.Println("Chatbot Goroutine timed out!")
			msgEmb := &discordgo.MessageEmbed{} 
			msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
			msgEmb.Description = "YeetusBot: Goodbye!"
			msgEmb.Color = 0xbad8eb
			_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
			msgEmb = &discordgo.MessageEmbed{} 
			msgEmb.Title = "YeetusBot Chatbot"
			msgEmb.Description = fmt.Sprintf("%s, Your chat has timed out!", e.Message.Author.Mention())
			msgEmb.Color = 0xbad8eb
			_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
			return
		case <- pchan:
			s = <- schan
			e = <- echan
			chatResponse := e.Message.Content
			if strings.Contains(strings.ToLower(e.Message.Content), prefix + "chat stop") || strings.Contains(strings.ToLower(e.Message.Content), prefix + "ch stop") {
				flg = true
				break
			}
			msgEmb := &discordgo.MessageEmbed{} 
			msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
			msgEmb.Description = fmt.Sprintf("YeetusBot: Hello, and %s to you too!", chatResponse)
			msgEmb.Color = 0xbad8eb
			_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
		
		}
	}
	fmt.Println("Chatbot Goroutine stopped!")
	msgEmb := &discordgo.MessageEmbed{} 
	msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
	msgEmb.Description = "YeetusBot: Goodbye!"
	msgEmb.Color = 0xbad8eb
	_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)

	return
}