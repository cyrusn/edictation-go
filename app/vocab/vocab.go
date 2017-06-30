package vocab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Vocab is ...
type Vocab struct {
	Title        string `json:"title,omitempty"`
	PartOfSpeech string `json:"partOfSpeech"`
	Definition   string `json:"definition"`
	Level        int    `json:"level"`
}

var vocabs []Vocab

func FilterVocabsByLevel(level int, vocabs []Vocab) []Vocab {
	return filter(vocabs, byLevel(level))
}

func GetVocabs(pathToJSON string) []Vocab {
	raw, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Vocab
	json.Unmarshal(raw, &c)
	return c
}

func byLevel(level int) func(Vocab) bool {
	return func(v Vocab) bool {
		if v.Level == level {
			return true
		}

		return false
	}
}

func filter(vocabs []Vocab, f func(Vocab) bool) []Vocab {
	var resultVocabs []Vocab

	for _, vocab := range vocabs {
		if f(vocab) {
			resultVocabs = append(resultVocabs, vocab)
		}
	}

	return resultVocabs
}
