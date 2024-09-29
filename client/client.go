package client

import (
	"crypto/tls"
	"encoding/json"

	"asynq-app/config"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

type Client struct {
	client *asynq.Client
}

func New(c config.Config) *Client {
	redisOpts := asynq.RedisClientOpt{
		Addr:     c.RedisConfig.Addr,
		Password: c.RedisConfig.Password,
		Username: c.RedisConfig.Username,
		DB:       0,
	}

	if c.RedisConfig.EnableTls {
		redisOpts.TLSConfig = &tls.Config{}
	}

	return &Client{
		client: asynq.NewClient(redisOpts),
	}
}

func (c *Client) Enqueue(taskType string, task any, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(task)
	if err != nil {
		return nil, errors.Wrap(err, "marshal error")
	}

	taskInfo, err := c.client.Enqueue(asynq.NewTask(taskType, payload), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "fail to enqueue task")
	}
	return taskInfo, nil
}
