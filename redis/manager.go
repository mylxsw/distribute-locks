package redis

import (
	"github.com/mylxsw/glacier/scheduler"
	"gopkg.in/redis.v5"
)

type LockManager struct {
	client *redis.Client
	name   string
}

func New(client *redis.Client, name string) scheduler.LockManager {
	return &LockManager{client: client, name: name}
}

func (m *LockManager) TryLock() error {
	return nil
}

func (m *LockManager) Release() error {
	return nil
}
