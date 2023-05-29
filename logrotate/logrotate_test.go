package logrotate

import (
	"testing"
	"time"
)

func test(t *testing.T) {
	log, _ := NewRotateLog(
		WithRotateFilePath(DefaultFilePath, DefaultFileName),
		WithDeleteExpiredFile(MaxAgeQuarter),
		WithDateControl(),
	)

	for {
		<-time.After(time.Second)
		for i := 0; i <= 10; i++ {
			t := time.Now().Format("2006-01-02 15:04:05")
			log.Write([]byte(t + "s ka\n"))
		}

	}
}
