package domain_model

import "github.com/google/uuid"

type ExamQuestionAnswer struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	QuestionFk uuid.UUID
	Question   *ExamQuestion `gorm:"foreignKey:QuestionFk"`
	Answer     string
	IsCorrect  bool `gorm:"type:boolean"`
}
