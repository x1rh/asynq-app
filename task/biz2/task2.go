package biz2

import (
	"asynq-boilerplate/asynqctx"
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"github.com/hibiken/asynq"
)

type Task2 struct {
	Id  int    `json:"id"`
	App string `json:"app"`
}

func HandleTask2(ctx *asynqctx.Context) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		var t Task2

		if time.Now().UTC().Unix()%2 == 0 && rand.Int()%2 == 0 {
			panic("random panic task2")
		}

		err := json.Unmarshal(task.Payload(), &t)
		if err != nil {
			slog.Error("fail to unmarshal task", slog.Any("err", err))
			return err
		}

		slog.Info("handle task2", slog.Any("task2", t))
		return nil
	}
}
