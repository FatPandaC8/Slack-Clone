package ws

import (
	"core/domain/message"
	"fmt"
)

type Broadcaster struct {}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{}
}

func (b *Broadcaster) Broadcast(roomID string, msg *message.Message) error {
	fmt.Printf("SERVER: RoomID (%s) - Message (%s)\n", roomID, msg.Content())
	return nil
}