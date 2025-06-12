package services

import (
	"errors"
	"github.com/google/uuid"
	"math"
	. "phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"time"
)

var (
	_                      ExamCompletionService = (*ExamCompletionServiceImpl)(nil)
	ErrQuestionNotExisting                       = errors.New("question does not exist")
)

type ExamCompletionService interface {
	CompleteExam(userId, examId uuid.UUID, answers *[]QuestionCompletionDto) (*ExperienceGain, error)
}

type ExamCompletionServiceImpl struct {
	ExamRepo          repositories.ExamRepository
	ExamCompRepo      repositories.ExamCompletionRepository
	ExperienceService ExperienceService
}

func (e *ExamCompletionServiceImpl) CompleteExam(userId, examId uuid.UUID, answers *[]QuestionCompletionDto) (*ExperienceGain, error) {
	exam, err := e.ExamRepo.Get(examId)
	if err != nil {
		return nil, err
	}
	score, err := e.calculateScore(exam, answers)
	if err != nil {
		return nil, err
	}
	exComp := e.createExamCompletion(userId, examId, answers, score)
	err = e.ExamCompRepo.Save(exComp)
	if err != nil {
		return nil, err
	}
	return e.ExperienceService.GetExperienceGainOfExamCompletion(userId, score)
}

func (e *ExamCompletionServiceImpl) createExamCompletion(userId, examId uuid.UUID, comps *[]QuestionCompletionDto, score int) *ExamCompletion {
	answers := make([]ExamCompletionAnswer, 0, len(*comps))
	for _, comp := range *comps {
		for _, answer := range comp.Answers {
			answers = append(answers, ExamCompletionAnswer{
				ID:       uuid.New(),
				AnswerFk: answer,
			})
		}
	}
	ec := ExamCompletion{
		ID:          uuid.New(),
		UserFk:      userId,
		ExamFk:      examId,
		CompletedAt: time.Now().UTC(),
		Answers:     answers,
		Score:       score,
	}
	return &ec
}

// calculateScore calculates the score in the range [0, 100]
func (e *ExamCompletionServiceImpl) calculateScore(exam *Exam, answers *[]QuestionCompletionDto) (int, error) {
	score := float64(0)
	scorePerQ := 100 / float64(len(exam.Questions))
	resps, err := e.createUserAndActualAnswer(exam, answers)
	if err != nil {
		return 0, err
	}
	for _, resp := range *resps {
		singleScore := e.calculateScoreOfQuestion(&resp)
		score += singleScore * scorePerQ
	}
	return int(math.Round(score)), nil
}

// calculateScoreOfQuestion calculates the score in the range [0, 1]
func (e *ExamCompletionServiceImpl) calculateScoreOfQuestion(a *userAndActualAnswer) float64 {
	reductionPerWrongA := 1 / float64(a.totalAnswers)
	wrong := a.getNumberOfWrongOrMissingAnswers()
	score := float64(1) - float64(wrong)*reductionPerWrongA
	return score
}

func (e *ExamCompletionServiceImpl) createUserAndActualAnswer(exam *Exam, qComps *[]QuestionCompletionDto) (*[]userAndActualAnswer, error) {
	qIdMap := make(map[uuid.UUID]userAndActualAnswer, len(exam.Questions))
	// fill correct answers and total answers
	for _, question := range exam.Questions {
		correctAs := make(UuidSet)
		for _, answer := range question.Answers {
			if answer.IsCorrect {
				correctAs[answer.ID] = struct{}{}
			}
		}
		qIdMap[question.ID] = userAndActualAnswer{
			totalAnswers:   len(question.Answers),
			correctAnswers: correctAs,
		}
	}
	// fill user answers
	for _, qComp := range *qComps {
		_, ok := qIdMap[qComp.QuestionId]
		if !ok {
			return nil, ErrQuestionNotExisting
		}
		q := qIdMap[qComp.QuestionId]
		q.userAnswers = toMap(&qComp.Answers)
		qIdMap[qComp.QuestionId] = q
	}
	responses := make([]userAndActualAnswer, len(exam.Questions))
	i := 0
	for _, resp := range qIdMap {
		responses[i] = resp
		i++
	}
	return &responses, nil
}

func toMap[V comparable](array *[]V) map[V]struct{} {
	m := make(map[V]struct{}, len(*array))
	for _, v := range *array {
		m[v] = struct{}{}
	}
	return m
}

type UuidSet map[uuid.UUID]struct{}

type userAndActualAnswer struct {
	totalAnswers   int
	correctAnswers UuidSet
	userAnswers    UuidSet
}

// A: set of answers the user thinks are correct
// B: set of all actual correct answers
// ex: A={1,4,5}, B={1,3} -> |(A - B) U (B - A)| = |{4,5} U {3}| = |{3,4,5}| = 3
func (u *userAndActualAnswer) getNumberOfWrongOrMissingAnswers() int {
	wrong := 0
	// B - A
	for val := range u.correctAnswers {
		if _, exists := u.userAnswers[val]; !exists {
			wrong++
		}
	}
	// A - B
	for val := range u.userAnswers {
		if _, exists := u.correctAnswers[val]; !exists {
			wrong++
		}
	}
	return wrong
}
