package chatbot

import (
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
	"main/internal/cmdErrors"
	"fmt"
	"strings"
	"main/internal/utils"
)

var cmdchannel = make(chan string)

type Chatbot struct{
	OnMessageListeners []chan string
}

func (c *Chatbot) Invokes() []string {
	return []string{"chat", "ch"}
}

func (c *Chatbot) Description() string {
	return "Starts a chat bot for you to communicate with! Required arguments: [start|stop]"
}

func (c *Chatbot) AdminRequired() bool {
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
func (c *Chatbot) Exec(ctx *commands.Context) error {
	if (strings.ToLower(ctx.Args[0])) == "start" {
		cmdchannel <- "start " + ctx.Message.Author.ID
		return nil
	}else if (strings.ToLower(ctx.Args[0])) == "stop"{
		cmdchannel <- "stop " + ctx.Message.Author.ID
	}else{
		return cmdErrors.BadArgumentsError([]string{"argument [start|stop]"})
	}

	return nil
}

type chatbotprocess struct {
	schan chan *discordgo.Session
	echan chan *discordgo.MessageCreate
}

type ChatbotMessageHandler struct {
	schan chan *discordgo.Session
	echan chan *discordgo.MessageCreate
	processMap map[string]*chatbotprocess

}

func NewMessageHandler() *ChatbotMessageHandler {
	return &ChatbotMessageHandler{}
}

//var chatflag = false

func (h *ChatbotMessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if h.processMap == nil {
    	h.processMap = make(map[string]*chatbotprocess)
	}
	//fmt.Printf("%+v\n", h.processMap)
	//_, ok := h.processMap[e.Message.Author.ID]
	//fmt.Printf("%v\n", ok)
	if _, ok := h.processMap[e.Message.Author.ID]; ok && !e.Message.Author.Bot{
		//print("owo\n")
		//fmt.Printf("%+v", h.processMap[e.Message.Author.ID])
		h.processMap[e.Message.Author.ID].schan <- s
		h.processMap[e.Message.Author.ID].echan <- e
	}
	for msg := range cmdchannel {
		print(msg + "\n")
		msgList := strings.Split(msg, " ")
		if msgList[0] == "start"{
			if _, ok := h.processMap[msgList[1]]; !ok {
				//print("uwu\n")
				tschan := make(chan *discordgo.Session)
				techan := make(chan *discordgo.MessageCreate)
				h.processMap[msgList[1]] = &chatbotprocess{
					schan: tschan,
					echan: techan,
				}
				msgEmb := &discordgo.MessageEmbed{} 
				msgEmb.Title = "YeetusBot Chatbot"
				msgEmb.Description = fmt.Sprintf("%s, Your chat has started!", utils.GetUserByID(&commands.Context{Session: s, Message: e.Message,}, msgList[1]).Mention())
				msgEmb.Color = 0xbad8eb
				//schan = make(chan *discordgo.Session)
				//echan = make(chan *discordgo.MessageCreate)
				_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
				go chathandler(tschan, techan)
			}else if ok {

			}
			//chatflag = true
			
		}else if msgList[0] == "stop" {
			if _, ok := h.processMap[msgList[1]]; ok {
				print("a\n")
				msgEmb := &discordgo.MessageEmbed{} 
				msgEmb.Title = "YeetusBot Chatbot"
				msgEmb.Description = fmt.Sprintf("%s, Your chat has stopped!", utils.GetUserByID(&commands.Context{Session: s, Message: e.Message,}, msgList[1]).Mention())
				msgEmb.Color = 0xbad8eb
				_, _ = s.ChannelMessageSendEmbed(e.Message.ChannelID, msgEmb)
				delete(h.processMap, msgList[1])
			}
		}
	}
}
