package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	port          string
	servingFolder string
)

// JSONResponse is ...
type JSONResponse struct {
	Title      string
	ID         int    `json:"id"`
	Definition string `json:"definition"`
	// Level      int    `json:"level"`
}

var jsonResponses = []JSONResponse{
	JSONResponse{"hello", 1, "你好"},
	JSONResponse{"apple", 1, "蘋果"},
}

func init() {
	flag.Usage = func() {
		const welcomeText = "Fresh card server for LPSS.\nUsage:"
		fmt.Println(welcomeText)
		flag.PrintDefaults()
	}

	const (
		defaultPort          = ":8080"
		defaultServingFolder = "./html"

		usagePort          = "server port"
		usageServingFolder = "location of static files"
	)

	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, usagePort+" shorthand")

	flag.StringVar(&servingFolder, "static", defaultServingFolder, usageServingFolder)
	flag.StringVar(&servingFolder, "s", defaultServingFolder, usageServingFolder+" shorthand")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/api/{level:[1-6]}/{id:[0-9]+}", vocabHandler)

	r.HandleFunc("/api/{level:[1-6]}", levelHandler)

	r.HandleFunc("/check/{level:[1-6]}/{id:[0-9]+}", checkHandler)

	r.HandleFunc("/mp3/{level:[1-6]}/{id:[0-9]+}", mp3Handler)

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
	w.Write([]byte(fmt.Sprintf(`{"noOfQuestions":%d}`, len(jsonResponses))))
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)

	answer := r.FormValue("answer")
	fmt.Println(answer)
	if answer == jsonResponses[id].Title {
		w.Write([]byte(`{"result":true}`))
	} else {
		w.Write([]byte(`{"result":false}`))
	}
}

func mp3Handler(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	level := mux.Vars(r)["level"]
	title := jsonResponses[id].Title

	filename := fmt.Sprintf("./mp3/level%v/%s.mp3", level, title)
	fmt.Println(filename)
	http.ServeFile(w, r, filename)
}

func vocabHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		errPrint(w, err)
	}

	// level, err := strconv.Atoi(params["level"])
	if err != nil {
		errPrint(w, err)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(jsonResponses[id]); err != nil {
		errPrint(w, err)
	}
}

func errPrint(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"error": %s}`, err)
}
