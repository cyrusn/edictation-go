package vocab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Vocab is basic struct for vocabulary.
type Vocab struct {
	ID           int    `json:"id"`
	Title        string `json:"title,omitempty"`
	PartOfSpeech string `json:"partOfSpeech"`
	Definition   string `json:"definition"`
	Level        int    `json:"level"`
}

var vocabs []Vocab

// FilterVocabsByIDs filter vocabs by id.
func FilterVocabsByIDs(ids []int, vocabs []Vocab) []Vocab {
	var result []Vocab
	for _, id := range ids {
	vocab_loop:
		for _, vocab := range vocabs {
			if vocab.ID == id {
				result = append(result, vocab)
				break vocab_loop
			}
		}
	}

	return result
}

// FilterVocabsByLevel filter vocabs by level.
func FilterVocabsByLevel(level int, vocabs []Vocab) []Vocab {
	return filter(vocabs, byLevel(level))
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

// byLevel is callback function for vocabs by given level.
func byLevel(level int) func(Vocab) bool {
	return func(v Vocab) bool {
		if v.Level == level {
			return true
		}

		return false
	}
}

// filter is helper function to filter vocabs.
func filter(vocabs []Vocab, f func(Vocab) bool) []Vocab {
	var resultVocabs []Vocab

	for _, vocab := range vocabs {
		if f(vocab) {
			resultVocabs = append(resultVocabs, vocab)
		}
	}

	return resultVocabs
}
