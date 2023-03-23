package handler

import (
	"github.com/POMBNK/trelloSaver/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Saver interface {
	Download(c *gin.Context)
}

type Handler struct {
	saver *service.Saver
}

func New(saver *service.Saver) *Handler {
	return &Handler{saver: saver}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	router.GET("/download", h.Download)

	return router
}

func (h *Handler) Download(c *gin.Context) {

	lists, err := h.saver.GetAllLists(h.saver.Client.BoardID)
	if err != nil {
		log.Printf("Can't get lists %s", err)
	}

	cards, err := h.saver.GetAllCards(lists)
	if err != nil {
		log.Printf("Can't get cards %s", err)
	}

	attachments, err := h.saver.GetAllAttachmentsURLS(cards)
	if err != nil {
		log.Printf("Can't get attachments %s", err)
	}

	if err = h.saver.DownloadFiles(attachments); err != nil {
		log.Printf("Can't get attachments %s", err)
	}
}
