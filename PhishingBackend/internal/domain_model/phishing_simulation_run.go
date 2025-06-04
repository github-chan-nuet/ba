package domain_model

import (
	"time"

	"github.com/google/uuid"
)

type PhishingSimulationRun struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk                   uuid.UUID
	User                     *User `gorm:"foreignKey:UserFk"`
	TemplateFk               uuid.UUID
	Template                 *PhishingSimulationContentTemplate           `gorm:"foreignKey:TemplateFk"`
	RecognitionFeatureValues []*PhishingSimulationRecognitionFeatureValue `gorm:"many2many:phishing_simulation_run_recognition_feature_value;"`
	SentAt                   *time.Time                                   `gorm:"type:timestamptz"`
	OpenedAt                 *time.Time                                   `gorm:"type:timestamptz"`
}

type PhishingSimulationRunPatch struct {
	ID       uuid.UUID
	SentAt   *time.Time
	OpenedAt *time.Time
}
