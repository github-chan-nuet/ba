package controllers

import (
	"encoding/json"
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"phishing_backend/internal/adapters/presentation/mappers"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"phishing_backend/internal/domain_services/services"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ExamController struct {
	Authenticator         services.Authenticator
	ExamRepository        repositories.ExamRepository
	ExamCompRepo          repositories.ExamCompletionRepository
	ExamCompletionService services.ExamCompletionService
}

func (e *ExamController) GetExams(w http.ResponseWriter, _ *http.Request) {
	exams, err := e.ExamRepository.GetAll()
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	apiExams := e.mapToExams(exams)
	writeJsonResponse(w, http.StatusOK, &apiExams)
}

func (e *ExamController) GetExam(w http.ResponseWriter, r *http.Request) {
	examId, err := getPathVariable(r, "examId")
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	exam, err := e.ExamRepository.Get(examId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	examDto := e.mapToExam(exam)
	writeJsonResponse(w, http.StatusOK, examDto)
}

func (e *ExamController) CompleteExam(w http.ResponseWriter, r *http.Request) {
	examId, err := getPathVariable(r, "examId")
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
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

func (e *ExamController) GetCompletedExam(w http.ResponseWriter, r *http.Request) {
	examId, err := getPathVariable(r, "examId")
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	authUserId, err := e.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}

	examComp, err := e.ExamCompRepo.GetCompletedExam(authUserId, examId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	apiExamComp := e.mapToCompletedExam(examComp)
	writeJsonResponse(w, http.StatusOK, &apiExamComp)
}

func (e *ExamController) mapToExams(exams *[]domain_model.Exam) *[]api.Exam {
	dtoExams := make([]api.Exam, len(*exams))
	for i, exam := range *exams {
		dtoExams[i] = *e.mapToExam(&exam)
	}
	return &dtoExams
}

func (e *ExamController) mapToExam(exam *domain_model.Exam) *api.Exam {
	dtoExam := api.Exam{
		Id:          (*exam).ID,
		Title:       &(*exam).Title,
		Description: &(*exam).Description,
	}
	dtoQuestions := make([]api.Question, len((*exam).Questions))
	for i, question := range exam.Questions {
		numCorrectAnswers := 0
		dtoAnswers := make([]api.Answer, len(question.Answers))
		for j, answer := range question.Answers {
			if answer.IsCorrect {
				numCorrectAnswers++
			}
			dtoAnswer := api.Answer{
				Answer: answer.Answer,
				Id:     answer.ID,
			}
			dtoAnswers[j] = dtoAnswer
		}

		dtoQuestion := api.Question{
			Answers:  dtoAnswers,
			Id:       question.ID,
			Question: question.Question,
		}
		if numCorrectAnswers == 1 {
			dtoQuestion.Type = api.QuestionTypeSingleChoice
		} else {
			dtoQuestion.Type = api.QuestionTypeMultipleChoice
		}
		dtoQuestions[i] = dtoQuestion
	}
	dtoExam.Questions = dtoQuestions
	return &dtoExam
}

func (e *ExamController) mapToCompletedExam(exComp *domain_model.ExamCompletion) *api.CompletedExam {
	if exComp == nil || exComp.Exam == nil {
		return nil
	}

	qIdQuestion := make(map[uuid.UUID]api.CompletedQuestion)
	// fill out user answers
	for _, answer := range exComp.Answers {
		qId := answer.Answer.QuestionFk
		entry, ok := qIdQuestion[qId]
		if !ok {
			entry = api.CompletedQuestion{
				Id:          qId,
				Question:    answer.Answer.Question.Question,
				UserAnswers: make([]uuid.UUID, 4),
			}
		}
		entry.UserAnswers = append(qIdQuestion[qId].UserAnswers, answer.Answer.ID)
		qIdQuestion[qId] = entry
	}
	// fill out actual answers
	for _, q := range exComp.Exam.Questions {
		nCorrectAnswers := 0
		answers := make([]api.AnswerWithCorrection, len(q.Answers))
		for i, answer := range q.Answers {
			answers[i] = api.AnswerWithCorrection{
				Answer:    answer.Answer,
				Id:        answer.ID,
				IsCorrect: answer.IsCorrect,
			}
			if answer.IsCorrect {
				nCorrectAnswers++
			}
		}
		entry, _ := qIdQuestion[q.ID]
		entry.Answers = answers
		if nCorrectAnswers == 1 {
			entry.Type = api.CompletedQuestionTypeSingleChoice
		} else {
			entry.Type = api.CompletedQuestionTypeMultipleChoice
		}
		qIdQuestion[q.ID] = entry
	}
	// map from map to array
	qs := make([]api.CompletedQuestion, len(qIdQuestion))
	i := 0
	for _, q := range qIdQuestion {
		qs[i] = q
		i++
	}
	apiExamComp := api.CompletedExam{
		CompletedAt: openapi_types.Date{Time: exComp.CompletedAt},
		Questions:   qs,
	}
	return &apiExamComp
}
