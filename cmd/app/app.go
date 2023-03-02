package main

import (
	"fmt"
	"github.com/POMBNK/trelloSaver/pkg/trello"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	boardID = "63f8c1ab0b7de813525024e1"
	listID  = "63f8c1bdcd591c805bf8abad"
	cardID  = "63f8c1c678b910f04963abc8"
)

func main() {
	//Get .env vars
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env vars %s", err.Error())
	}

	// dirty testing :)
	client := trello.New(os.Getenv("TOKEN"), os.Getenv("KEY"))
	//-------------------------LISTS----------------------------------------
	resp, _ := client.GetList(listID)
	fmt.Printf("%+v\n", resp)
	resp2, _ := client.GetLists(boardID)
	fmt.Printf("%+v\n", resp2)
	for _, v := range resp2 {
		fmt.Println(v.ID, v.Name)
	}
	//-------------------------CARDS----------------------------------------
	respCards, _ := client.GetCardsInList(listID)
	for _, v := range respCards {
		fmt.Printf("id:%s name:%s\n  updated at:%+v\n  listID:%+v\n  desc:%+v\n", v.ID, v.Name, v.DateLastActivity, v.IDList, v.Desc)
	}
	//-------------------------ATTACHMENTS----------------------------------------
	respAttach, _ := client.GetAttachmentsFromCard(cardID)
	for _, v := range respAttach {
		fmt.Printf("Att_id:%s\n  FileName:%s\n  Type:%+v\n  IsUpload:%+v\n  Date:%+v\n", v.ID, v.FileName, v.MimeType, v.IsUpload, v.Date)
	}
}

// requirement request: {{BasedURL}}cards/{{CardID1}}/attachments/{{attachmentID}}?key={{Key}}&token={{Token}}
// 1. BaseURL OK
// 2. Cards IDs <- ListsID <-BoardID
// 3. Attachments IDs
// 4. Key OK
// 5. Token OK
// or https://trello.com/1/cards/{{cardID}}/attachments/{{attachmentID}}/download/{filename.??}
// with Auth Headers
// ???

// GET CARD IDS IN CURRENT LIST
// {{BasedURL}}lists/{{ListID}}/cards?fields=id&key={{Key}}&token={{Token}}
