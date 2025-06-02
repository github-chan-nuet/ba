package domain_model

import (
	"net/mail"
	"phishing_backend/internal/domain_model/validation"

	"github.com/google/uuid"
)

type User struct {
	ID                               uuid.UUID `gorm:"type:uuid;primary_key;"`
	Firstname                        string
	Lastname                         string
	Password                         []byte `gorm:"type:bytea"`
	Email                            string
	ParticipatesInPhishingSimulation bool
}

type UserPatch struct {
	ID                               uuid.UUID
	Email                            *string
	Firstname                        *string
	Lastname                         *string
	Password                         *[]byte
	ParticipatesInPhishingSimulation *bool
}

type UserPatchDto struct {
	Email                            *string
	Firstname                        *string
	Lastname                         *string
	Password                         *string
	ParticipatesInPhishingSimulation *bool
}

func (u *UserPatchDto) Validate() error {
	vErr := validation.NewValidationError()
	if u.Email == nil && u.Firstname == nil && u.Lastname == nil && u.Password == nil {
		vErr.Add("#/", validation.NoFieldSet)
		return vErr
	}
	if u.Email != nil && !isValidEmail(*u.Email) {
		vErr.Add("#/email", validation.InvalidEmail)
	}
	if vErr.HasErr() {
		return vErr
	}
	return nil
}

type UserPostDto struct {
	Email                            string
	Firstname                        string
	Lastname                         string
	Password                         string
	ParticipatesInPhishingSimulation *bool
}

func (u *UserPostDto) Validate() error {
	vErr := validation.NewValidationError()
	if u.Firstname == "" {
		vErr.Add("#/firstname", validation.Mandatory)
	}
	if u.Lastname == "" {
		vErr.Add("#/lastname", validation.Mandatory)
	}
	if u.Email == "" {
		vErr.Add("#/email", validation.Mandatory)
	} else if !isValidEmail(u.Email) {
		vErr.Add("#/email", validation.InvalidEmail)
	}
	if u.Password == "" {
		vErr.Add("#/password", validation.Mandatory)
	}
	if vErr.HasErr() {
		return vErr
	}
	return nil
}

func isValidEmail(email string) bool {
	addr, err := mail.ParseAddress(email)
	return err == nil && addr.Address == email
}
