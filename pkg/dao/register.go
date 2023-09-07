package dao

import (
	"context"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/db"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/gomodule/redigo/redis"
)

const (
	kLockRegisterExpireTimeSecond = 5 // Redis 注册锁过期时间
	kUserExistExpireTimeSecond    = 5 // Regis 用户存在过期时间

	kLockRegisterRedisKeyPrefix = "LOCK_REGISTER_"
	kUserExistRedisKeyPrefix    = "USER_EXIST_"
)

// LockRegister 通过 SETNX 保证同意用户同一时刻不会重复进入注册流程
// @param userName 用户名
// @param token 保证自己加锁只能自己解锁 (或者等待过期失效)
// @return isLockSuccess 是否加锁成功
func LockRegister(ctx context.Context, userName string, token string) (isLockSuccess bool) {
	var (
		redisKey = kLockRegisterRedisKeyPrefix + userName
		err      error
	)

	_, err = db.PtrRedisClient.SETNX(redisKey, token, kLockRegisterExpireTimeSecond)
	if err == redis.ErrNil {
		logger.WarnTrace(ctx, "user [%s] is already lock for register", userName)
		return false
	}

	if err != nil {
		logger.ErrorTrace(ctx, "lock user [%s] for register fail with err [%v]", userName, err)
		return false
	}

	logger.InfoTrace(ctx, "lock user [%s] for register successfully", userName)
	return true
}

// IsUserRegistered 判断用户是否已经注册了
func IsUserRegistered(ctx context.Context) (isRegister bool, err error) {
	// 先查 redis 缓存

	return
}
