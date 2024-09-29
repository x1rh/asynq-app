package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"asynq-app/appctx"
	"asynq-app/task"
	"asynq-app/task/biz1"
	"asynq-app/task/biz2"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
)

type Server struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	appctx *appctx.Context
}

func New(ctx *appctx.Context) (*Server, error) {
	c := ctx.Config
	redisOpts, err := redis.ParseURL(
		fmt.Sprintf("rediss://%s:%s@%s",
			c.RedisConfig.Username, c.RedisConfig.Password, c.RedisConfig.Addr,
		))
	if err != nil {
		return nil, errors.New("fail to parse redis config")
	}

	asynqRedisOpts := asynq.RedisClientOpt{
		Addr:     redisOpts.Addr,
		Password: redisOpts.Password,
		DB:       0,
	}

	if ctx.Config.RedisConfig.EnableTls {
		asynqRedisOpts.TLSConfig = &tls.Config{}
	}

	server := asynq.NewServer(asynqRedisOpts, asynq.Config{
		Concurrency: 8,
		Queues:      nil,
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return time.Second
		},
	})

	return &Server{
		server: server,
		mux:    asynq.NewServeMux(),
	}, nil
}

func (svr *Server) Start() {
	svr.HandleFunc(task.TypeTask1, biz1.HandleTask1(svr.appctx))
	svr.HandleFunc(task.TypeTask2, biz2.HandleTask2(svr.appctx))

	slog.Info("asynq server start...")
	if err := svr.server.Run(svr.mux); err != nil {
		slog.Error("asynq server run failed!!!", slog.Any("error", err))
	}
}

func (svr *Server) HandleFunc(pattern string, taskHandler func(ctx context.Context, task *asynq.Task) error) {
	svr.mux.HandleFunc(pattern, taskHandler)
}

func (svr *Server) Stop() {
	slog.Info("asynq server stop...")
	svr.server.Stop()
}
