package utils

import (
	//"fmt"
	//"github.com/bwmarrin/discordgo"
	//"main/internal/commands"
	"strings"
)

func ParseIDFromMention(mention string) string {
	if strings.Contains(mention, "<@!") {
		return mention[3:len(mention)-1]
	}else {
		return mention[2:len(mention)-1]
	}
	
}	