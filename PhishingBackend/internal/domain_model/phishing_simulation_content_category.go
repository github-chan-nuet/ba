package domain_model

import "github.com/google/uuid"

type PhishingSimulationContentCategory struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string
}
