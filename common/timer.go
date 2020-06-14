package common

import (
	"errors"
	"github.com/mittacy/blog-gin/cache"
	"github.com/mittacy/blog-gin/log"
	"time"
)

// startTimer 定时清理缓存任务任务
func StartTimer() {
	go func() {
		for {
			// 1. 更新admin新增访问量到数据库
			if err := cache.UpdateViewsToDatabase(); err != nil {
				log.RecordErr(errors.New("更新博客新增访问量到数据库失败, err: " + err.Error()))
			}
			// 2. 更新文章访问量到数据库
			if err := cache.UpdateArticleViewsToDatabase(); err != nil {
				log.RecordErr(errors.New("更新文章新增访问量到数据库失败, err: " + err.Error()))
				log.RecordErr(err)
			}
			// 3. 更新内存缓冲器缓存
			if err := cache.InitCache(); err != nil {
				log.ErrLogger.Fatalln(err)
			}
			// 4. 计算下一次更新时间
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
