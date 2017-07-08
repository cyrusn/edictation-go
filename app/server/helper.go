package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
