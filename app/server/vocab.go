package server

import (
	"encoding/json"
	"net/http"
)

func vocabHandler(w http.ResponseWriter, r *http.Request) {
	v, err := getVocabFromAPI(r)
	if err != nil {
		errPrint(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	// remove title, so student can't retrive the vocab directly from api
	v.Title = ""

	if err := json.NewEncoder(w).Encode(v); err != nil {
		errPrint(w, err)
	}
}
