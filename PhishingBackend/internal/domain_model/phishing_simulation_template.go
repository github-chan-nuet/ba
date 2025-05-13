package domain_model

import "github.com/google/uuid"

type PhishingSimulationTemplate struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Template string
}
