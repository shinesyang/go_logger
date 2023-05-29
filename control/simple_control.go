package control

import (
	"fmt"
	"os"
	"time"
)

// 判断文件存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type SimpleControl struct {
}

func (s *SimpleControl) RotateFilePath(filePath string, afterTime time.Time) string {
	// 检查判断filePath是否存在
	if ok, err := pathExists(filePath); err != nil || !ok {
		fmt.Printf("文件不存在或者查询出错: %v\n", err)
		return filePath
	}

	// 将当天日志文件重新命名为隔天的
	yesterdayTime := afterTime.Add(-time.Hour * 24)
	yesterdayTimeString := yesterdayTime.Format("20060102")
	latestLogPath := filePath + "." + yesterdayTimeString
	if err := os.Rename(filePath, latestLogPath); err != nil {
		fmt.Printf("文件重命名失败: %v\n", err)
	}
	return filePath
}

/*
	程序第一次启动的时候先判断有没有存在当前检索的文件(xxx.log)，
	如果则要判断一下这个日志最后修改什么时候然后把这个日志文件修改成最后修改时带后缀的文件名
	不能直接使用,不然生成的当前的日志就和老的日志文件同时存在一个文件中
*/
func (s *SimpleControl) WithFilePath(filePath string, now time.Time) string {
	// 检查判断filePath是否存在
	if ok, err := pathExists(filePath); err != nil || !ok {
		return filePath
	}

	// 获取文件最后修改时间
	finfo, _ := os.Stat(filePath)
	//modTime := finfo.ModTime().Format("2006-01-02 15:04:05")
	formatTimeString := finfo.ModTime().Format("20060102")

	// 获取当前时间
	nowTimeString := now.Format("20060102")

	// 用文件创建时间组成一个新的文件
	if formatTimeString != nowTimeString {
		fileDatePath := filePath + "." + formatTimeString
		if err := os.Rename(filePath, fileDatePath); err != nil {
			fmt.Printf("文件重命名失败: %v\n", err)
		}
	}

	return filePath
}

//func test() {
//	finfo, _ := os.Stat(filePath)
//	osType := runtime.GOOS
//	if osType == "windows" {
//		status := finfo.Sys().(*syscall.Win32FileAttributeData)
//		seconds := time.Duration(status.CreationTime.Nanoseconds()).Seconds()
//		fmt.Printf("时间戳错 ：%v\n", int(seconds))
//		//formatTimeStr := time.Unix(int64(seconds), 0).Format("2006-01-02 15:04:05")
//		formatTimeString = time.Unix(int64(seconds), 0).Format("20060102")
//	} else if osType == "linux" {
//		status := finfo.Sys().(*syscall.Stat_t)
//		seconds := int64(status.Ctim.Sec)
//		formatTimeString = time.Unix(int64(seconds), 0).Format("20060102")
//	}
//}
