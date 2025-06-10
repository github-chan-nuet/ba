package domain_model

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Sender    string
	Recipient string
	Subject   string
	Content   string
	SentAt    *time.Time `gorm:"type:timestamptz"`
	ClickedAt *time.Time `gorm:"type:timestamptz"`
}

type EmailPatch struct {
	ID        uuid.UUID
	SentAt    *time.Time
	ClickedAt *time.Time
}
