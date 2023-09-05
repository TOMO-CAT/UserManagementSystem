package db

import (
	"fmt"
	"time"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/gomodule/redigo/redis"
)

var (
	PtrRedisClient *redisClient
)

type redisClient struct {
	pool *redis.Pool
}

func InitRedisClient() {
	PtrRedisClient = &redisClient{
		pool: &redis.Pool{
			MaxIdle:     config.GlobalUmsConfig.Redis.MaxIdle,
			MaxActive:   config.GlobalUmsConfig.Redis.MaxActive,
			IdleTimeout: time.Second * time.Duration(config.GlobalUmsConfig.Redis.IdleTimeoutSeconds),
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.GlobalUmsConfig.Redis.IP, config.GlobalUmsConfig.Redis.Port))
				if err != nil {
					logger.Error("redis dial fail with err [%v]", err)
					return nil, err
				}
				// if _, err = c.Do("AUTH", "password"); err != nil {
				// 	c.Close()
				// 	logger.Error("redis auth fail with err [%v]", err)
				// 	return nil, err
				// }
				// if _, err = c.Do("SELECT", "index"); err != nil {
				// 	c.Close()
				// 	logger.Error("redis select fail with err [%v]", err)
				// 	return nil, err
				// }
				return c, nil
			},
		},
	}
}

// Redis Command Warpper: https://redis.io/commands

// redis:6379> MSET firstname Jack lastname Stuntman age 35
// "OK"
// redis:6379> KEYS *name*
// 1) "lastname"
// 2) "firstname"
// redis:6379> KEYS a??
// 1) "age"
// redis:6379> KEYS *
// 1) "age"
// 2) "lastname"
// 3) "firstname"
func (c *redisClient) KEYS(pattern string) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

// GET nonexisting
// (nil)
// SET mykey "Hello"
// "OK"
// GET mykey
// "Hello"
func (c *redisClient) GET(key string) (value string, err error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// SETNX mykey "Hello"
// (integer) 1
// SETNX mykey "World"
// (integer) 0
// 由于 SETNX 不支持设置过期时间, 改成 SET key value EX expireTimeSecond NX
func (c *redisClient) SETNX(key string, value interface{}, expireTimeSecond int) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	if expireTimeSecond <= 0 {
		return redis.String(conn.Do("SET", key, value, "NX"))
	} else {
		return redis.String(conn.Do("SET", key, value, "EX", expireTimeSecond, "NX"))
	}
}

// DEL key
// (integer) 2
func (c *redisClient) DEL(key string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

// redis:6379> HSET myhash field1 "Hello"
// (integer) 1
// redis:6379> HGET myhash field1
// "Hello"
func (c *redisClient) HSET(key string, field string, value interface{}) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HSET", key, field, value))
}

// HMSET myhash field1 "Hello" field2 "World"
// "OK"
func (c *redisClient) HMSET(key string, fieldValues []interface{}) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	var args []interface{}
	args = append(args, key)
	args = append(args, fieldValues...)
	return redis.String(conn.Do("HMSET", args...))
}

// HGETALL myhash
// 1) "field1"
// 2) "Hello"
// 3) "field2"
// 4) "World"
func (c *redisClient) HGETALL(key string) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HGETALL", key))
}

// HVALS myhash
// 1) "Hello"
// 2) "World"
func (c *redisClient) HVALS(key string) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HVALS", key))
}

// redis:6379> SET mykey "Hello"
// "OK"
// redis:6379> EXPIRE mykey 10
// (integer) 1
// redis:6379> TTL mykey
// (integer) 10
func (c *redisClient) EXPIRE(key string, expireTimeSecond int) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXPIRE", key, expireTimeSecond))
}

// redis:6379> SADD myset "Hello"
// (integer) 1
// redis:6379> SADD myset "World"
// (integer) 1
// redis:6379> SMEMBERS myset
// 1) "World"
// 2) "Hello"
func (c *redisClient) SMEMBERS(key string) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("SMEMBERS", key))
}

// redis:6379> SADD myset "Hello"
// (integer) 1
// redis:6379> SADD myset "World"
// (integer) 1
// redis:6379> SADD myset "World"
// (integer) 0
func (c *redisClient) SADD(key string, value interface{}) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("SADD", key, value))
}
