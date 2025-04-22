package model

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	Firstname string
	Lastname  string
	Password  []byte `gorm:"type:bytea"`
	Email     string
}
