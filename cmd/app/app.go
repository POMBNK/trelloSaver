package main

import (
	"fmt"
	"github.com/POMBNK/trelloSaver/pkg/trello"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//Get .env vars
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env vars %s", err.Error())
	}

	// testing :)
	client := trello.New(os.Getenv("TOKEN"), os.Getenv("KEY"))
	resp, _ := client.GetList("63f8c1bdcd591c805bf8abad")
	fmt.Printf("%+v\n", resp)
}
