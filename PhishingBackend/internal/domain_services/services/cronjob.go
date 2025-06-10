package services

import (
	"log/slog"
	"math/rand"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCronJob(d time.Duration, fn func(utc time.Time)) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		now := <-ticker.C
		fn(now.UTC())
	}
}

func StartCronStyleJob(cronExpr string, fn func(utc time.Time)) (*cron.Cron, cron.EntryID, error) {
	c := cron.New()

	entryID, err := c.AddFunc(cronExpr, func() {
		now := time.Now().UTC()
		fn(now)
	})
	if err != nil {
		return nil, 0, err
	}

	c.Start()
	return c, entryID, nil
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
