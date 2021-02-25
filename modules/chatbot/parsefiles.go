package chatbot

import (
	"encoding/json"
	"bufio"
	"os"
)

type sentances struct {
	SentanceList []string `json:"Sentances"`
}

func parseJSONFile(fileName string) (s *sentances, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	s = new(sentances)
	err = json.NewDecoder(f).Decode(s)

	return
}


func parseTXTFile(path string) (*sentances, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines) 

	var lines []string

	for scanner.Scan() {
	lines = append(lines, scanner.Text())
	}

	return &sentances{
		SentanceList: lines,
	}, nil
}
