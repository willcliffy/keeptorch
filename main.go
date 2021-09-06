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
	// create a server "object"
	// `Handler` is any `http.Handler`, in this case we're using a `chi` router
	server := http.Server{
		Addr: ":8080",
		Handler: Router(),
	}

	// `chan` is short for "channel". Channels are used to communicate between `goroutines`, 
	// which are the built-in golang implementation of threads
	shutdown := make(chan struct{})
	
	// this is a goroutine, which is a thread.
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		// this line blocks the goroutine until we receive an interrupt signal from the 
		// process that started the api
		<-sigint

		// gracefully handle shutdown
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("API Shutdown Error: %v\n", err)
		}
		close(shutdown)
	}()

	// main entry point:
	log.Printf("HTTP api listening on 8080\n")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("err: %v\n", err)
	}

	// block the main thread to prevent prematurely killing your goroutines.
	// This allows us to gracefully close the api before exiting
	<-shutdown
}

func Router() chi.Router {
	router := chi.NewRouter()

	// apply middleware here as needed
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)

	// this is the only request that the API accepts, which simply returns a status 204 and the string "Hello, world!"
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
