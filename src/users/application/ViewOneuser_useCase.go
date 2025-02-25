package application

import (
	"chat/src/users/domain"
	"chat/src/users/domain/entities"
	"errors"
)

type ViewUser struct {
	repo domain.IUser
}

func NewViewUser(repo domain.IUser) *ViewUser {
	return &ViewUser{repo: repo}
}

func (vu *ViewUser) Execute(id int) (*entities.User, error) {
	// Validación del ID antes de consultar
	if id <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	user, err := vu.repo.ViewOne(id)
	if err != nil {
		return nil, errors.New("error al obtener el usuario: " + err.Error())
	}

	return user, nil
}
