package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	id := getIDFromAPI(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errPrint(w, err)
		return
	}
	defer r.Body.Close()

	type ans struct {
		Answer string `json:"answer"`
	}
	var v ans

	if err := json.Unmarshal(body, &v); err != nil {
		errPrint(w, err)
		return
	}

	vocab := allVocabs[id]
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	if vocab.Title == v.Answer {
		w.Write([]byte("true"))
		return
	}

	w.Write([]byte("false"))

}
