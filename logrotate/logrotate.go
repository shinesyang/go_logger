package logrotate

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/shinesyang/go_logger/ticker"
)

type RotateLog struct {
	Control            Control
	file               *os.File
	FilePath           string
	rotate             <-chan time.Time
	Mutex              *sync.Mutex
	rotateLog          int
	maxAge             time.Duration
	deleteFileWildcard string
}

type Control interface {
	RotateFilePath(string, time.Time) string
	WithFilePath(string, time.Time) string
}

func NewRotateLog(options ...OptionsFunc) (*RotateLog, error) {
	r := &RotateLog{
		Mutex:     &sync.Mutex{},
		maxAge:    DefaultMaxAge,
		rotateLog: RotateIsLog,
	}
	for _, opt := range options {
		opt(r)
	}

	// 创建文件目录
	if err := os.Mkdir(filepath.Dir(r.FilePath), 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	nowTime := time.Now()

	latestLogPath := r.Control.WithFilePath(r.FilePath, nowTime)
	if err := r.rotateFile(nowTime, latestLogPath); err != nil {
		return nil, err
	}

	go r.handleEvent()

	return r, nil
}

func (r *RotateLog) rotateFile(now time.Time, latestLogPath string) error {
	// 判断是否需要切割
	if r.rotateLog == RotateIsLog {
		r.rotate = time.After(ticker.CalRotateTimeDuration(now))
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	file, err := os.OpenFile(latestLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if r.file != nil {
		r.file.Close()
	}
	r.file = file

	//if len(r.FilePath) > 0 {
	//	os.Remove(r.FilePath)
	//	os.Link(latestLogPath, r.FilePath)
	//}

	if r.maxAge > 0 && r.deleteFileWildcard != "" { // at present
		go r.deleteExpiredFile(now)
	}

	return nil
}

func (r *RotateLog) handleEvent() {
	for {
		eventTime := <-r.rotate
		r.Close()
		latestLogPath := r.Control.RotateFilePath(r.FilePath, eventTime)
		r.rotateFile(eventTime, latestLogPath)

	}

}

// 删除文件
func (r *RotateLog) deleteExpiredFile(now time.Time) {
	cutoffTime := now.Add(-r.maxAge)                    // maxAge之前时间
	matches, err := filepath.Glob(r.deleteFileWildcard) // 获取当前所有文件
	if err != nil {
		return
	}

	toUnlink := make([]string, 0, len(matches))
	for _, path := range matches {
		fileInfo, err := os.Stat(path)
		if err != nil {
			continue
		}

		// 判断文件是否为cutoffTime时间之后修改的文件
		if fileInfo.ModTime().After(cutoffTime) {
			continue
		}

		// 判断文件是否为不带日期的当前文件
		if fileInfo.Name() == filepath.Base(r.FilePath) {
			continue
		}
		toUnlink = append(toUnlink, path)
	}

	for _, path := range toUnlink {
		os.Remove(path)
	}
}

func (r *RotateLog) Write(p []byte) (n int, err error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.file.Write(p)
}

func (r *RotateLog) Close() error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.file.Close()
}
