package xLog

import (
	"fmt"
	"github.com/myskull/common/xdate"
	"runtime"
	"time"
)

const (
	LOG_INFO    = 2
	LOG_DEBUG   = 4
	LOG_WARNING = 8
	LOG_ERROR   = 16
	LOG_ALL     = LOG_INFO + LOG_DEBUG + LOG_WARNING + LOG_ERROR
)
const (
	textColorBlack = iota + 30
	textColorRed
	textColorGreen
	textColorYellow
	textColorBlue
	textColorPurple
	textColorCyan
	textColorWrite
)

var logLevel = LOG_ALL

func SetLevel(level int) {
	logLevel = level
}

func Info(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)

	text = fmt.Sprintf("%v:INFO:%v %v", getFileLine(), xdate.ToDate(time.Now().Unix()), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorCyan, text))
}

func Warning(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:WARNING:%v %v", getFileLine(), xdate.ToDate(time.Now().Unix()), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorYellow, text))
}

func Debug(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:DEBUG:%v %v", getFileLine(), xdate.ToDate(time.Now().Unix()), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorGreen, text))
}

func Error(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:ERROR:%v %v", getFileLine(), xdate.ToDate(time.Now().Unix()), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorRed, text))
}
func getFileLine() string {
	//fmt.Println(runtime.Caller(0))
	//fmt.Println(runtime.Caller(1))
	//fmt.Println(runtime.Caller(2))
	_, file, line, _ := runtime.Caller(2)
	//if strings.HasSuffix("") {
	//
	//}
	return fmt.Sprintf("%v:%v", file, line)
}
