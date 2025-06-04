package domain_model

import (
	"errors"
	"log/slog"
	"phishing_backend/internal/utils"
	"regexp"

	"github.com/google/uuid"
)

type PhishingSimulationContentTemplate struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;"`
	Subject           string
	Content           string
	ContentCategoryFk uuid.UUID
	ContentCategory   *PhishingSimulationContentCategory `gorm:"foreignKey:ContentCategoryFk"`
}

type PhishingSimulationContentTemplateResult struct {
	Subject string
	Content string
}

type ScoredTemplate struct {
	Score               float32
	Template            *PhishingSimulationContentTemplate
	RecognitionFeatures []PhishingSimulationRecognitionFeature
}

func (template *PhishingSimulationContentTemplate) GetScoredCombinations(vulnerabilities []PhishingSimulationUserVulnerability, recognitionFeatureDefinitions []PhishingSimulationRecognitionFeature) []ScoredTemplate {
	var scoredCombinations []ScoredTemplate
	recognitionFeatures, err := template.getApplicableRecognitionFeatures(recognitionFeatureDefinitions)
	if err != nil {
		slog.Error("Template Error", "error", err)
		return scoredCombinations
	}

	combinations := utils.Combinations(recognitionFeatures.List(), 2)
	for _, comb := range combinations {
		var score float32 = 0
		for _, feat := range comb {
			vulnerability := utils.Find(
				vulnerabilities,
				func(vuln PhishingSimulationUserVulnerability) bool {
					return *vuln.ContentCategory == *template.ContentCategory && *vuln.RecognitionFeature == feat
				})
			if vulnerability != nil {
				score += vulnerability.Score
			} else {
				mostMatchingVulnerabilities := utils.FindAll(
					vulnerabilities,
					func(vuln PhishingSimulationUserVulnerability) bool {
						return *vuln.RecognitionFeature == feat
					})

				var arithMean float32 = 0
				for _, mostMatchingVuln := range mostMatchingVulnerabilities {
					arithMean += mostMatchingVuln.Score
				}
				if len(mostMatchingVulnerabilities) > 0 {
					arithMean = arithMean / float32(len(mostMatchingVulnerabilities))
				}
				if arithMean == 0 {
					score += 1
				} else {
					score += arithMean
				}
			}
		}

		scoredCombination := ScoredTemplate{
			Score:               score,
			Template:            template,
			RecognitionFeatures: comb,
		}
		scoredCombinations = append(scoredCombinations, scoredCombination)
	}
	return scoredCombinations
}

func (template *PhishingSimulationContentTemplate) getApplicableRecognitionFeatures(recognitionFeatureDefinitions []PhishingSimulationRecognitionFeature) (*utils.Set[PhishingSimulationRecognitionFeature], error) {
	resultSet := utils.NewSet[PhishingSimulationRecognitionFeature]()

	applicableInContent := extractRecognitionFeaturesUsed(template.Content)
	applicableInSubject := extractRecognitionFeaturesUsed(template.Subject)
	applicableTotal := applicableInContent.Union(&applicableInSubject)

	for _, def := range recognitionFeatureDefinitions {
		if def.IsAlwaysApplicable {
			resultSet.Add(def)
		}
	}

	for _, feat := range applicableTotal.List() {
		defition := utils.Find(recognitionFeatureDefinitions, func(def PhishingSimulationRecognitionFeature) bool {
			return def.Name == feat
		})

		if defition == nil {
			err := errors.New("Undefined RecognitionFeature used in template")
			return nil, err
		}
		resultSet.Add(*defition)
	}
	return resultSet, nil
}

func extractRecognitionFeaturesUsed(v string) utils.Set[string] {
	var values utils.Set[string]

	re := regexp.MustCompile(`\{RecognitionFeature\{([^}]+)\}\}`)
	matches := re.FindAllStringSubmatch(v, -1)

	for _, match := range matches {
		if len(match) > 1 {
			values.Add(match[1])
		}
	}
	return values
}
