package out

import "core/domain/message"

type MessageBroadcaster interface {
	Broadcast(roomID string, msg *message.Message) error
}