package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/duckonomy/my-mlh-go-example/internal/handlers"
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":5000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
