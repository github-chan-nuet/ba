package services

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/domain_model"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var _ PhishingEmailGenerationService = (*PhishingEmailGenerationServiceImpl)(nil)

type PhishingEmailGenerationService interface {
	GenerateEmail(run *domain_model.PhishingSimulationRun) *domain_model.Email
}

type PhishingEmailGenerationServiceImpl struct {
}

func (s *PhishingEmailGenerationServiceImpl) GenerateEmail(run *domain_model.PhishingSimulationRun) *domain_model.Email {
	domainValue := getRecognitionFeatureValueToApply("Domain", run)
	if domainValue == nil {
		slog.Error("Missing Feature Value for Feature: Domain")
		return nil
	}

	subject, err := parseTemplate(run.Template.Subject, run)
	if err != nil {
		slog.Error("Failed to parse subject")
		return nil
	}

	content, err := parseTemplate(run.Template.Content, run)
	if err != nil {
		slog.Error("Failed to parse content")
		return nil
	}

	email := domain_model.Email{
		ID:        uuid.New(),
		Sender:    "info@" + domainValue.Value,
		Recipient: run.User.Email,
		Subject:   subject,
		Content:   content,
	}

	return &email
}

func parseTemplate(input string, run *domain_model.PhishingSimulationRun) (string, error) {
	// Find outermost {{...}} block
	re := regexp.MustCompile(`{{(.*?)}}`)
	for {
		matches := re.FindAllStringSubmatch(input, -1)
		if len(matches) == 0 {
			break
		}

		for _, match := range matches {
			full := match[0]
			body := match[1]

			// Case 1: Literal argument: Handler{Arg}
			if strings.Contains(body, "{") && strings.HasSuffix(body, "}") {
				openBrace := strings.Index(body, "{")
				handlerName := body[:openBrace]
				arg := body[openBrace+1 : len(body)-1]

				handler, ok := placeholderHandlers[handlerName]
				if !ok {
					return "", errors.New("Unknown handler: " + handlerName)
				}

				result, err := handler(arg, run)
				if err != nil {
					return "", err
				}

				input = strings.Replace(input, full, result, 1)
				continue
			}

			// Case 2: Nested Placeholder
			resolvedInner, err := parseTemplate(body, run)
			if err != nil {
				return "", err
			}

			handlerName := resolvedInner
			arg := ""

			if strings.Contains(resolvedInner, " ") {
				parts := strings.SplitN(resolvedInner, " ", 2)
				handlerName = parts[0]
				arg = parts[1]
			}

			handler, ok := placeholderHandlers[handlerName]
			if !ok {
				return "", errors.New("Unknown handler: " + handlerName)
			}

			result, err := handler(arg, run)
			if err != nil {
				return "", err
			}

			input = strings.Replace(input, full, result, 1)
		}
	}

	return input, nil
}

type PlaceholderHandler func(arg string, run *domain_model.PhishingSimulationRun) (string, error)

var placeholderHandlers = map[string]PlaceholderHandler{
	"RecognitionFeature": handleRecognitionFeature,
	"EducationLink":      handleEducationLink,
}

func handleRecognitionFeature(arg string, run *domain_model.PhishingSimulationRun) (string, error) {
	recognitionFeatureValue := getRecognitionFeatureValueToApply(arg, run)
	if recognitionFeatureValue != nil {
		return recognitionFeatureValue.Value, nil
	}
	return "", errors.New("Missing Feature Value for Feature: " + arg)
}

func handleEducationLink(arg string, run *domain_model.PhishingSimulationRun) (string, error) {
	domainFeatureValue := getRecognitionFeatureValueToApply("Domain", run)
	if domainFeatureValue != nil {
		return "https://www." + domainFeatureValue.Value + "?r=" + run.ID.String(), nil
	}
	return "", errors.New("Missing Feature Value for Feature: Domain")
}

func getRecognitionFeatureValueToApply(recognitionFeatureName string, run *domain_model.PhishingSimulationRun) *domain_model.PhishingSimulationRecognitionFeatureValue {
	for _, fv := range run.RecognitionFeatureValues {
		if fv.RecognitionFeature.Name == recognitionFeatureName {
			return &fv
		}
	}
	return nil
}
