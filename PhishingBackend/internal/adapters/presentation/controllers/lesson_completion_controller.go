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

type LessonCompletionController struct {
	LessonCompletionService    services.LessonCompletionService
	LessonCompletionRepository repositories.LessonCompletionRepository
	Authenticator              services.Authenticator
	ExperienceService          services.ExperienceService
}

func (c *LessonCompletionController) CreateLessonCompletion(w http.ResponseWriter, r *http.Request) {
	courseIdStr := r.PathValue("courseId")
	courseId, err := uuid.Parse(courseIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var lesson api.Lesson
	err = json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}

	isNew, err := c.LessonCompletionService.Create(courseId, lesson.LessonId, userId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	if !isNew {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	expGain, err := c.ExperienceService.GetExperienceGainOfLessonCompletion(userId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	writeJsonResponse(w, http.StatusCreated, mappers.ToApiExpGain(expGain))
}

func (c *LessonCompletionController) GetAllLessonCompletionsOfUser(w http.ResponseWriter, r *http.Request) {
	userId, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}

	completions, err := c.LessonCompletionRepository.GetAllCompletedLessonsInAllCourses(userId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	ccs := toApiCourseCompletions(completions)
	writeJsonResponse(w, http.StatusOK, ccs)
}

func (c *LessonCompletionController) GetLessonCompletionsOfCourseAndUser(w http.ResponseWriter, r *http.Request) {
	courseIdStr := r.PathValue("courseId")
	courseId, err := uuid.Parse(courseIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := c.Authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	lcs, err := c.LessonCompletionRepository.GetLessonCompletionsOfCourseAndUser(userId, courseId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	lessonIds := aggregateLessonIds(lcs)
	writeJsonResponse(w, http.StatusOK, lessonIds)
}

func toApiCourseCompletions(completions []domain_model.LessonCompletion) []api.CourseCompletion {
	courseToLessons := make(map[uuid.UUID][]uuid.UUID)
	for _, c := range completions {
		lessons, ok := courseToLessons[c.CourseId]
		if !ok {
			courseToLessons[c.CourseId] = make([]uuid.UUID, 10)
		}
		courseToLessons[c.CourseId] = append(lessons, c.LessonId)
	}

	ccs := make([]api.CourseCompletion, 0, len(courseToLessons))
	for courseId, lessons := range courseToLessons {
		ccs = append(ccs, api.CourseCompletion{
			CourseId:         courseId,
			CompletedLessons: lessons,
		})
	}
	return ccs
}

func aggregateLessonIds(lcs []domain_model.LessonCompletion) []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(lcs))
	for _, lc := range lcs {
		ids = append(ids, lc.LessonId)
	}
	return ids
}
