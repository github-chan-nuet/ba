package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAdd(t *testing.T) {
	// given
	s := NewSet[int]()

	// when
	s.Add(1)

	// then
	_, ok := s.elements[1]
	assert.True(t, ok)
}

func TestSetAddSameValue(t *testing.T) {
	// given
	s := NewSet[int]()
	s.Add(1)
	s.Add(1)
	s.Add(1)

	// when
	s.Add(1)

	// then
	_, ok := s.elements[1]
	assert.True(t, ok)
	assert.Equal(t, 1, len(s.elements))
}

func TestSetUnion(t *testing.T) {
	// given
	s1 := NewSet[int]()
	s2 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s2.Add(2)
	s2.Add(3)

	// when
	result := s1.Union(s2)

	// then
	assert.ElementsMatch(t, []int{1, 2, 3}, getKeys(result.elements))
}

func TestSetList(t *testing.T) {
	// given
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// when
	result := s.List()

	// then
	assert.ElementsMatch(t, []int{1, 2, 3}, result)
}

func TestSetSize(t *testing.T) {
	// given
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// when
	result := s.Size()

	// then
	assert.Equal(t, 3, result)
}

func TestSetContains(t *testing.T) {
	// given
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// when
	result := s.Contains(1)

	// then
	assert.Equal(t, true, result)
}

func TestSetDoesNotContain(t *testing.T) {
	// given
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// when
	result := s.Contains(9)

	// then
	assert.ElementsMatch(t, false, result)
}

func TestSetDifferenceWithOverlaps(t *testing.T) {
	// given
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)

	// when
	result := s1.Difference(s2)

	// then
	assert.ElementsMatch(t, []int{1}, result.List())
}

func TestSetDifferenceWithoutOverlaps(t *testing.T) {
	// given
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)

	s2 := NewSet[int]()
	s2.Add(3)
	s2.Add(4)

	// when
	result := s1.Difference(s2)

	// then
	assert.ElementsMatch(t, []int{1, 2}, result.List())
}

func getKeys[T comparable](m map[T]struct{}) []T {
	result := make([]T, len(m))
	i := 0
	for k := range m {
		result[i] = k
		i++
	}
	return result
}
