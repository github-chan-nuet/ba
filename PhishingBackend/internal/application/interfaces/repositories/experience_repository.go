package repositories

import "github.com/google/uuid"

type ExperienceRepository interface {
	GetTotalExperience(userId uuid.UUID) (int, error)
}
