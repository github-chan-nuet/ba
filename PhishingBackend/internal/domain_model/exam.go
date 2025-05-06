package domain_model

import "github.com/google/uuid"

type Exam struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Answers []ExamQuestion
}
