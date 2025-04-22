package services

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	// Given
	duration := 10 * time.Millisecond
	var wasCalled bool
	fn := func(utc time.Time) {
		wasCalled = true
	}

	// When
	go StartCronJob(duration, fn)

	// Then
	time.Sleep(duration * 2)
	if !wasCalled {
		t.Errorf("Expected function to be called, but it wasn't")
	}
}
