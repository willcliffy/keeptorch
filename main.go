package main

import (
	"bytes"
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

	router.Get("/get-hello-world", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET request from client\n")
		SendResponseToClient(w, http.StatusOK, "Hello, world!")
	})

	router.Post("/post-hello-world", func(w http.ResponseWriter, r *http.Request) {	
		var payload SamplePayload

		bodyBytes := new(bytes.Buffer)
		_, err := bodyBytes.ReadFrom(r.Body)
	
		if err != nil {
			log.Printf("Error reading request: %v\n", err)
			SendResponseToClient(w, http.StatusInternalServerError, err)
			return
		}
	
		if err = payload.UnmarshalJSON(bodyBytes.Bytes()); err != nil {
			log.Printf("Error unmarshalling payload: %v\n", err)
			SendResponseToClient(w, http.StatusInternalServerError, err)
			return
		}
	
		log.Printf("POST request from client. Payload: %v\n", payload)
	
		SendResponseToClient(w, http.StatusNoContent, nil)
	})

	return router
}

func SendResponseToClient(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.Header().Set("Cache-Control", "private, max-age=60")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("error marshalling payload: ", err)
	}
}
