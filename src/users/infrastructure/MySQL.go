package infrastructure

import (
	"chat/src/Users/domain/entities"
	"chat/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

// NewMySQL inicializa una nueva conexión MySQL.
func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Save agrega un nuevo usuario a la base de datos.
func (mysql *MySQL) Save(username, password string) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, username, password)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

// Delete elimina un usuario de la base de datos por su ID.
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

// ViewAll obtiene todos los usuarios de la base de datos.
func (mysql *MySQL) ViewAll() ([]entities.User, error) {
	query := "SELECT id, username, password FROM users"

	// Corregido: Ahora capturamos el error de FetchRows
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		users = append(users, user)
	}

	// Validar errores después de iterar sobre las filas
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return users, nil
}

// ViewOne obtiene un usuario específico por su ID.
func (mysql *MySQL) ViewOne(id int) (*entities.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = ?"

	// Corregido: Capturar error de FetchRows
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer rows.Close()

	var user entities.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningún usuario con el ID %d", id)
	}

	// Validar errores después de iterar sobre las filas
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return &user, nil
}
