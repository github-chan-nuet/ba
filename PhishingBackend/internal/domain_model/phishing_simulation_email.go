package domain_model

import (
	"time"

	"github.com/google/uuid"
)

type PhishingSimulationEmail struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk                   uuid.UUID
	User                     *User `gorm:"foreignKey:UserFk"`
	TemplateFk               uuid.UUID
	Template                 *PhishingSimulationContentTemplate           `gorm:"foreignKey:TemplateFk"`
	RecognitionFeatureValues []*PhishingSimulationRecognitionFeatureValue `gorm:"many2many:phishing_simulation_email_recognition_feature_values;"`
	SentAt                   time.Time                                    `gorm:"type:date"`
	OpenedAt                 time.Time                                    `gorm:"type:date"`
}
