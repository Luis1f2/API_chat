package application

import (
	"chat/src/users/domain"
	"errors"
)

type DeleteUser struct {
	repo domain.IUser
}

func NewDeleteUser(repo domain.IUser) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (du *DeleteUser) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID de usuario inválido")
	}

	err := du.repo.Delete(id)
	if err != nil {
		return errors.New("error al eliminar el usuario: " + err.Error())
	}
	return nil
}
