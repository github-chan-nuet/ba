package services

import (
	"log/slog"
	"math/rand"
	"time"
)

func StartCronJob(d time.Duration, fn func(utc time.Time)) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		now := <-ticker.C
		fn(now.UTC())
	}
}

func StartRandomCronJob(min, max time.Duration, fn func(utc time.Time)) {
	if max <= min {
		panic("max must be greater than min")
	}
	for {
		sleepDuration := min + time.Duration(rand.Int63n(int64(max-min)))
		slog.Info("Sleeping for", "duration", sleepDuration.String())
		time.Sleep(sleepDuration)

		fn(time.Now().UTC())
	}
}

func ExecuteFunctionIn(d time.Duration, fn func(utc time.Time)) {
	ticker := time.NewTicker(d)
	now := <-ticker.C
	ticker.Stop()
	go fn(now.UTC())
}
