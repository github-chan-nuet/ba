package model

import "time"

type JobLog struct {
	ExecutionTime time.Time `db:"EXECUTION_TIME" fieldtag:"pk"`
	ExecutorId    string    `db:"EXECUTOR_ID"`
}

func (j *JobLog) IsOlder(toCompare time.Time, limit time.Duration) bool {
	buffer := time.Millisecond * 1
	return j.ExecutionTime.IsZero() || toCompare.Add(buffer).Sub(j.ExecutionTime) > limit
}
