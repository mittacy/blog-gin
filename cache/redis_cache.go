package cache

import (
	"github.com/crazychat/blog-gin/database"
	"time"
)

// RedisIncr redis对象加1
func RedisIncr(key string) error {
	return database.RedisDB.Incr(key).Err()
}
// RedisDel 删除redis对象
func RedisDel(key string) error {
	return  database.RedisDB.Del(key).Err()
}
// RedisIncr redis获取对象值
func RedisGet(key string) (string, error) {
	return  database.RedisDB.Get(key).Result()
}
// RedisSet redis保存对象
func RedisSet(key, val string, duration time.Duration) error {
	return  database.RedisDB.Set(key, val, duration).Err()
}
