package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/internal/application/services"
	"phishing_backend/internal/infrastructure/presentation/api"
)

type LessonCompletionController struct {
	LessonCompletionService services.LessonCompletionService
	Authenticator           services.Authenticator
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
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	isNew, err := c.LessonCompletionService.Create(courseId, lesson.LessonId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if isNew {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	//api.ExperienceGain{}
}
