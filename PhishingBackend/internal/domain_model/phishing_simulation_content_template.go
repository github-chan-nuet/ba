package domain_model

import "github.com/google/uuid"

type PhishingSimulationContentTemplate struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;"`
	Subject           string
	Content           string
	ContentCategoryFk uuid.UUID
	ContentCategory   *PhishingSimulationContentCategory `gorm:"foreignKey:ContentCategoryFk"`
}
