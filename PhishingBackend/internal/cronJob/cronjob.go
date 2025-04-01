package cronjob

import (
	"time"
)

func Start(d time.Duration, fn func(utc time.Time)) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		now := <-ticker.C
		fn(now.UTC())
	}
}
