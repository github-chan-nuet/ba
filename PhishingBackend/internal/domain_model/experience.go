package domain_model

const LessonCompletionGain = 100

type UserExperience struct {
	TotalExperience int
	Level           int
}

type ExperienceGain struct {
	NewExperienceGained int
	TotalExperience     int
	NewLevel            *int
}
