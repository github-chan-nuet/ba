package domain_model

import (
	"github.com/google/uuid"
	"time"
)

type PhishingSimulationRun struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk     uuid.UUID
	User       *User     `gorm:"foreignKey:UserFk"`
	SentTime   time.Time `gorm:"type:date"`
	OpenedAt   time.Time `gorm:"type:date"`
	TemplateFk uuid.UUID
	Template   *PhishingSimulationTemplate `gorm:"foreignKey:TemplateFk"`
}
