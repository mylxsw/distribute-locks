package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mylxsw/glacier/scheduler"
	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/must"
)

func TestRandomToken(t *testing.T) {
	for i := 0; i < 100; i++ {
		token1 := must.Must(randomToken())
		token2 := must.Must(randomToken())
		assert.True(t, token1 != token2)
	}
}

func TestTryLock(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	locker := New(rdb, "test_lock", 10*time.Second)
	assert.NoError(t, locker.TryLock(context.TODO()))
	assert.NoError(t, locker.TryLock(context.TODO()))

	assert.NoError(t, locker.Release(context.TODO()))

	assert.NoError(t, locker.TryLock(context.TODO()))
	assert.NoError(t, locker.TryLock(context.TODO()))

	assert.NoError(t, locker.Release(context.TODO()))
	assert.NoError(t, locker.Release(context.TODO()))

	locker2 := New(rdb, "test_lock", 10*time.Second)
	assert.NoError(t, locker2.TryLock(context.TODO()))
	assert.Equal(t, scheduler.ErrLockFailed, locker.TryLock(context.TODO()))
	assert.NoError(t, locker2.TryLock(context.TODO()))

	assert.NoError(t, locker.Release(context.TODO()))
	assert.Equal(t, scheduler.ErrLockFailed, locker.TryLock(context.TODO()))

	assert.NoError(t, locker2.Release(context.TODO()))
	assert.NoError(t, locker.TryLock(context.TODO()))
	assert.Equal(t, scheduler.ErrLockFailed, locker2.TryLock(context.TODO()))

	assert.NoError(t, locker.Release(context.TODO()))
}
