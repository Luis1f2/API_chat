package application

import (
	"chat/src/Message/domain"
	"chat/src/Message/domain/entities"
	"errors"
)

type ViewMessage struct {
	repo domain.IMessage
}

func NewViewMessage(repo domain.IMessage) *ViewMessage {
	return &ViewMessage{repo: repo}
}

func (vm *ViewMessage) Execute(id int) (*entities.Message, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}

	return vm.repo.ViewOne(id)
}
