package cli

import (
	"flag"
	"fmt"
)

var (
	// Port store the port string for server
	Port string
	// VoiceSpeed control the speed of the voice
	VoiceSpeed float64
	// ServingFolder is location that serve the public files (index.html ...)
	ServingFolder string
	// VocabJSONPath is the location of all vocabularies
	VocabJSONPath string
)

func init() {
	flag.Usage = func() {
		const welcomeText = "eDictation server for LPSS.\nUsage:"
		fmt.Println(welcomeText)
		flag.PrintDefaults()
	}

	const (
		defaultPort          = ":5050"
		defaultVoiceSpeed    = 0.5
		defaultServingFolder = "./html/dist"
		defaultVocabJSONPath = "./vocab/vocab.json"

		usagePort          = "server port"
		usageVoiceSpeed    = "speed of voice"
		usageServingFolder = "location of static files"
		usageVocabJSONPath = "location of vocab.json"
	)

	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+" shorthand")

	flag.Float64Var(&VoiceSpeed, "speed", defaultVoiceSpeed, usageVoiceSpeed)
	flag.Float64Var(&VoiceSpeed, "vs", defaultVoiceSpeed, usageVoiceSpeed+" shorthand")

	flag.StringVar(&ServingFolder, "static", defaultServingFolder, usageServingFolder)
	flag.StringVar(&ServingFolder, "s", defaultServingFolder, usageServingFolder+" shorthand")

	flag.StringVar(&VocabJSONPath, "vocab", defaultVocabJSONPath, usageVocabJSONPath)
	flag.StringVar(&VocabJSONPath, "v", defaultVocabJSONPath, usageVocabJSONPath+" shorthand")
}

// Parse parses the flag for server
func Parse() {
	flag.Parse()
}
