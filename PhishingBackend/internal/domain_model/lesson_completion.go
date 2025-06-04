package domain_model

import (
	"github.com/google/uuid"
	"time"
)

type LessonCompletion struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	CourseId uuid.UUID `gorm:"type:uuid"`
	LessonId uuid.UUID `gorm:"type:uuid"`
	UserFk   uuid.UUID
	User     *User     `gorm:"foreignKey:UserFk"`
	Time     time.Time `gorm:"type:date"`
}

type LastLessonCompletion struct {
	UserID uuid.UUID
	Time   time.Time
}
