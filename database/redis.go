package database

import (
	//"github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
)

var RedisDB redis.Conn

func ConnectRedis() (err error) {
	RedisDB, err = redis.Dial("tcp", localhost + ":6379")
	if err != nil {
		return
	}
	return nil
}

// CloseRedis 关闭redis连接
func CloseRedis() {
	RedisDB.Close()
}


