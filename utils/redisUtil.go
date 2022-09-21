package utils

import (
	"log"

	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	RedisToken = "Token_"
)

func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password == "" {
				return c, err
			}
			_, err = c.Do("AUTH", password)
			if err != nil {
				err := c.Close()
				if err != nil {
					log.Println("[NewPool] Poll Conn Close Failed")
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	Pool *redis.Pool
)

// SetRedisStringValue : Set redis value
func SetRedisStringValue(key string, value string, expire int) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[SetRedisStringValue] Poll Conn Close Failed")
		}
	}(conn)
	if expire < 0 {
		if _, err := conn.Do("SET", key, value); err != nil {
			log.Printf("[[SetRedisStringValue]] set data to redis failed. key: %s, err: %s\n", key, err)
		}
		return
	}

	if _, err := conn.Do("SETEX", key, expire, value); err != nil {
		log.Printf("[[SetRedisStringValue]] set expireTime to redis failed. key: %s, err: %s\n", key, err)
	}
}

// SetRedisIntValue : Set redis value
func SetRedisIntValue(key string, value int64, expire int) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[SetRedisIntValue] Poll Conn Close Failed")
		}
	}(conn)
	if expire <= 0 {
		if _, err := conn.Do("SET", key, value); err != nil {
			log.Printf("[[SetRedisIntValue]] set data to redis failed. key: %s, err: %s\n", key, err)
		}
		return
	}
	if _, err := conn.Do("SETEX", key, expire, value); err != nil {
		log.Printf("[[SetRedisIntValue]] set expireTime to redis failed. key: %s, err: %s\n", key, err)
	}
}

// GetRedisBytes : Get redis data and return byte array
func GetRedisBytes(key string) ([]byte, error) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[GetRedisBytes] Poll Conn Close Failed")
		}
	}(conn)
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Printf("[GetRedisBytes] key: %s, redis get failed: %s", key, err)
	}
	return data, err
}

// GetRedisString : Get redis data and return string
func GetRedisString(key string) (strData string, err error) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[GetRedisString] Poll Conn Close Failed")
		}
	}(conn)
	strData, err = redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("[[GetRedisString]] get stringInfo failed, key: %s, err: %s\n", key, err)
	}
	return
}

// GetRedisInt : Get redis data and return int64
func GetRedisInt(key string) (int64, error) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[GetRedisInt] Poll Conn Close Failed")
		}
	}(conn)
	return redis.Int64(conn.Do("GET", key))
}

// GetRedisTTL : Get redis key TTL
func GetRedisTTL(key string) (int, error) {
	conn := Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("[GetRedisTTL] Poll Conn Close Failed")
		}
	}(conn)
	return redis.Int(conn.Do("TTL", key))
}

func SetToken(username, token string) {
	username = strings.ToLower(username)
	key := RedisToken + username
	SetRedisStringValue(key, token, 60*60*24) // 24hr
}

func VerifyToken(username, token string) bool {
	username = strings.ToLower(username)
	key := RedisToken + username
	value, err := GetRedisString(key)
	if err != nil {
		log.Println("Get redis value error : ", err)
		return false
	}
	return token == value
}
