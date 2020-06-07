package common

import "time"

// startTimer 定时清理缓存任务任务
func StartTimer() {
	go func() {
		for {
			//SaveBlogViews()
			//SaveAdminInfo()
			//SaveRecentArticles()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
