package chatbot

import (
	"github.com/bwmarrin/discordgo"
	//"main/internal/commands"
	//"main/internal/cmdErrors"
	"main/internal/config"
	"fmt"
	"strings"
)

var cfg, _ = config.ParseConfigFromJSONFile("./config/config.json")
var prefix = cfg.Prefix
func chathandler(schan chan *discordgo.Session, echan chan *discordgo.MessageCreate) {
	fmt.Println("Chatbot Goroutine started!")
	var s *discordgo.Session
	var e *discordgo.MessageCreate

	for {
		s = <- schan
		e = <- echan
		print(e.Message.Author.Username + "\n")
		chatResponse := "response!"
		if strings.Contains(strings.ToLower(e.Message.Content), prefix + "chat stop") || strings.Contains(strings.ToLower(e.Message.Content), prefix + "ch stop") {
			//chatResponse = "Goodbye!"
			break
		}
		msgEmb := &discordgo.MessageEmbed{} 
		msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
		msgEmb.Description = fmt.Sprintf("YeetusBot: %s", chatResponse)
		msgEmb.Color = 0xbad8eb
		_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
	}
	fmt.Println("Chatbot Goroutine stopped!")
	msgEmb := &discordgo.MessageEmbed{} 
	msgEmb.Title = fmt.Sprintf("%s#%s's YeetusBot Chatbot chat!", e.Message.Author.Username, e.Message.Author.Discriminator)
	msgEmb.Description = "YeetusBot: Goodbye!"
	msgEmb.Color = 0xbad8eb
	_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)

	return
}