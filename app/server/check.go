package server

import (
	"net/http"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	ans := r.FormValue("ans")
	vocab, err := getVocabFromAPI(r)
	if err != nil {
		errPrint(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	if vocab.Title == ans {
		w.Write([]byte("true"))
		return
	}

	w.Write([]byte("false"))

}
