package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cyrusn/edictation-go/app/vocab"
	"github.com/gorilla/mux"
)

// httpLogger is a simple logger to log the request info
func httpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func errPrint(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"error": %s}`, err)
}

func getAssessmentNameFromAPI(r *http.Request) string {
	assessmentName := mux.Vars(r)["name"]
	return assessmentName
}

func getIndexFromAPI(r *http.Request) int {
	indexString := mux.Vars(r)["index"]
	index, _ := strconv.Atoi(indexString)
	return index
}

func getAssessmentByName(name string) (assessment, error) {
	for _, a := range assessments {
		if a.name == name {
			return a, nil
		}
	}
	return assessment{}, errors.New("No such assessment")
}

func getVocabsFromAPI(r *http.Request) ([]vocab.Vocab, error) {
	name := getAssessmentNameFromAPI(r)
	a, err := getAssessmentByName(name)
	if err != nil {
		return []vocab.Vocab{}, err
	}

	return a.vocabs, nil
}

func getVocabFromAPI(r *http.Request) (vocab.Vocab, error) {
	index := getIndexFromAPI(r)

	vocabs, err := getVocabsFromAPI(r)
	if err != nil {
		return vocab.Vocab{}, err
	}
	return vocabs[index], err
}
