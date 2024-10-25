package server

import (
	"asynq-boilerplate/asynqctx"
	"asynq-boilerplate/task"
	"context"
	"log/slog"

	"github.com/hibiken/asynq"
)

func errorHandlerWrapper(appctx *asynqctx.Context) asynq.ErrorHandlerFunc {
	return func(ctx context.Context, t *asynq.Task, err error) {
		_ = appctx // make it happy
		retried, _ := asynq.GetRetryCount(ctx)
		maxRetry, _ := asynq.GetMaxRetry(ctx)
		if retried >= maxRetry {
			slog.Error("task retry exhausted", slog.String("task type", t.Type()), slog.Any("err", err))
			if t.Type() == task.TypeTask1 {
				// do something when a task if failed and retry exhausted
			}
		}
	}
}
