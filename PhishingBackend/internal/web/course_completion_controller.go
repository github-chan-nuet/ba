package web

import (
	"encoding/json"
	"net/http"
	"phishing_backend/api"
	"phishing_backend/internal/courseCompletion"
	"phishing_backend/internal/domain/db"
)

var courseCompletionService courseCompletion.Service = &courseCompletion.ServiceImpl{
	Repo: &db.CourseCompletionRepositoryImpl{},
}

func createLessonCompletion(w http.ResponseWriter, r *http.Request) {
	courseId := r.PathValue("courseId")
	var lesson api.Lesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
