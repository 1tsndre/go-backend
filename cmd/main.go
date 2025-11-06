package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/1tsndre/go-backend/internal/app"
)

func main() {
	addr := ":8080"
	if v := os.Getenv("SERVER_PORT"); v != "" {
		addr = ":" + v
	}

	a := app.New()
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Handler(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
	log.Println("server stopped")
}
