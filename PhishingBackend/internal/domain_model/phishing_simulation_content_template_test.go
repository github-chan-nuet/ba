package domain_model

import (
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
	assert.Equal(t, wantSet, c)
}

func TestExtractRecognitionFeaturesUsed(t *testing.T) {
	// given
	input := "{RecognitionFeature{Face}} {RecognitionFeature{Voice}} {RecognitionFeature{A}} {RecognitionFeature{Aa}}"

	// when
	s := extractRecognitionFeaturesUsed(input)

	// then
	assert.Equal(t, s, s)
}
