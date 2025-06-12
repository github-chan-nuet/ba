package domain_model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPowerSet(t *testing.T) {
	// given
	s := []int{1, 2, 3}

	// when
	c := getPowerSet(s, 2)

	// then
	wantSet := [][]int{{1}, {1, 2}, {1, 3}, {2}, {2, 3}, {3}}
	assert.ElementsMatch(t, wantSet, c)
}

func TestExtractRecognitionFeaturesUsed(t *testing.T) {
	// given
	input := "Lorem Ipsum {RecognitionFeature{Voice}} {RecognitionFeature{A}} {RecognitionFeature{Aa}}"

	// when
	s := extractRecognitionFeaturesUsed(input)

	// then
	assert.ElementsMatch(t, []string{"Voice", "A", "Aa"}, s.List())
}

func TestFind(t *testing.T) {
	// given
	ints := []int{1, 3, 5, 7, 2}
	isEven := func(i int) bool { return i%2 == 0 }

	// when
	result := find(ints, isEven)

	// then
	assert.Equal(t, 2, *result)
}

func TestFindNoMatch(t *testing.T) {
	// given
	ints := []int{1, 3, 5, 7}
	isEven := func(i int) bool { return i%2 == 0 }

	// when
	result := find(ints, isEven)

	// then
	assert.Nil(t, result)
}

func TestFindAll(t *testing.T) {
	// given
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(i int) bool { return i%2 == 0 }

	// when
	result := findAll(ints, isEven)

	// then
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, result)
}

func TestFindAllNoMatch(t *testing.T) {
	// given
	ints := []int{1, 3, 5, 7, 9}
	isEven := func(i int) bool { return i%2 == 0 }

	// when
	result := findAll(ints, isEven)

	// then
	assert.Empty(t, result)
}

func TestGetApplicableRecognitionFeaturesSuccess(t *testing.T) {
	// given
	template := PhishingSimulationContentTemplate{
		Subject: "{RecognitionFeature{Name1}}",
		Content: "{RecognitionFeature{Name2}}",
	}
	def1 := PhishingSimulationRecognitionFeature{Name: "Name1", IsAlwaysApplicable: false}
	def2 := PhishingSimulationRecognitionFeature{Name: "Name2", IsAlwaysApplicable: false}
	def3 := PhishingSimulationRecognitionFeature{Name: "Always", IsAlwaysApplicable: true}

	// when
	features, err := template.getApplicableRecognitionFeatures([]PhishingSimulationRecognitionFeature{def1, def2, def3})

	// then
	assert.NoError(t, err)
	assert.Equal(t, 3, len(features.List()))
}

func TestGetApplicableRecognitionFeaturesUndefinedFeature(t *testing.T) {
	// given
	template := PhishingSimulationContentTemplate{
		Content: "{RecognitionFeature{Unknown}}",
	}

	// when
	features, err := template.getApplicableRecognitionFeatures([]PhishingSimulationRecognitionFeature{})

	// then
	assert.Nil(t, features)
	assert.Error(t, err)
	assert.Equal(t, "Undefined RecognitionFeature used in template", err.Error())
}

func TestGetScoredCombinationsBasic(t *testing.T) {
	// given
	categoryID := uuid.New()
	featureID := uuid.New()
	template := &PhishingSimulationContentTemplate{
		Subject:           "{RecognitionFeature{Name1}}",
		ContentCategoryFk: categoryID,
		ContentCategory:   &PhishingSimulationContentCategory{ID: categoryID},
	}
	value := PhishingSimulationRecognitionFeatureValue{
		Level: 2,
	}
	feature := PhishingSimulationRecognitionFeature{
		ID:                       featureID,
		Name:                     "Name1",
		RecognitionFeatureValues: &[]PhishingSimulationRecognitionFeatureValue{value},
	}
	vuln := PhishingSimulationUserVulnerability{
		ContentCategory:    &PhishingSimulationContentCategory{ID: categoryID},
		RecognitionFeature: &feature,
		Score:              2.5,
	}

	// when
	result := template.GetScoredCombinations([]PhishingSimulationUserVulnerability{vuln}, []PhishingSimulationRecognitionFeature{feature})

	// then
	assert.NotEmpty(t, result)
	assert.InDelta(t, 2.5, result[0].Score, 0.01)
	assert.Equal(t, template, result[0].Template)
	assert.Equal(t, 1, len(result[0].RecognitionFeatureValues))
}

func TestGetApplicableRecognitionFeaturesOnlyAlwaysApplicable(t *testing.T) {
	// given
	template := PhishingSimulationContentTemplate{
		Content: "No feature used",
		Subject: "Still no feature",
	}
	def := PhishingSimulationRecognitionFeature{Name: "Always", IsAlwaysApplicable: true}

	// when
	result, err := template.getApplicableRecognitionFeatures([]PhishingSimulationRecognitionFeature{def})

	// then
	assert.NoError(t, err)
	assert.ElementsMatch(t, []PhishingSimulationRecognitionFeature{def}, result.List())
}

