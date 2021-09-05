package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Altruist-Motion/keeptorch/transport"
)

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler: transport.Router(),
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
