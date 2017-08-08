package server

import (
	"encoding/json"
	"net/http"
)

func assessmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(assessmentNames); err != nil {
		errPrint(w, err)
	}
}
