package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cyrusn/edictation-go/app/vocab"
)

func reportHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errPrint(w, err)
		return
	}

	defer r.Body.Close()
	var indexes []int

	if err = json.Unmarshal(body, &indexes); err != nil {
		errPrint(w, err)
		return
	}

	vocabs, err := getVocabsFromAPI(r)
	if err != nil {
		errPrint(w, err)
		return
	}

	filteredVocabs := vocab.FilterVocabsByIndexArray(indexes, vocabs)
	if err := json.NewEncoder(w).Encode(filteredVocabs); err != nil {
		errPrint(w, err)
	}
}
