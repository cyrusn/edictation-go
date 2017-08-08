package server

import "net/http"

type api struct {
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

var apis = []api{
	// get the list of assessment name
	api{"GET", "/api/assessment", assessmentHandler},

	// Get Vocab by vocabID
	api{"GET", "/api/assessment/{name}/index/{index:[0-9]+}", vocabHandler},

	// get the no of vocab in given assessment name
	api{"GET", "/api/assessment/{name}/size", assessmentSizeHandler},

	// get the src of tts by given assessment and index
	api{"GET", "/api/voice/assessment/{name}/index/{index:[0-9]+}", voiceHandler},

	// post a vocabID and with answer in body, return bool
	api{"GET", "/api/check/assessment/{name}/index/{index:[0-9]+}", checkHandler},

	// Post a list of VocabID in body, and will return the list of Vocabs
	// for reporting
	api{"POST", "/api/assessment/{name}/report", reportHandler},
}
