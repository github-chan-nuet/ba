package services

import (
	"sync"
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

func ExecuteEachDayAfterDuration(d time.Duration, fn func()) {
	for {
		nextInvocation := time.Now().Truncate(time.Hour * 24).Add(time.Hour*24 + d)
		timeTillNext := nextInvocation.Sub(time.Now())
		var wg sync.WaitGroup
		wg.Add(1)
		time.AfterFunc(timeTillNext, func() {
			fn()
			wg.Done()
		})
		wg.Wait()
	}
}
