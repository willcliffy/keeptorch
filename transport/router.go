package transport

import (
	"bytes"
	"log"
	"net/http"

	"github.com/Altruist-Motion/keeptorch/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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
		var payload model.SamplePayload

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
