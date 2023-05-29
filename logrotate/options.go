package logrotate

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/shinesyang/common/logger/control"
)

type OptionsFunc func(*RotateLog)

var DefaultFilePath = GetPath()

const (
	RotateIsLog     = -1 // 表示切割日志
	RotateNotLog    = 0  // 表示不切割日志
	DefaultFileName = "logs.log"
	DefaultMaxAge   = -1                   // 默认不切割日志
	MaxAgeWeek      = 7 * time.Hour * 24   // 七天
	MaxAgeMonth     = 30 * time.Hour * 24  // 一个月
	MaxAgeQuarter   = 120 * time.Hour * 24 // 一个季度
)

// 生成完整的日志文件路径
func WithRotateFilePath(pathString string, fileName string) OptionsFunc {
	return func(r *RotateLog) {
		r.FilePath = path.Join(pathString, fileName)
	}
}

// 过期删除文件
func WithDeleteExpiredFile(maxAge time.Duration) OptionsFunc {
	return func(r *RotateLog) {
		r.maxAge = maxAge
		r.deleteFileWildcard = r.FilePath + ".*"
	}
}

/*
	生成当前的日志文件不带后缀,
	当前日志格式: xxxx.log
	隔天的日志格式: xxxx.log.20230529
*/
func WithSimpleControl() OptionsFunc {
	return func(r *RotateLog) {
		r.Control = &control.SimpleControl{}
	}
}

/*
	按时间日志生成日志文件,
	当前日志格式: xxxx.log.20230530
	隔天的日志格式: xxxx.log.20230529
*/
func WithDateControl() OptionsFunc {
	return func(r *RotateLog) {
		r.Control = &control.DateControl{}
	}
}

// 调用不切割日志
func WithRotateNotLog() OptionsFunc {
	return func(r *RotateLog) {
		r.rotateLog = RotateNotLog
	}
}

// 获取目录,设置为默认日志文件目录
func GetPath() string {
	//var abPath string
	//_, filename, _, ok := runtime.Caller(0)
	//if ok {
	//	fmt.Println(filename)
	//	abPath = path.Dir(filename)
	//}

	//abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	//executable, _ := os.Executable()
	//s, _ := filepath.Abs(filepath.Dir(executable))

	dir, _ := os.Getwd()
	abPath, _ := filepath.Abs(dir)
	logPath := path.Join(abPath, "logs")

	return logPath
}
