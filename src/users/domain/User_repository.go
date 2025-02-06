package domain

import "chat/src/Users/domain/entities"

type IUser interface {
	Save(username, password string) error
	ViewOne(id int) (*entities.User, error)
	ViewAll() ([]entities.User, error)
	Delete(id int) error
}
