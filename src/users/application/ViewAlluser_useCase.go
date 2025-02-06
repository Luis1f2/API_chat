package application

import (
	"chat/src/users/domain/entities"
	"chat/src/users/domain"
)

type ViewUsers struct {
	db domain.IUser
}

func NewViewUsers(db domain.IUser) *ViewUsers {
	return &ViewUsers{db: db}
}

func (vu *ViewUsers) Execute() ([]entities.User, error) {
	return vu.db.ViewAll()
}
