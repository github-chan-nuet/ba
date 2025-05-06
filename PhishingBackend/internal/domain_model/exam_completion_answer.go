package domain_model

import "github.com/google/uuid"

type ExamCompletionAnswer struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	ExamFk   uuid.UUID
	Exam     *Exam `gorm:"foreignKey:ExamFk"`
	AnswerFk uuid.UUID
	Answer   *ExamQuestionAnswer `gorm:"foreignKey:AnswerFk"`
}
