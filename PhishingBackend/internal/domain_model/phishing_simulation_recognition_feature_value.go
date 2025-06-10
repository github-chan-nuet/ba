package domain_model

import "github.com/google/uuid"

type PhishingSimulationRecognitionFeatureValue struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Value                string
	Level                int
	UserInstruction      string
	RecognitionFeatureFk uuid.UUID
	RecognitionFeature   *PhishingSimulationRecognitionFeature `gorm:"foreignKey:RecognitionFeatureFk"`
	ContentCategoryFk    uuid.UUID
	ContentCategory      *PhishingSimulationContentCategory `gorm:"foreignKey:ContentCategoryFk"`
}
