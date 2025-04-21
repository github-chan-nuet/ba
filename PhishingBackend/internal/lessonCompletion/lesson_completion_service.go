package lessonCompletion

import (
	"github.com/google/uuid"
	"phishing_backend/internal/domain/db"
	"phishing_backend/internal/domain/model"
	"time"
)

type IsNewEntry bool

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	Create(courseId, lessonId, userId uuid.UUID) (IsNewEntry, error)
}

type ServiceImpl struct {
	Repo db.LessonCompletionRepository
}

func (c *ServiceImpl) Create(courseId, lessonId, userId uuid.UUID) (IsNewEntry, error) {
	lc := model.LessonCompletion{
		LessonId: lessonId,
		CourseId: courseId,
		UserFk:   userId,
		Time:     time.Now().UTC(),
	}
	count, err := c.Repo.Create(&lc)
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