func TestGetApplicableRecognitionFeaturesDuplicateUsage(t *testing.T) {
	// given
	template := PhishingSimulationContentTemplate{
		Content: "{RecognitionFeature{Feat}} and again {RecognitionFeature{Feat}}",
	}
	def := PhishingSimulationRecognitionFeature{Name: "Feat", IsAlwaysApplicable: false}

	// when
	result, err := template.getApplicableRecognitionFeatures([]PhishingSimulationRecognitionFeature{def})

	// then
	assert.NoError(t, err)
	assert.Len(t, result.List(), 1)
}

func TestGetScoredCombinationsNoMatchingVulnerability(t *testing.T) {
	// given
	categoryID := uuid.New()
	featureID := uuid.New()
	template := &PhishingSimulationContentTemplate{
		Subject:           "{RecognitionFeature{Feat}}",
		ContentCategoryFk: categoryID,
		ContentCategory:   &PhishingSimulationContentCategory{ID: categoryID},
	}
	featureValue := PhishingSimulationRecognitionFeatureValue{Level: 1}
	feature := PhishingSimulationRecognitionFeature{
		ID:                       featureID,
		Name:                     "Feat",
		RecognitionFeatureValues: &[]PhishingSimulationRecognitionFeatureValue{featureValue},
	}
	vuln := PhishingSimulationUserVulnerability{
		RecognitionFeature: &PhishingSimulationRecognitionFeature{
			ID:   uuid.New(),
			Name: "Unrelated",
		},
		ContentCategory: &PhishingSimulationContentCategory{ID: categoryID},
		Score:           4,
	}

	// when
	result := template.GetScoredCombinations([]PhishingSimulationUserVulnerability{vuln}, []PhishingSimulationRecognitionFeature{feature})

	// then
	assert.NotEmpty(t, result)
	assert.Equal(t, float32(1), result[0].Score)
}

func TestGetScoredCombinationsArithmeticMeanFallback(t *testing.T) {
	// given
	categoryID1 := uuid.New()
	categoryID2 := uuid.New()
	featureID := uuid.New()
	template := &PhishingSimulationContentTemplate{
		Subject:           "{RecognitionFeature{Feat}}",
		ContentCategoryFk: categoryID1,
		ContentCategory:   &PhishingSimulationContentCategory{ID: categoryID1},
	}
	featureValue := PhishingSimulationRecognitionFeatureValue{Level: 2}
	feature := PhishingSimulationRecognitionFeature{
		ID:                       featureID,
		Name:                     "Feat",
		RecognitionFeatureValues: &[]PhishingSimulationRecognitionFeatureValue{featureValue},
	}
	vuln1 := PhishingSimulationUserVulnerability{
		RecognitionFeature: &feature,
		ContentCategory:    &PhishingSimulationContentCategory{ID: categoryID2}, // different category
		Score:              2,
	}
	vuln2 := PhishingSimulationUserVulnerability{
		RecognitionFeature: &feature,
		ContentCategory:    &PhishingSimulationContentCategory{ID: categoryID2},
		Score:              4,
	}

	// when
	result := template.GetScoredCombinations([]PhishingSimulationUserVulnerability{vuln1, vuln2}, []PhishingSimulationRecognitionFeature{feature})

	// then
	assert.NotEmpty(t, result)
	assert.InDelta(t, 3.0, result[0].Score, 0.01) // mean of 2 and 4
}

func TestGetScoredCombinationsLevelCorrectedToMin1(t *testing.T) {
	// given
	categoryID := uuid.New()
	featureID := uuid.New()
	template := &PhishingSimulationContentTemplate{
		Subject:           "{RecognitionFeature{Feat}}",
		ContentCategoryFk: categoryID,
		ContentCategory:   &PhishingSimulationContentCategory{ID: categoryID},
	}
	featureValue1 := PhishingSimulationRecognitionFeatureValue{Level: 0}
	featureValue2 := PhishingSimulationRecognitionFeatureValue{Level: 1}
	feature := PhishingSimulationRecognitionFeature{
		ID:                       featureID,
		Name:                     "Feat",
		RecognitionFeatureValues: &[]PhishingSimulationRecognitionFeatureValue{featureValue1, featureValue2},
	}
	vuln := PhishingSimulationUserVulnerability{
		RecognitionFeature: &feature,
		ContentCategory:    &PhishingSimulationContentCategory{ID: categoryID},
		Score:              0, // level = 0, but should be corrected
	}

	// when
	result := template.GetScoredCombinations([]PhishingSimulationUserVulnerability{vuln}, []PhishingSimulationRecognitionFeature{feature})

	// then
	assert.NotEmpty(t, result)
	assert.GreaterOrEqual(t, (result[0].RecognitionFeatureValues)[0].Level, 1)
}
