package handler

import (
	"chatting/pkg/domain"
	"chatting/pkg/usecases/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	MessageInteractor message.MessageInteractor
}

func (m *MessageController) SendMessage(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.MessageInteractor.SendMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (m *MessageController) GetMessages(c *gin.Context) {
	senderID := c.Param("senderID")
	receiverID := c.Param("receiverID")
	messages, err := m.MessageInteractor.GetMessages(senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
