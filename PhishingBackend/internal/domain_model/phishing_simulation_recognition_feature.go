package domain_model

import "github.com/google/uuid"

type PhishingSimulationRecognitionFeature struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name                     string
	IsAlwaysApplicable       bool
	Title                    string
	UserInstruction          string
	RecognitionFeatureValues *[]PhishingSimulationRecognitionFeatureValue `gorm:"foreignKey:RecognitionFeatureFk"`
}
