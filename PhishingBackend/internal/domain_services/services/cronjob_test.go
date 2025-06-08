package services

import (
	"github.com/stretchr/testify/assert"
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
	assert.True(t, wasCalled, "Expected function to be called, but it wasn't")
}

func TestExecuteEachDayAfterDuration(t *testing.T) {
	// Given
	now := time.Now()
	// nano from midnight till now
	durTilNow := now.Sub(now.Truncate(24 * time.Hour))
	duration := durTilNow - 24*time.Hour + 1*time.Second
	var wasCalled bool
	fn := func() {
		wasCalled = true
	}

	// When
	go ExecuteEachDayAfterDuration(duration, fn)

	// Then
	time.Sleep(time.Second * 2)
	assert.True(t, wasCalled, "Expected function to be called, but it wasn't")
}
