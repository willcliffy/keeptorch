package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/jsonapi"
)

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler: Router(),
	}

	shutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("API Shutdown Error: %v\n", err)
		}
		close(shutdown)
	}()

	log.Printf("HTTP api listening on 8080\n")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("err: %v\n", err)
	}

	<-shutdown
}

func Router() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET request from client\n")

		w.Header().Set("Content-Type", jsonapi.MediaType)
		w.Header().Set("Cache-Control", "private, max-age=60")
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode("Hello, world!"); err != nil {
			log.Println("error marshalling payload: ", err)
		}
	})

	return router
}
