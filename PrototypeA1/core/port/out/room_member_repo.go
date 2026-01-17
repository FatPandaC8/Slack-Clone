package out

type RoomMemberRepository interface {
	IsMember(roomID, userID string) (bool, error)
	AddMember(roomID, userID string, role string) error
}