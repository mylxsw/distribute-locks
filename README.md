# Glacier 开发框架任务调度器分布式锁实现

目前支持 redis 实现。

```go

import (
	"time"

	"github.com/go-redis/redis/v8"
	redisLock "github.com/mylxsw/distribute-locks/redis"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/glacier/scheduler"
)

type Provider struct{}

func (Provider) Aggregates() []infra.Provider {
	return []infra.Provider{
		scheduler.Provider(
			func(cc infra.Resolver, creator scheduler.JobCreator) {
				// add your jobs here
			},
			scheduler.SetLockManagerOption(func(resolver infra.Resolver) scheduler.LockManagerBuilder {
                // get redis instance
				redisClient := resolver.MustGet(&redis.Client{}).(*redis.Client)
				return func(name string) scheduler.LockManager {
					// create redis lock
                    return redisLock.New(redisClient, name, 10*time.Minute)
				}
			}),
		),
	}
}

func (Provider) Register(cc infra.Binder) {}
```