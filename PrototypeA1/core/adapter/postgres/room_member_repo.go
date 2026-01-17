package postgres

import "database/sql"

type RoomMemberRepository struct {
	db *sql.DB
}

func NewRoomMemberRepository(db *sql.DB) *RoomMemberRepository {
	return &RoomMemberRepository{
		db: db,
	}
}

func (r *RoomMemberRepository) IsMember(roomID, userID string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS (
			SELECT 1 FROM room_members WHERE room_id=$1 AND user_id=$2
		)`,
		roomID, userID,
	).Scan(&exists)

	return exists, err
}

func (r *RoomMemberRepository) AddMember(roomID, userID string, role string) error {
	_, err := r.db.Exec(
		`INSERT INTO room_members (room_id, user_id, role, joined_at)
		 VALUES ($1, $2, $3, NOW())`,
		roomID, userID, role,
	)
	return err
}