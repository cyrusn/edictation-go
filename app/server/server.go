package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyrusn/edictation-tsf/app/cli"
	"github.com/cyrusn/edictation-tsf/app/vocab"
	"github.com/cyrusn/edictation-tsf/app/voice"
	"github.com/gorilla/mux"
)

var allVocabs []vocab.Vocab
var levelVocabs [][]vocab.Vocab

func init() {
	cli.Parse()

	vocabJSONPath := cli.VocabJSONPath
	allVocabs = vocab.GetVocabs(vocabJSONPath)

	for i := 0; i < 6; i++ {
		vocabs := vocab.FilterVocabsByLevel(i+1, allVocabs)
		levelVocabs = append(levelVocabs, vocabs)
	}

}

// Serve ...
func Serve() {
	port := cli.Port
	servingFolder := cli.ServingFolder

	r := mux.NewRouter()
	r.HandleFunc("/api/{level:[1-6]}/{id:[0-9]+}", vocabHandler)

	r.HandleFunc("/api/{level:[1-6]}", levelHandler)

	r.HandleFunc("/check/{level:[1-6]}/{id:[0-9]+}", checkHandler).
		Methods("POST")

	r.HandleFunc("/voice/{level:[1-6]}/{id:[0-9]+}", voiceHandler)

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

func levelHandler(w http.ResponseWriter, r *http.Request) {
	levelIndex := getLevelIndex(w, r)

	w.Write([]byte(fmt.Sprintf(`{"noOfQuestions":%d}`, len(levelVocabs[levelIndex]))))
}

func checkHandler(w http.ResponseWriter, r *http.Request) {

	levelIndex := getLevelIndex(w, r)
	idIndex := getIDIndex(w, r)

	answer := r.FormValue("answer")
	vocab := levelVocabs[levelIndex][idIndex]

	if answer == vocab.Title {
		w.Write([]byte(`{"result":true}`))
	} else {
		w.Write([]byte(`{"result":false}`))
	}
}

func voiceHandler(w http.ResponseWriter, r *http.Request) {
	levelIndex := getLevelIndex(w, r)
	idIndex := getIDIndex(w, r)

	v := levelVocabs[levelIndex][idIndex]
	fileResponse, err := voice.GetVoiceResponse(v.Title)
	if err != nil {
		errPrint(w, err)
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	fileResponse.Write(w)
}

func vocabHandler(w http.ResponseWriter, r *http.Request) {
	idIndex := getIDIndex(w, r)
	levelIndex := getLevelIndex(w, r)

	w.Header().Set("Content-Type", "application/json")
	v := levelVocabs[levelIndex][idIndex]
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
	index := id - 1
	return index
}
