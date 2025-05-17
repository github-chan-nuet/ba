package domain_model

import (
	"github.com/google/uuid"
)

type ExamCompletionAnswer struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	ExamCompFk uuid.UUID
	ExamComp   *ExamCompletion `gorm:"foreignKey:ExamCompFk"`
	AnswerFk   uuid.UUID
	Answer     *ExamQuestionAnswer `gorm:"foreignKey:AnswerFk"`
}

type QuestionCompletionDto struct {
	Answers    []uuid.UUID
	QuestionId uuid.UUID
}
