package service

import (
	"github.com/POMBNK/trelloSaver/pkg/trello"
	"log"
)

//Todo: make folders with list's names and sort files by folders

type Saver struct {
	Client *trello.Client
}

func NewSaver(client *trello.Client) *Saver {
	return &Saver{Client: client}
}

func (s *Saver) GetAllAttachmentsURLS(cards trello.Cards) (map[string]string, error) {
	files := make(map[string]string)
	var attachmentsList trello.Attachments
	for _, card := range cards {
		attachments, err := s.Client.GetAttachmentsFromCard(card.ID)
		if err != nil {
			return nil, err
		}
		attachmentsList = append(attachmentsList, attachments...)
	}

	for _, attachment := range attachmentsList {
		files[attachment.FileName] = attachment.URL
	}

	return files, nil
}

func (s *Saver) DownloadFiles(files map[string]string) error {
	for name, file := range files {

		err := s.Client.DownloadFile(name, file)
		log.Printf("Downloading... %s", file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Saver) GetAllLists(boardID string) (trello.Lists, error) {
	lists, err := s.Client.GetLists(boardID)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (s *Saver) GetAllCards(lists trello.Lists) (trello.Cards, error) {
	var allCards trello.Cards

	for _, list := range lists {
		cards, err := s.Client.GetCardsInList(list.ID)
		if err != nil {
			return nil, err
		}
		allCards = append(allCards, cards...)
	}

	return allCards, nil
}

//func createFolder(folderName string) error {
//	os.Mkdir("downloaded", 0700)
//	out, err := os.Create("downloaded/" + folderName)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	return nil
//}
