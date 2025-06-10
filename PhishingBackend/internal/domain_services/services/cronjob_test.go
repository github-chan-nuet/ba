package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestRandomStart(t *testing.T) {
	// Given
	minDuration := 100 * time.Millisecond
	maxDuration := 2 * time.Second

	var timesCalled int
	calledAt := []time.Time{}
	fn := func(utc time.Time) {
		calledAt = append(calledAt, utc)
		timesCalled += 1
	}

	// When
	go StartRandomCronJob(minDuration, maxDuration, fn)

	// Then
	time.Sleep(10 * time.Second)
	assert.LessOrEqual(t, timesCalled, 100)
	assert.GreaterOrEqual(t, timesCalled, 5)

	for i := 1; i < len(calledAt); i++ {
		diff := calledAt[i].Sub(calledAt[i-1])
		assert.LessOrEqual(t, diff, maxDuration)
		assert.GreaterOrEqual(t, diff, minDuration)
	}
}
