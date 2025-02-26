package application

import (
	"chat/src/Users/domain"
	"chat/src/Users/domain/entities"
	"errors"
)

type SaveUser struct {
	repo domain.IUser
}

func NewSaveUser(repo domain.IUser) *SaveUser {
	return &SaveUser{repo: repo}
}

func (su *SaveUser) Execute(username, password string) error {
	// Verificar si el usuario ya existe
	exists, err := su.repo.Exists(username)
	if err != nil {
		return errors.New("error verificando existencia del usuario: " + err.Error())
	}
	if exists {
		return errors.New("el usuario ya existe")
	}

	// Crear una nueva instancia del usuario con la contraseña encriptada
	user, err := entities.NewUser(username, password)
	if err != nil {
		return errors.New("error al crear el usuario: " + err.Error())
	}

	// Guardar en el repositorio
	return su.repo.Save(user)
}
