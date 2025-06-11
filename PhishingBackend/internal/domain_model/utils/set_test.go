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

	// when
}
