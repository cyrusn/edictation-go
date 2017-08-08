package server

import (
	"net/http"

	"github.com/cyrusn/edictation-go/app/voice"
)

func voiceHandler(w http.ResponseWriter, r *http.Request) {
	v, err := getVocabFromAPI(r)
	if err != nil {
		errPrint(w, err)
		return
	}

	url, err := voice.GetVoiceSource(v.Title, voiceSpeed)
	if err != nil {
		errPrint(w, err)
		return
	}

	fileResponse, err := http.Get(url)
	if err != nil {
		errPrint(w, err)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	fileResponse.Write(w)
}
