package main

import (
	"github.com/POMBNK/trelloSaver/internal/service"
	"github.com/POMBNK/trelloSaver/pkg/trello"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env vars %s", err.Error())
	}

	trelloClient := trello.New(os.Getenv("TOKEN"), os.Getenv("KEY"))
	saveService := service.NewSaver(trelloClient)

	lists, err := saveService.GetAllLists(os.Getenv("BOARD_ID"))
	if err != nil {
		log.Fatalf("Can't get lists %s", err)
	}
	cards, err := saveService.GetAllCards(lists)
	if err != nil {
		log.Fatalf("Can't get cards %s", err)
	}

	attachments, err := saveService.GetAllAttachmentsURLS(cards)
	if err != nil {
		log.Fatalf("Can't get attachments %s", err)
	}

	if err = saveService.DownloadFiles(attachments); err != nil {
		log.Fatalf("Can't get attachments %s", err)
	}

}
