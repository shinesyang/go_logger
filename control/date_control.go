package control

import "time"

type DateControl struct {
}

func createFilePath(filePath string, now time.Time) string {
	nowTime := now
	nowTimeDayString := nowTime.Format("20060102")

	latestLogPath := filePath + "." + nowTimeDayString
	return latestLogPath
}

// 日志文件加当前时间
func (d *DateControl) RotateFilePath(filePath string, afterTime time.Time) string {
	return createFilePath(filePath, afterTime)
}

// 日志文件加当前时间
func (d *DateControl) WithFilePath(filePath string, now time.Time) string {
	return createFilePath(filePath, now)
}
