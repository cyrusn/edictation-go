package server

import (
	"github.com/cyrusn/edictation-go/app/cli"
	"github.com/cyrusn/edictation-go/app/vocab"
)

func init() {
	// import cli package
	cli.Parse()

	vocabJSONPath := cli.VocabJSONPath
	allVocabs = vocab.GetVocabs(vocabJSONPath)
}
