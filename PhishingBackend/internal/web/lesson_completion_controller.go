package web

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"phishing_backend/api"
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/lessonCompletion"
)

var lessonCompletionService lessonCompletion.Service = &lessonCompletion.ServiceImpl{
	Repo: &db.LessonCompletionRepositoryImpl{},
}

func createLessonCompletion(w http.ResponseWriter, r *http.Request) {
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
	user, err := authenticator.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	isNew, err := lessonCompletionService.Create(courseId, lesson.LessonId, user.ID)
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
