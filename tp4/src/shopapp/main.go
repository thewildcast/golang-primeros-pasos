package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"server"
)

const addr = ":8080"

// Code structure
// https://getgophish.com/blog/post/2018-12-02-building-web-servers-in-go/
func main() {
	handler := server.NewRouteHandler()
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start the server
	go func() {
		log.Fatal(httpServer.ListenAndServe())
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
