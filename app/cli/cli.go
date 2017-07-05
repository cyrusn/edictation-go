package cli

import (
	"flag"
	"fmt"
)

var (
	// Port store the port string for server
	Port string
	// VoiceServerPort store the port string for voice server
	VoiceServerPort string
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
		defaultVoicePort     = ":5001"
		defaultServingFolder = "./html/dist"
		defaultVocabJSONPath = "./vocab/vocab.json"

		usagePort          = "server port"
		usageVoicePort     = "voice server's port"
		usageServingFolder = "location of static files"
		usageVocabJSONPath = "location of vocab.json"
	)

	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+" shorthand")

	flag.StringVar(&VoiceServerPort, "vport", defaultVoicePort, usageVoicePort)
	flag.StringVar(&VoiceServerPort, "vp", defaultVoicePort, usageVoicePort+" shorthand")

	flag.StringVar(&ServingFolder, "static", defaultServingFolder, usageServingFolder)
	flag.StringVar(&ServingFolder, "s", defaultServingFolder, usageServingFolder+" shorthand")

	flag.StringVar(&VocabJSONPath, "vocab", defaultVocabJSONPath, usageVocabJSONPath)
	flag.StringVar(&VocabJSONPath, "v", defaultVocabJSONPath, usageVocabJSONPath+" shorthand")
}

// Parse parses the flag for server
func Parse() {
	flag.Parse()
}
