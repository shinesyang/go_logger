#### 根据日期实现自动日志切割功能
`使用time.Ticker来触发实现根据日期实现日志自动按照当前日期进行切割,本项目在其他大佬的项目基础上进行改编而来,地址: https://github.com/Me1onRind/logrotate`

##### Example
```go
	writer, err := logrotate.NewRotateLog(
		logrotate.WithRotateFilePath(logrotate.DefaultFilePath, logrotate.DefaultFileName),
		logrotate.WithDeleteExpiredFile(logrotate.MaxAgeQuarter),
		logrotate.WithSimpleControl(),
	)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer writer.Close()
    if _, err := writer.Write([]byte("Hello,World!\n")); err != nil {
    fmt.Println(err)
    return
    }
	
	// 结合开源项目zap日志工具使用:
	zapcore.AddSync(r)

```

`在options中定义了一些基础的调用参数,你可以使用它,或者重新定义`

`zap地址: https://github.com/uber-go/zap`