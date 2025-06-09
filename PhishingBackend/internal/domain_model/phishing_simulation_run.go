package domain_model

import (
	"github.com/google/uuid"
)

type PhishingSimulationRun struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk                   uuid.UUID
	User                     *User `gorm:"foreignKey:UserFk"`
	TemplateFk               uuid.UUID
	Template                 *PhishingSimulationContentTemplate          `gorm:"foreignKey:TemplateFk"`
	RecognitionFeatureValues []PhishingSimulationRecognitionFeatureValue `gorm:"many2many:phishing_simulation_run_recognition_feature_value;"`
	EmailFk                  *uuid.UUID
	Email                    *Email `gorm:"foreignKey:EmailFk"`
}

type PhishingSimulationRunPatch struct {
	ID      uuid.UUID
	EmailFk *uuid.UUID
}
