package user

import "time"

type User struct {
	id 				ID
	name 			string
	email 			string
	passwordHash 	string
	createdAt 		time.Time
}