package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cyrusn/edictation-go/app/cli"
	"github.com/gorilla/mux"
)

// Serve ...
func Serve() {
	port := cli.Port
	voiceSpeed = cli.VoiceSpeed
	servingFolder := cli.ServingFolder

	r := mux.NewRouter()

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
		Addr:    "localhost" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting up, Serving " + http.Dir(servingFolder))
	fmt.Println("Available on http://localhost" + port)
	log.Fatal(srv.ListenAndServe())
}
