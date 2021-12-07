package main

import (
	"context"
	"github.com/gichohi/go-cqrs-rabbitmq/api"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/db"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	db.InitDB()

	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	httpHandler := api.NewHandler()
	s := &http.Server{
		Handler: httpHandler,
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	} else {
		log.Printf("Listening on port: %s", addr)
	}

	<-ctx.Done()

	defer Stop(s)
}


func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}