package domain_model

const (
	LessonCompletionGain = 100
	ExamCompletionGain   = 1000
)

type UserExperience struct {
	TotalExperience int
	Level           int
}

type ExperienceGain struct {
	NewExperienceGained int
	TotalExperience     int
	NewLevel            *int
}
