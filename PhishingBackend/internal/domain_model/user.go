package domain_model

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Firstname string
	Lastname  string
	Password  []byte `gorm:"type:bytea"`
	Email     string
}

type UserPatchDto struct {
	Email     *string
	Firstname *string
	Lastname  *string
	Password  *string
}

type UserPostDto struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
}
