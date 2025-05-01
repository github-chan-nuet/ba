package domain_model

import (
	"github.com/google/uuid"
	"time"
)

type Reminder struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk   uuid.UUID
	User     *User     `gorm:"foreignKey:UserFk"`
	SentTime time.Time `gorm:"type:date"`
	Count    uint      `gorm:"type:integer"`
}
