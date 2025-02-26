package application

import (
	"chat/src/Users/domain"
	"chat/src/Users/domain/entities"
	"errors"
)

type ViewUsers struct {
	repo domain.IUser
}

func NewViewUsers(repo domain.IUser) *ViewUsers {
	return &ViewUsers{repo: repo}
}

func (vu *ViewUsers) Execute() ([]entities.User, error) {
	users, err := vu.repo.ViewAll()
	if err != nil {
		return nil, errors.New("error al obtener la lista de usuarios: " + err.Error())
	}

	return users, nil
}
