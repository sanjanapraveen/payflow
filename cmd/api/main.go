package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"payflow/internal/config"
	"payflow/internal/payment/repository"
	"payflow/internal/payment/service"
	httptransport "payflow/internal/transport/http"
)

func main() {
	cfg := config.Load()

	repo, err := repository.NewPostgresPaymentRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}
	defer func() {
		if cerr := repo.Close(); cerr != nil {
			log.Printf("repository close error: %v", cerr)
		}
	}()

	svc := service.NewService(repo)
	handler := httptransport.NewHandler(svc)

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("starting HTTP server on %s", cfg.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v", err)
		}
	}()

	<-stop
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
}
