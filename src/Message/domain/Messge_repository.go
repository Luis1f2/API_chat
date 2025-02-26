package domain

import "chat/src/Message/domain/entities"

type IMessage interface {
	Save(message *entities.Message) error
	ViewOne(id int) (*entities.Message, error)
	ViewAll() ([]entities.Message, error)
	Delete(id int) error
	UpdateStatus(id int, status string) error
}
