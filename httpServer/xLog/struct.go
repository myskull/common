package xLog

import (
	"fmt"
	"runtime"
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
	text = fmt.Sprintf("%v:INFO %v", getFileLine(), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorCyan, text))
}

func Warning(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:WARNING %v", getFileLine(), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorYellow, text))
}

func Debug(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:DEBUG %v", getFileLine(), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorGreen, text))
}

func Error(log string, format ...interface{}) {
	if logLevel&LOG_INFO <= 0 {
		return
	}
	text := fmt.Sprintf(log, format...)
	text = fmt.Sprintf("%v:ERROR %v", getFileLine(), text)
	fmt.Println(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", textColorRed, text))
}
func getFileLine() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}
