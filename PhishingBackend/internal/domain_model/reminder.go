package domain_model

import (
	"time"

	"github.com/google/uuid"
)

type Reminder struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk        uuid.UUID
	User          *User     `gorm:"foreignKey:UserFk"`
	SentTime      time.Time `gorm:"type:date"`
	Count         int
	TemplateFk    int
	EmailTemplate *ReminderEmailTemplate `gorm:"foreignKey:TemplateFk"`
}
