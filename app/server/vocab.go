package server

import (
	"encoding/json"
	"net/http"
)

func vocabHandler(w http.ResponseWriter, r *http.Request) {
	id := getIDFromAPI(r)

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	v := allVocabs[id]
	// remove title, so student can't retrive the vocab simple from
	// calling the api directly
	v.Title = ""

	if err := json.NewEncoder(w).Encode(v); err != nil {
		errPrint(w, err)
	}
}
