package postgres

import (
	"core/domain/room"
	"database/sql"
	"time"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) Save(room *room.Room) error {
	_, err := r.db.Exec(
		`INSERT INTO rooms (roomid, name, invite_code, created_at)
		VALUES ($1, $2, $3, $4)`,
		room.ID(), room.Name(), room.InviteCode(), time.Now(),
	)
	return err
}

func (r *RoomRepository) FindByID(roomID string) (*room.Room, error) {
	row := r.db.QueryRow(
		`SELECT roomid, name, invite_code, created_at
		 FROM rooms
		 WHERE roomid = $1`,
		roomID,
	)

	var (
		e_roomID    string
		name      string
		invite    string
		createdAt time.Time
	)

	if err := row.Scan(&e_roomID, &name, &invite, &createdAt); err != nil {
		return nil, err
	}

	return room.NewRoom(e_roomID, name, invite, createdAt), nil
}

func (r *RoomRepository) FindByInviteCode(code string) (*room.Room, error) {
	row := r.db.QueryRow(
		`SELECT roomid, name, invite_code, created_at
		 FROM rooms
		 WHERE invite_code = $1`,
		code,
	)

	var (
		roomID    string
		name      string
		invite    string
		createdAt time.Time
	)

	if err := row.Scan(&roomID, &name, &invite, &createdAt); err != nil {
		return nil, err
	}

	return room.NewRoom(roomID, name, invite, createdAt), nil
}