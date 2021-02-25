package chatbot

import (
	"github.com/bwmarrin/discordgo"
	//"main/internal/commands"
	//"main/internal/cmdErrors"
	"main/internal/config"
	"main/internal/utils"
	"fmt"
	"strings"
	//"context"
)

var cfg, _ = config.ParseConfigFromJSONFile("./config/config.json")
var prefix = cfg.Prefix

func chathandler(cp *chatbotprocess) {
	defer utils.CatchGoroutinePanic()
	fmt.Println("Chatbot Goroutine started!")
	echan := cp.echan
	s := cp.session
	timeoutctx := cp.timeoutctx
	pchan := cp.pchan
	var e *discordgo.MessageCreate
	nsw2r := newNSW2RegMap()
	nsw2r.initNSW2Reg()
	flg := false
	for {
		if flg == true {
			break
		}
		select{
		case <-timeoutctx.Done():
			fmt.Println("Chatbot Goroutine timed out!")
			newDM, _ := s.UserChannelCreate(cp.author.ID)
			msgEmb := &discordgo.MessageEmbed{} 
			msgEmb.Title = "YeetusBot Chatbot"
			msgEmb.Description = fmt.Sprintf("%s, Your chat has timed out!", cp.author.Mention())
			msgEmb.Color = 0xbad8eb
			_, _ = s.ChannelMessageSendEmbed(newDM.ID, msgEmb)
			return
		case <- pchan:
			e = <- echan
			if strings.Contains(strings.ToLower(e.Message.Content), prefix + "chat stop") || strings.Contains(strings.ToLower(e.Message.Content), prefix + "ch stop") {
				flg = true
				break
			}
			chatResponse := getCosineSimilaritySentance(e.Message.Content, nsw2r)
			msgEmb := &discordgo.MessageEmbed{} 
			msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
			msgEmb.Description = fmt.Sprintf("YeetusBot: %s", chatResponse)
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