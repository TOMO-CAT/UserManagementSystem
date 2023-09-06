package dao

import (
	"github.com/TOMO-CAT/UserManagementSystem/pkg/db"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/gomodule/redigo/redis"
)

const (
	kLockRegisterExpireTimeSecond = 5

	kLockRegisterRedisKeyPrefix = "LOCK_REGISTER_"
)

// LockRegister 通过 SETNX 保证用户不会重复注册
// @param userName 用户名
// @param token 保证自己加锁只能自己解锁 (或者等待过期失效)
func LockRegister(userName string, token string) (isLock bool, err error) {
	var (
		redisKey = kLockRegisterRedisKeyPrefix + userName
	)

	_, err = db.PtrRedisClient.SETNX(redisKey, token, kLockRegisterExpireTimeSecond)
	if err == redis.ErrNil {
		logger.Warn("user [%s] is already lock for register", userName)
	}

	return
}

func IsUserRegistered() (isRegister bool, err error) {
	return
}
