package main

import (
	"context"
	"github.com/POMBNK/trelloSaver/internal/handler"
	"github.com/POMBNK/trelloSaver/internal/server"
	"github.com/POMBNK/trelloSaver/internal/service"
	"github.com/POMBNK/trelloSaver/pkg/trello"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env vars %s", err.Error())
	}

	trelloClient := trello.New(os.Getenv("TOKEN"), os.Getenv("KEY"))

	saveService := service.NewSaver(trelloClient)
	handlers := handler.New(saveService)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			log.Fatalf("error while running server %s", err)
		}
	}()

	log.Println("Service started")

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Service shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}
}
