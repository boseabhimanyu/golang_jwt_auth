package main

import (
	httpserver "golang-jwt-auth/internal/server"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httpserver.NewRouter()

	//standard go type that runs a server
	srv := &http.Server{
		Addr:        ":5000",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("API running on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("server closed")
			return
		}

		log.Fatalf("server error: %v", err)
	}

}
