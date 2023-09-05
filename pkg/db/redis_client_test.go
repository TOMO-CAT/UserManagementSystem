package db

import (
	"testing"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func init4test(t *testing.T) {
	config.GlobalUmsConfig.Redis.IP = "127.0.0.1"
	config.GlobalUmsConfig.Redis.Port = 6379
	config.GlobalUmsConfig.Redis.MaxIdle = 3
	config.GlobalUmsConfig.Redis.MaxActive = 10
	config.GlobalUmsConfig.Redis.IdleTimeoutSeconds = 10

	InitRedisClient()

	if err := logger.InitLoggerDefault(); err != nil {
		t.Errorf("init logger fail with err [%v]", err)
	}
}

// TestSETNXWithoutExpireTime 不设置过期时间的 SETNX
func TestSETNXWithoutExpireTime(t *testing.T) {
	init4test(t)
	defer logger.Close()

	const (
		kRedisKey = "REDIS_KEY_SETNX"
	)

	PtrRedisClient.DEL(kRedisKey)
	res, err := PtrRedisClient.SETNX(kRedisKey, "TestSETNXValue", 0)
	assert.Nil(t, err)

	logger.Info("SETNX without expire time||res=%v||err=%v", res, err)
	res, err = PtrRedisClient.SETNX(kRedisKey, "TestSETNXValue", 0)
	assert.Equal(t, err, redis.ErrNil)
	logger.Info("SETNX with same key && without expire time||res=%v||err=%v", res, err)

	PtrRedisClient.DEL(kRedisKey)
}
