package biz1

import (
	"asynq-app/appctx"
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"github.com/hibiken/asynq"
)

type Task1 struct {
	Id  int    `json:"id"`
	App string `json:"app"`
}

func HandleTask1(ctx *appctx.Context) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		if time.Now().UTC().Unix()%2 == 0 && rand.Int()%2 == 0 {
			panic("random panic task2")
		}

		var t Task1
		err := json.Unmarshal(task.Payload(), &t)
		if err != nil {
			slog.Error("fail to unmarshal task", slog.Any("err", err))
			return err
		}

		slog.Info("handle task1", slog.Any("task1", t))
		return nil
	}
}
