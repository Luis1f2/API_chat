package domain

import "chat/src/Users/domain/entities"

type IUser interface {
	Save(user *entities.User) error
	ViewOne(id int) (*entities.User, error)
	ViewAll() ([]entities.User, error)
	Delete(id int) error
	Exists(username string) (bool, error)
	Authenticate(username, password string) (*entities.User, error)
}
