package http

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Run() error {
	mux := MakeMuxRouter()

	httpAddr := os.Getenv("ADDR1")
	log.Println("Listening on ", os.Getenv("ADDR1"))

	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
