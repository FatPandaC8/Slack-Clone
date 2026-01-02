package dto

import "chat-core-go/domain/user"

type CreateChannelCommand struct {
	ChannelID 		string
	Members  		[]user.ID
}