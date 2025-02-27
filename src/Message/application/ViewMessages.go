package application

import (
	"chat/src/Message/domain"
	"chat/src/Message/domain/entities"
	"errors"
)

type ViewMessages struct {
	repo domain.IMessage
}

func NewViewMessages(repo domain.IMessage) *ViewMessages {
	return &ViewMessages{repo: repo}
}

// ✅ Obtener mensajes de un usuario
func (vm *ViewMessages) ExecuteForUser(userID int) ([]entities.Message, error) {
	if userID <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	messages, err := vm.repo.ViewByUser(userID)
	if err != nil {
		return nil, errors.New("error al obtener los mensajes del usuario: " + err.Error())
	}

	return messages, nil
}

func (vm *ViewMessages) ExecuteAll() ([]entities.Message, error) {
	messages, err := vm.repo.ViewAll()
	if err != nil {
		return nil, errors.New("error al obtener la lista de mensajes: " + err.Error())
	}

	return messages, nil
}
