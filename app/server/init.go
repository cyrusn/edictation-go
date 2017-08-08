package server

import (
	"log"
	"os"
	"path"

	"github.com/cyrusn/edictation-go/app/cli"
	"github.com/cyrusn/edictation-go/app/vocab"
)

// assessment will store the vocab and name of an assessment
type assessment struct {
	name   string
	vocabs []vocab.Vocab
}

// init variables for initiation
var (
	vocabPath       string
	assessmentNames []string
	assessments     []assessment
	voiceSpeed      float64
)

func init() {
	// import cli package
	cli.Parse()

	vocabPath = cli.VocabJSONPath

	vocabDir, err := os.Open(vocabPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filenames, err := vocabDir.Readdirnames(0)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, filename := range filenames {
		vocabJSONPath := path.Join(vocabPath, filename)
		ext := path.Ext(filename)
		assessmentName := filename[:len(filename)-len(ext)]
		assessmentNames = append(assessmentNames, assessmentName)

		a := assessment{
			assessmentName,
			vocab.GetVocabs(vocabJSONPath),
		}

		assessments = append(assessments, a)
	}
}
