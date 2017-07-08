package server

import (
	"encoding/json"
	"net/http"

	"github.com/cyrusn/edictation-go/app/vocab"
)

func levelHandler(w http.ResponseWriter, r *http.Request) {
	level := getLevelFromAPI(r)

	var result []int
	vocabs := vocab.FilterVocabsByLevel(level, allVocabs)
	for _, v := range vocabs {
		result = append(result, v.ID)
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		errPrint(w, err)
	}
}
