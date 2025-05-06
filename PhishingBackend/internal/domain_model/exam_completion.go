package domain_model

import (
	"github.com/google/uuid"
	"time"
)

type ExamCompletion struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserFk         uuid.UUID
	User           *User `gorm:"foreignKey:UserFk"`
	ExamFk         uuid.UUID
	Exam           *Exam     `gorm:"foreignKey:ExamFk"`
	CompletionTime time.Time `gorm:"type:date"`
	Answers        []ExamCompletionAnswer
}
