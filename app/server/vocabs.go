package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cyrusn/edictation-go/app/vocab"
)

func vocabsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errPrint(w, err)
		return
	}
	defer r.Body.Close()
	var v []int

	if err := json.Unmarshal(body, &v); err != nil {
		errPrint(w, err)
		return
	}

	vocabs := vocab.FilterVocabsByIDs(v, allVocabs)
	if err := json.NewEncoder(w).Encode(vocabs); err != nil {
		errPrint(w, err)
	}
}
