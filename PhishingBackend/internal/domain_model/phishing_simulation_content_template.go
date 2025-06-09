package domain_model

import (
	"errors"
	"log/slog"
	"math"
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

type ScoredTemplate struct {
	Score                    float32
	Template                 *PhishingSimulationContentTemplate
	RecognitionFeatureValues []PhishingSimulationRecognitionFeatureValue
}

func (template *PhishingSimulationContentTemplate) GetScoredCombinations(
	vulnerabilities []PhishingSimulationUserVulnerability,
	recognitionFeatureDefinitions []PhishingSimulationRecognitionFeature,
) []ScoredTemplate {
	var scoredCombinations []ScoredTemplate
	recognitionFeatures, err := template.getApplicableRecognitionFeatures(recognitionFeatureDefinitions)
	if err != nil {
		slog.Error("Template Error", "error", err)
		return scoredCombinations
	}

	combinations := utils.Combinations(recognitionFeatures.List(), 2)
	for _, comb := range combinations {
		// Iterate over applicableRecognitionFeatures
		// if in comb then select the respective level x > 1
		// else select lowest level
		var featureValuesComb []PhishingSimulationRecognitionFeatureValue
		var score float32 = 0
		for _, feature := range recognitionFeatures.List() {
			level := 0
			if utils.Find(comb, func(f PhishingSimulationRecognitionFeature) bool {
				return f.ID == feature.ID
			}) != nil {
				vulnerability := utils.Find(
					vulnerabilities,
					func(vuln PhishingSimulationUserVulnerability) bool {
						return *vuln.ContentCategory == *template.ContentCategory && vuln.RecognitionFeature.ID == feature.ID
					})
				if vulnerability != nil {
					score += vulnerability.Score
					level = int(math.Round(float64(vulnerability.Score)))
				} else {
					mostMatchingVulnerabilities := utils.FindAll(
						vulnerabilities,
						func(vuln PhishingSimulationUserVulnerability) bool {
							return vuln.RecognitionFeature.ID == feature.ID
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
						level = int(math.Round(float64(arithMean)))
					}
				}
				if level < 1 {
					level = 1
				}
			}

			if feature.RecognitionFeatureValues != nil && len(*feature.RecognitionFeatureValues) > 0 {
				mostSuitableValue := (*feature.RecognitionFeatureValues)[0]
				for _, value := range (*feature.RecognitionFeatureValues)[1:] {
					if math.Abs(float64(level-value.Level)) < math.Abs(float64(level-mostSuitableValue.Level)) {
						mostSuitableValue = value
					}
				}
				featureValuesComb = append(featureValuesComb, mostSuitableValue)
			}
		}

		scoredCombination := ScoredTemplate{
			Score:                    score,
			Template:                 template,
			RecognitionFeatureValues: featureValuesComb,
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
