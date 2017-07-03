package cli

import (
	"flag"
	"fmt"
)

var (
	// Port is ..
	Port string
	// ServingFolder is ...
	ServingFolder string
	// VocabJSONPath is ...
	VocabJSONPath string
)

func init() {
	flag.Usage = func() {
		const welcomeText = "Fresh card server for LPSS.\nUsage:"
		fmt.Println(welcomeText)
		flag.PrintDefaults()
	}

	const (
		defaultPort          = ":8080"
		defaultServingFolder = "./html/dist"
		defaultVocabJSONPath = "./vocab/vocab.json"

		usagePort          = "server port"
		usageServingFolder = "location of static files"
		usageVocabJSONPath = "location of vocab.json"
	)

	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+" shorthand")

	flag.StringVar(&ServingFolder, "static", defaultServingFolder, usageServingFolder)
	flag.StringVar(&ServingFolder, "s", defaultServingFolder, usageServingFolder+" shorthand")

	flag.StringVar(&VocabJSONPath, "vocab", defaultVocabJSONPath, usageVocabJSONPath)
	flag.StringVar(&VocabJSONPath, "v", defaultVocabJSONPath, usageVocabJSONPath+" shorthand")
}

// Run ...
func Parse() {
	flag.Parse()
}
