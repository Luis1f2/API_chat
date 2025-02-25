package domain

import "chat/src/users/domain/entities"

type IUser interface {
	Save(user *entities.User) error
	ViewOne(id int) (*entities.User, error)
	ViewAll() ([]entities.User, error)
	Delete(id int) error
	Exists(username string) (bool, error)
}
