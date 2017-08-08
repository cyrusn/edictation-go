package server

import (
	"net/http"
	"strconv"
)

func assessmentSizeHandler(w http.ResponseWriter, r *http.Request) {
	vocabs, err := getVocabsFromAPI(r)
	if err != nil {
		errPrint(w, err)
		return
	}
	size := len(vocabs)
	w.Write([]byte(strconv.Itoa(size)))
}
