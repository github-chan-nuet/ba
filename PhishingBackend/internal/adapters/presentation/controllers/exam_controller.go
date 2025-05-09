package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"phishing_backend/internal/adapters/presentation/mappers"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"phishing_backend/internal/domain_services/services"
)

type ExamController struct {
	Authenticator         services.Authenticator
	ExamRepository        repositories.ExamRepository
	ExamCompletionService services.ExamCompletionService
}

func (e *ExamController) GetExam(w http.ResponseWriter, r *http.Request) {
	examIdStr := r.PathValue("examId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exam, err := e.ExamRepository.Get(examId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	examDto := e.mapToExam(exam, examId)
	writeJsonResponse(w, http.StatusOK, examDto)
}

func (e *ExamController) CompleteExam(w http.ResponseWriter, r *http.Request) {
	examIdStr := r.PathValue("examId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authUserId, err := e.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	var answers []api.QuestionCompletion
	err = json.NewDecoder(r.Body).Decode(&answers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expGain, err := e.ExamCompletionService.CompleteExam(authUserId, examId, mappers.ToQuestionCompletionDto(&answers))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	writeJsonResponse(w, http.StatusOK, mappers.ToApiExpGain(expGain))
}

func (e *ExamController) mapToExam(exam *domain_model.Exam, examId uuid.UUID) *api.Exam {
	dtoExam := api.Exam{Id: examId}
	dtoQuestions := make([]api.Question, 0, len(exam.Questions))
	for _, question := range exam.Questions {
		numCorrectAnswers := 0
		dtoAnswers := make([]api.Answer, 0, len(question.Answers))
		for _, answer := range question.Answers {
			if answer.IsCorrect {
				numCorrectAnswers++
			}
			dtoAnswer := api.Answer{
				Answer: answer.Answer,
				Id:     answer.ID,
			}
			dtoAnswers = append(dtoAnswers, dtoAnswer)
		}

		dtoQuestion := api.Question{
			Answers:  dtoAnswers,
			Id:       question.ID,
			Question: question.Question,
		}
		if numCorrectAnswers == 1 {
			dtoQuestion.Type = api.SingleChoice
		} else {
			dtoQuestion.Type = api.MultipleChoice
		}
		dtoQuestions = append(dtoQuestions, dtoQuestion)
	}
	dtoExam.Questions = dtoQuestions
	return &dtoExam
}
