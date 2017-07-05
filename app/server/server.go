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

// init variables for initiation
var (
	allVocabs       []vocab.Vocab
	voiceServerPort string
)

func init() {
	// import cli package
	cli.Parse()

	vocabJSONPath := cli.VocabJSONPath
	allVocabs = vocab.GetVocabs(vocabJSONPath)

}

// Serve ...
func Serve() {
	port := cli.Port
	voiceServerPort = cli.VoiceServerPort
	servingFolder := cli.ServingFolder

	r := mux.NewRouter()

	type api struct {
		method  string
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}

	var apis = []api{
		// Get Vocab by vocabID
		api{"GET", "/api/id/{id:[0-9]+}", vocabIDHandler},

		// get the list of vocabs by given level
		api{"GET", "/api/level/{level:[1-6]}", levelHandler},

		// get the src of tts by given vocabID
		api{"GET", "/voice/id/{id:[0-9]+}", voiceHandler},

		// Post a list of VocabID, and will return the list of Vocabs
		// for reporting
		api{"POST", "/api/vocabs", vocabsHandler},

		// post a vocabID and with answer, return bool
		api{"POST", "/check/id/{id:[0-9]+}", checkHandler},
	}

	for _, a := range apis {
		r.HandleFunc(a.path, a.handler).Methods(a.method)
	}

	// server static folder
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

// httpLogger is a simple logger to log the request info
func httpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

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

func voiceHandler(w http.ResponseWriter, r *http.Request) {
	id := getIDFromAPI(r)

	v := allVocabs[id]
	fileResponse, err := voice.GetVoiceResponse(v.Title, voiceServerPort)
	if err != nil {
		errPrint(w, err)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	fileResponse.Write(w)
}

func vocabIDHandler(w http.ResponseWriter, r *http.Request) {
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

func errPrint(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"error": %s}`, err)
}

func getLevelFromAPI(r *http.Request) int {
	slevel := mux.Vars(r)["level"]
	level, _ := strconv.Atoi(slevel)
	return level
}

func getIDFromAPI(r *http.Request) int {
	sID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(sID)
	return id
}
