package application

import (
	"chat/src/Users/domain"
	"chat/src/Users/domain/entities"
	"errors"
)

type LoginUser struct {
	repo domain.IUser
}

func NewLoginUser(repo domain.IUser) *LoginUser {
	return &LoginUser{repo: repo}
}

func (lu *LoginUser) Execute(username, password string) (*entities.User, error) {
	user, err := lu.repo.Authenticate(username, password)
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}
	return user, nil
}
