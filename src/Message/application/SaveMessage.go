package application

import (
	"chat/src/Message/domain"
	"chat/src/Message/domain/entities"
	"errors"
	"time"
)

type SaveMessage struct {
	repo domain.IMessage
}

func NewSaveMessage(repo domain.IMessage) *SaveMessage {
	return &SaveMessage{repo: repo}
}

func (sm *SaveMessage) Execute(senderID, receiverID int, content string) error {
	if senderID <= 0 || receiverID <= 0 || content == "" {
		return errors.New("todos los campos son obligatorios")
	}

	message := &entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Timestamp:  time.Now(),
		Status:     "sent",
	}

	return sm.repo.Save(message)
}
