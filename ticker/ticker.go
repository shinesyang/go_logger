package ticker

import "time"

func CalRotateTimeDuration(now time.Time) time.Duration {

	// 获取本地明天0点的时间戳(单位毫秒)
	nowTimeDayString := now.Format("2006-01-02")
	nowTimeDay, _ := time.ParseInLocation("2006-01-02", nowTimeDayString, time.Local)
	tomorrowTimeDay := nowTimeDay.Add(time.Hour * 24)
	tomorrowTimeDayMilli := tomorrowTimeDay.UnixMilli()
	// 当前时间戳(单位毫秒)
	nowTimeMilli := now.UnixMilli()

	// 距离明天0点的时间戳剩余时间
	diifSecondMilli := tomorrowTimeDayMilli - nowTimeMilli
	countTime := time.Duration(diifSecondMilli) * time.Millisecond
	return countTime
}
