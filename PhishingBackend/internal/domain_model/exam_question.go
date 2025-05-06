package domain_model

import "github.com/google/uuid"

type ExamQuestion struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	ExamFk   uuid.UUID
	Exam     *Exam `gorm:"foreignKey:ExamFk"`
	Question string
	Answers  []ExamQuestionAnswer
}
