package vocab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Vocab is basic struct for vocabulary.
type Vocab struct {
	Title        string `json:"title,omitempty"`
	PartOfSpeech string `json:"partOfSpeech"`
	Definition   string `json:"definition"`
}

var vocabs []Vocab

// FilterVocabsByIndexArray filter vocabs by index.
func FilterVocabsByIndexArray(indexes []int, vocabs []Vocab) []Vocab {
	var result []Vocab
	for _, index := range indexes {
	vocab_loop:
		for n, vocab := range vocabs {
			if n == index {
				result = append(result, vocab)
				break vocab_loop
			}
		}
	}
	return result
}

// GetVocabs will read vocabulary from location which is a json file that store all vocabularies.
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
