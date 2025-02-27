package infrastructure

import (
	"chat/src/Message/domain"
	"chat/src/Message/domain/entities"
	"chat/src/core"
	"database/sql"
	"fmt"
)

type MessageRepository struct {
	conn *core.ConnMySQL
}

var _ domain.IMessage = &MessageRepository{}

func NewMessageRepository(db *core.ConnMySQL) domain.IMessage {
	return &MessageRepository{conn: db}
}

func (repo *MessageRepository) Save(message *entities.Message) error {
	query := "INSERT INTO messages (sender_id, receiver_id, content, status) VALUES (?, ?, ?, ?)"
	_, err := repo.conn.ExecutePreparedQuery(query, message.SenderID, message.ReceiverID, message.Content, message.Status)
	return err
}

func (repo *MessageRepository) ViewOne(id int) (*entities.Message, error) {
	query := "SELECT id, sender_id, receiver_id, content, timestamp, status FROM messages WHERE id = ?"
	row := repo.conn.FetchRow(query, id)

	var message entities.Message
	err := row.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.Timestamp, &message.Status)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("mensaje no encontrado")
	} else if err != nil {
		return nil, fmt.Errorf("error al obtener el mensaje: %w", err)
	}

	return &message, nil
}

func (repo *MessageRepository) ViewAll() ([]entities.Message, error) {
	query := "SELECT id, sender_id, receiver_id, content, timestamp, status FROM messages"
	rows, err := repo.conn.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.Timestamp, &message.Status)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (repo *MessageRepository) Delete(id int) error {
	query := "DELETE FROM messages WHERE id = ?"
	_, err := repo.conn.ExecutePreparedQuery(query, id)
	return err
}

func (repo *MessageRepository) UpdateStatus(id int, status string) error {
	query := "UPDATE messages SET status = ? WHERE id = ?"
	_, err := repo.conn.ExecutePreparedQuery(query, status, id)
	return err
}

func (repo *MessageRepository) ViewByUser(userID int) ([]entities.Message, error) {
	query := "SELECT id, sender_id, receiver_id, content, timestamp, status FROM messages WHERE receiver_id = ?"
	rows, err := repo.conn.FetchRows(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.Timestamp, &message.Status)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
