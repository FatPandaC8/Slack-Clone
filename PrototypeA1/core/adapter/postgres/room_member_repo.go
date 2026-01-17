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
			SELECT 1 FROM room_members WHERE roomid=$1 AND userid=$2
		)`,
		roomID, userID,
	).Scan(&exists)

	return exists, err
}

func (r *RoomMemberRepository) AddMember(roomID, userID string, role string) error {
	_, err := r.db.Exec(
		`INSERT INTO room_members (roomid, userid, role, joined_at)
		 VALUES ($1, $2, $3, NOW())`,
		roomID, userID, role,
	)
	return err
}