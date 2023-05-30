package control

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/shinesyang/go_logger/lib"
)

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
	// 昨天的日志名
	yesterdayTime := afterTime.Add(-time.Hour * 24)
	yesterdayTimeString := yesterdayTime.Format("20060102")
	yesterdayLogPath := filePath + "." + yesterdayTimeString

	// 今天的日志名
	nowTimeDayString := afterTime.Format("20060102")
	latestLogPath := filePath + "." + nowTimeDayString

	// 判断日志是否存在
	if ok, err := lib.PathExists(yesterdayLogPath); err == nil && ok {
		finfo, _ := os.Stat(yesterdayLogPath)
		if finfo.Size() <= 0 {
			_ = os.Rename(yesterdayLogPath, latestLogPath) // 当日志文件为空文件，日志名重置为今天
		}
	}

	return latestLogPath
}

// 日志文件加当前时间
func (d *DateControl) WithFilePath(filePath string, now time.Time) string {
	nowTimeDayString := now.Format("20060102")
	latestLogPath := filePath + "." + nowTimeDayString

	matches, err := filepath.Glob(filePath + ".*") // 获取当前所有文件
	if err != nil {
		return latestLogPath
	}

	for _, ph := range matches {
		fileInfo, err := os.Stat(ph)
		if err != nil {
			continue
		}

		dir := filepath.Dir(filePath)
		oldLogPath := path.Join(dir, fileInfo.Name())

		if fileInfo.Size() <= 0 {
			_ = os.Rename(oldLogPath, latestLogPath) // 当日志文件为空文件，日志名重置为今天
			break
		}
	}

	return latestLogPath
}
