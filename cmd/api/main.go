package main

import (
	"context"
	"golang-jwt-auth/internal/app"
	httpserver "golang-jwt-auth/internal/server"
	"log"
	"net/http"
	"time"
)

func main() {
	//root context
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Statup failed: %v", &err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Shutdown warning:%v", err)
		}
	}()
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
