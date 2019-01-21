package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", HandleWriteBlock).Methods("POST")
	return muxRouter
}
