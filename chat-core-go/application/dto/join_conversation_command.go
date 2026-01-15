package dto

import (
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
)

type JoinConversationCommand struct {
	Principal  *identity.Principal
	InviteCode valueobject.InviteCode
}