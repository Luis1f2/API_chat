package application

import "chat/src/users/domain"

type DeleteUser struct {
	db domain.IUser
}

func NewDeleteUser(db domain.IUser) *DeleteUser {
	return &DeleteUser{db: db}
}

func (du *DeleteUser) Execute(id int) error {
	return du.db.Delete(id)
}
