package moderation

import(
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

type server struct {
	ID string
	NotificationChannelID string
	MutedRole *discordgo.Role
}

type servers struct {
	serverMap map[string]*server
}


func parseConfigFromJSONFile(fileName string) (c *servers, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(servers)
	err = json.NewDecoder(f).Decode(c)

	return
}

func writeDataToJSONFile(fileName string) (c *servers, err error) {
	file, _ := json.MarshalIndent(c, "", " ")
 
	_ = ioutil.WriteFile("modules/moderation/persist/data.json", file, 0644)
	return
}