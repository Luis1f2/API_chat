package application

import (
	"chat/src/users/domain"
	"chat/src/users/domain/entities"
)

type ViewUser struct {
	db domain.IUser
}

func NewViewUser(db domain.IUser) *ViewUser {
	return &ViewUser{db: db}
}

func (vu *ViewUser) Execute(id int) (*entities.User, error) {
	return vu.db.ViewOne(id)
}
