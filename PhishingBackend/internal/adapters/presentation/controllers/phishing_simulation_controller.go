package controllers

import (
	"net/http"
	"phishing_backend/internal/adapters/presentation/api"
	"phishing_backend/internal/adapters/presentation/error_handling"
	"phishing_backend/internal/domain_model"
	"phishing_backend/internal/domain_services/interfaces/repositories"
	"phishing_backend/internal/domain_services/services"
)

type PhishingSimulationController struct {
	PhishingRunService           services.PhishingRunService
	PhishingSimulationRepository repositories.PhishingSimulationRepository
}

func (c *PhishingSimulationController) GetRun(w http.ResponseWriter, r *http.Request) {
	runId, err := getPathVariable(r, "phishingSimulationRunId")
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	run, err := c.PhishingSimulationRepository.GetRun(runId)
	if err != nil {
		error_handling.WriteErrorDetailResponse(w, err)
		return
	}
	if run.OpenedAt == nil {
		c.PhishingRunService.TrackRunClick(run)
	}
	runDto := toApiPhishingSimulationRun(run)
	writeJsonResponse(w, http.StatusOK, &runDto)
}

func toApiPhishingSimulationRun(run *domain_model.PhishingSimulationRun) *api.PhishingSimulationRun {
	dtoRun := api.PhishingSimulationRun{
		PhishingSimulationRunId: (*run).ID,
	}
	dtoRecognitionFeatureValues := make([]api.PhishingSimulationRecognitionFeatureValue, len((*run).RecognitionFeatureValues))
	for i, recognitionFeatureValue := range run.RecognitionFeatureValues {
		dtoRecognitionFeatureValue := api.PhishingSimulationRecognitionFeatureValue{
			RecognitionFeatureValueId: recognitionFeatureValue.ID,
		}
		dtoRecognitionFeatureValues[i] = dtoRecognitionFeatureValue
	}
	dtoRun.RecognitionFeatureValues = dtoRecognitionFeatureValues
	return &dtoRun
}
