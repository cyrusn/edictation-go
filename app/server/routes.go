package server

import "net/http"

type api struct {
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

var apis = []api{
	// Get Vocab by vocabID
	api{"GET", "/api/vocab/id/{id:[0-9]+}", vocabHandler},

	// Post a list of VocabID in body, and will return the list of Vocabs
	// for reporting
	api{"POST", "/api/vocab/id", vocabsHandler},

	// get the list of vocabs by given level
	api{"GET", "/api/level/{level:[1-6]}", levelHandler},

	// get the src of tts by given vocabID
	api{"GET", "/api/voice/vocab/id/{id:[0-9]+}", voiceHandler},

	// post a vocabID and with answer in body, return bool
	api{"post", "/api/check/vocab/id/{id:[0-9]+}", checkHandler},
}
