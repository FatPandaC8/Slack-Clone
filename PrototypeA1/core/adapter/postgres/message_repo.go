package postgres

import (
	"core/domain/message"
	"database/sql"
	"time"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) Save(msg *message.Message) error {
	_, err := r.db.Exec(
		`INSERT INTO messages (id, roomid, senderid, content, createdat)
		 VALUES ($1, $2, $3, $4, $5)`,
		msg.ID(), msg.RoomID(), msg.SenderID(), msg.Content(), time.Now(),
	)
	return err
}

func (r *MessageRepository) FindByRoom(roomID string) ([]*message.Message, error) {
	rows, err := r.db.Query(
		`SELECT id, roomid, senderid, content, createdat
		 FROM messages
		 WHERE roomid = $1
		 ORDER BY createdat ASC`,
		roomID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*message.Message

	for rows.Next() {
		var (
			id        string
			roomID    string
			senderID string
			content  string
			created  time.Time
		)

		if err := rows.Scan(&id, &roomID, &senderID, &content, &created); err != nil {
			return nil, err
		}

		messages = append(messages,
			message.NewMessage(id, roomID, senderID, content, created),
		)
	}

	return messages, nil
}
