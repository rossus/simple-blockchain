package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {

	response, err := json.MarshalIndent(payload, "", "  ")

	if err != nil {
		log.Println(":':'---->")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
