package infrastructure

import (
	"chat/src/Users/domain"
	"chat/src/Users/domain/entities"
	"chat/src/core"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	conn *core.ConnMySQL
}

var _ domain.IUser = &UserRepository{}

func NewUserRepository(db *core.ConnMySQL) domain.IUser {
	return &UserRepository{conn: db}
}

func (repo *UserRepository) Exists(username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE username = ?"

	var count int
	err := repo.conn.FetchRow(query, username).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error al verificar si el usuario existe: %w", err)
	}

	return count > 0, nil
}

func (repo *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := repo.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el usuario: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}
	return nil
}

func (repo *UserRepository) ViewAll() ([]entities.User, error) {
	query := "SELECT id, username, password FROM users"
	rows, err := repo.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, fmt.Errorf("error al escanear el usuario: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre los usuarios: %w", err)
	}

	return users, nil
}

func (repo *UserRepository) ViewOne(id int) (*entities.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = ?"
	row := repo.conn.FetchRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, fmt.Errorf("error al obtener el usuario: %w", err)
	}

	return &user, nil
}

func (repo *UserRepository) Save(user *entities.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := repo.conn.ExecutePreparedQuery(query, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}
	user.ID = int(id)

	return nil
}
