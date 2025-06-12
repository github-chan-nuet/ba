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
	input := "Lorem Ipsum {RecognitionFeature{Voice}} {RecognitionFeature{A}} {RecognitionFeature{Aa}}"

	// when
	s := extractRecognitionFeaturesUsed(input)

	// then
	assert.Equal(t, []string{"Voice", "A", "Aa"}, s.List())
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
