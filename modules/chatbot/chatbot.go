package chatbot

import (
	"github.com/bwmarrin/discordgo"
	"main/internal/commands"
	"main/internal/cmdErrors"
	"fmt"
	"strings"
	"main/internal/utils"
	"time"
	"context"
)

var cmdchannel = make(chan string)

type Chatbot struct{
	OnMessageListeners []chan string
}

func (c *Chatbot) Invokes() []string {
	return []string{"chat", "ch"}
}

func (c *Chatbot) Description() string {
	return "Starts a chat bot for you to communicate with! Required arguments: [start|stop]. ONLY LASTS 5 MINUTES. COMMAND LIABLE TO BE REMOVED."
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
		cmdchannel <- "start " + ctx.Message.Author.ID + " " + ctx.Message.ChannelID
		return nil
	}else if (strings.ToLower(ctx.Args[0])) == "stop"{
		cmdchannel <- "stop " + ctx.Message.Author.ID + " " + ctx.Message.ChannelID
	}else{
		return cmdErrors.BadArgumentsError([]string{"argument [start|stop]"})
	}

	return nil
}

type chatbotprocess struct {
	timeoutctx context.Context
	cancelFunc context.CancelFunc
	session *discordgo.Session
	echan chan *discordgo.MessageCreate
	pchan chan bool
	author *discordgo.User

}

type ChatbotMessageHandler struct {
	processMap map[string]*chatbotprocess

}

func NewMessageHandler() *ChatbotMessageHandler {
	return &ChatbotMessageHandler{}
}

//var chatflag = false

func (h *ChatbotMessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
    defer utils.CatchGoroutinePanic()
	if h.processMap == nil {
    	h.processMap = make(map[string]*chatbotprocess)
	}
	if _, ok := h.processMap[e.Message.Author.ID]; ok && !e.Message.Author.Bot{
		h.processMap[e.Message.Author.ID].pchan <- true
		h.processMap[e.Message.Author.ID].echan <- e
	}
	for msg := range cmdchannel {
		print(msg + "\n")
		msgList := strings.Split(msg, " ")
		if msgList[0] == "start"{
			if _, ok := h.processMap[msgList[1]]; !ok {
				//print("uwu\n")
				techan := make(chan *discordgo.MessageCreate)
				tpchan := make(chan bool)
				timeoutctx, cancelf := context.WithTimeout(context.Background(), 5*time.Minute)
				h.processMap[msgList[1]] = &chatbotprocess{
					timeoutctx: timeoutctx,
					cancelFunc: cancelf,
					session: s,
					echan: techan,
					pchan: tpchan,
					author: utils.GetUserByID(&commands.Context{Session: s, Message: e.Message,}, msgList[1]),
				}
				msgEmb := &discordgo.MessageEmbed{} 
				msgEmb.Title = "YeetusBot Chatbot"
				msgEmb.Description = fmt.Sprintf("%s, Your chat has started!", utils.GetUserByID(&commands.Context{Session: s, Message: e.Message,}, msgList[1]).Mention())
				msgEmb.Color = 0xbad8eb
				_, _ = s.ChannelMessageSendEmbed(msgList[2], msgEmb)
				go chathandler(h.processMap[msgList[1]])
				go handleInstanceTimeouts(h, h.processMap[msgList[1]])
			}else if ok {

			}
			//chatflag = true
			
		}else if msgList[0] == "stop" {
			if _, ok := h.processMap[msgList[1]]; ok {
				msgEmb := &discordgo.MessageEmbed{} 
				msgEmb.Title = "YeetusBot Chatbot"
				msgEmb.Description = fmt.Sprintf("%s, Your chat has stopped!", utils.GetUserByID(&commands.Context{Session: s, Message: e.Message,}, msgList[1]).Mention())
				msgEmb.Color = 0xbad8eb
				_, _ = s.ChannelMessageSendEmbed(msgList[2], msgEmb)
				delete(h.processMap, msgList[1])
			}
		}
	}
}

func handleInstanceTimeouts(h *ChatbotMessageHandler, c *chatbotprocess) {
	defer utils.CatchGoroutinePanic()
	select {
	case <- c.timeoutctx.Done():
		//fmt.Println("aaa")
		delete(h.processMap, c.author.ID)
		time.Sleep(10*time.Second)
		fmt.Println("Chatbot Process deleted")
	}
}