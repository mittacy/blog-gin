package cache

import (
	"github.com/mittacy/blog-gin/config"
	"github.com/mittacy/blog-gin/database"
	"github.com/gomodule/redigo/redis"
)
// GetIPTimes 获取ip访问次数
func GetIPTimes(ip string) (int, error) {
	exists, err := redis.Bool(database.RedisDB.Do("EXISTS", ip))
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, nil
	}
	return redis.Int(database.RedisDB.Do("GET", ip))
}
// IPIncr ip访问次数+1
func IPIncr(ip string) error {
	exists, err := redis.Bool(database.RedisDB.Do("EXISTS", ip))
	if err != nil {
		return err
	}
	if exists {
		err = database.RedisDB.Send("INCR", ip)
		return err
	}
	_, err = database.RedisDB.Do("SET", ip, 1, "EX", config.ProhibitIPTime)
	return err
}
// IPDel 删除IP访问次数记录
func IPDel(ip string) (err error) {
	_, err = database.RedisDB.Do("DEL", ip)
	return
}
