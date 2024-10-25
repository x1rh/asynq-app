package server

import (
	"time"

	"github.com/hibiken/asynq"
)

func RetryAfterDuration(d time.Duration) asynq.RetryDelayFunc {
	return func(n int, e error, t *asynq.Task) time.Duration {
		return d
	}
}
