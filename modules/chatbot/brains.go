package chatbot

import (
	"math/rand"
	"github.com/bbalet/stopwords"
	"github.com/recoilme/tfidf"
	"github.com/recoilme/tfidf/similarity"
	//"fmt"
	"strings"
)

type NSW2RegMap struct{
	NSW2Reg map[string]string
}

var stcs, _ = parseTXTFile("./modules/chatbot/corpus/sentances.txt")

func newNSW2RegMap() *NSW2RegMap {
	return &NSW2RegMap{
		NSW2Reg: make(map[string]string),
	}
}

func (n *NSW2RegMap) initNSW2Reg() {
	for _, st := range stcs.SentanceList {
		n.NSW2Reg[strings.ToLower(stopwords.CleanString(st, "en", true))] = st 
	}
	return
}

func getRandSentance() string {
	st := stcs.SentanceList[rand.Intn(len(stcs.SentanceList))]
	//fmt.Printf("%+v\n", st)
	return st
}

func getCosineSimilaritySentance(s string, n *NSW2RegMap) string {
	s = strings.ToLower(s)
	f := tfidf.New()
	for k := range n.NSW2Reg {
		f.AddDocs(k)
	}
	snsw := stopwords.CleanString(s, "en", true) //input w/ no stopwords
	highesti := 0.0
	highests := ""
	w1 := f.Cal(snsw)
	for k := range n.NSW2Reg {
		sim := similarity.Cosine(w1, f.Cal(k))
		if sim > highesti {
			highesti = sim
			highests = n.NSW2Reg[k]
		}
	}
	return highests
}