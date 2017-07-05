package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyrusn/edictation-go/app/cli"
	"github.com/cyrusn/edictation-go/app/vocab"
	"github.com/cyrusn/edictation-go/app/voice"
	"github.com/gorilla/mux"
)

var allVocabs []vocab.Vocab
var levelVocabs [][]vocab.Vocab

func init() {
	cli.Parse()

	vocabJSONPath := cli.VocabJSONPath
	allVocabs = vocab.GetVocabs(vocabJSONPath)

}

// Serve ...
func Serve() {
	port := cli.Port
	servingFolder := cli.ServingFolder

	r := mux.NewRouter()
	r.HandleFunc("/api/id/{id:[0-9]+}", vocabHandler)

	r.HandleFunc("/api/level/{level:[1-6]}", levelHandler)

	r.HandleFunc("/api/vocabs", vocabsHandler)

	r.HandleFunc("/check/id/{id:[0-9]+}", checkHandler).Methods("POST")

	r.HandleFunc("/voice/id/{id:[0-9]+}", voiceHandler)

	r.PathPrefix("/").Handler(
		http.StripPrefix(
			"/",
			http.FileServer(http.Dir(servingFolder)),
		),
	)

	srv := &http.Server{
		Handler: httpLogger(r),
		Addr:    "127.0.0.1" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting up, Serving " + http.Dir(servingFolder))
	fmt.Println("Available on http://localhost" + port)
	log.Fatal(srv.ListenAndServe())
}

func httpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func vocabsHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var v []int

	if err := json.Unmarshal(body, &v); err != nil {
		errPrint(w, err)
	}

	vocabs := vocab.FilterVocabsByIDs(v, allVocabs)
	if err := json.NewEncoder(w).Encode(vocabs); err != nil {
		errPrint(w, err)
	}

}

func levelHandler(w http.ResponseWriter, r *http.Request) {
	levelIndex := getLevelIndex(w, r)

	var result []int
	vocabs := vocab.FilterVocabsByLevel(levelIndex+1, allVocabs)
	for _, v := range vocabs {
		result = append(result, v.ID)
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		errPrint(w, err)
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {

	idIndex := getIDIndex(w, r)

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	answer := r.FormValue("answer")
	vocab := allVocabs[idIndex]

	if answer == vocab.Title {
		w.Write([]byte(`{"result":true}`))
	} else {
		w.Write([]byte(`{"result":false}`))
	}
}

func voiceHandler(w http.ResponseWriter, r *http.Request) {
	idIndex := getIDIndex(w, r)

	v := allVocabs[idIndex]
	fileResponse, err := voice.GetVoiceResponse(v.Title)
	if err != nil {
		errPrint(w, err)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	fileResponse.Write(w)
}

func vocabHandler(w http.ResponseWriter, r *http.Request) {
	idIndex := getIDIndex(w, r)

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	v := allVocabs[idIndex]
	v.Title = ""

	if err := json.NewEncoder(w).Encode(v); err != nil {
		errPrint(w, err)
	}
}

func errPrint(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"error": %s}`, err)
}

func getLevelIndex(w http.ResponseWriter, r *http.Request) int {
	slevel := mux.Vars(r)["level"]
	level, err := strconv.Atoi(slevel)
	if err != nil {
		errPrint(w, err)
	}
	var index = level - 1
	return index
}

func getIDIndex(w http.ResponseWriter, r *http.Request) int {
	sID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sID)
	if err != nil {
		errPrint(w, err)
	}
	index := id
	return index
}
