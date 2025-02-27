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

// ✅ Obtener un solo mensaje por ID
func (vm *ViewMessage) Execute(id int) (*entities.Message, error) {
	if id <= 0 {
		return nil, errors.New("ID de mensaje inválido")
	}

	message, err := vm.repo.ViewOne(id)
	if err != nil {
		return nil, errors.New("error al obtener el mensaje: " + err.Error())
	}

	return message, nil
}
