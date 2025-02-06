package application

import (
	"chat/src/users/domain"
)

type SaveUser struct {
	db domain.IUser
}

func NewSaveUser(db domain.IUser) *SaveUser {
	return &SaveUser{db: db}
}

func (su *SaveUser) Execute(username, password string) error {
	return su.db.Save(username, password)
}
