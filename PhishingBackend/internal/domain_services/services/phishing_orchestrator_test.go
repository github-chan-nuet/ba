package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"testing"
	"time"
)

func TestGeneratePhishingRuns(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	m := repositories.NewMockUserRepository(ctrl)
	s := repositories.NewMockPhishingSimulationRepository(ctrl)
	r := NewMockPhishingRunService(ctrl)
	users := []domain_model.User{{
		ID:                               uuid.New(),
		Firstname:                        "A",
		Lastname:                         "B",
		Email:                            "C@C",
		ParticipatesInPhishingSimulation: true,
	}}

	sut := PhishingOrchestratorImpl{
		UserRepository:               m,
		PhishingSimulationRepository: s,
		PhishingRunService:           r,
	}
	now := time.Now().UTC()
	twoWeeksAgo := now.Add(time.Hour * 24 * -14)

	tests := []struct {
		name string
		run  *domain_model.PhishingSimulationRun
	}{
		{
			name: "No latest run",
			run:  nil,
		},
		{
			name: "latest run without email",
			run:  &domain_model.PhishingSimulationRun{},
		},
		{
			name: "latest run with empty Email",
			run: &domain_model.PhishingSimulationRun{
				Email: &domain_model.Email{},
			},
		},
		{
			name: "latest run that was done 2 weeks ago",
			run: &domain_model.PhishingSimulationRun{
				Email: &domain_model.Email{SentAt: &twoWeeksAgo},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.EXPECT().GetLatestRun(users[0].ID).Return(tt.run, nil)
			m.EXPECT().GetUsersForPhishingSimulation().Return(users, nil)
			var capture *domain_model.User
			r.EXPECT().GenerateRun(gomock.Any()).DoAndReturn(func(arg *domain_model.User) error {
				capture = arg
				return nil
			})

			// when
			sut.generatePhishingRuns(now)

			// then
			assert.Equal(t, users[0], *capture)
		})
	}
}
