package services

import (
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

func ExecuteFunctionIn(d time.Duration, fn func(utc time.Time)) {
	ticker := time.NewTicker(d)
	now := <-ticker.C
	ticker.Stop()
	go fn(now.UTC())
}
