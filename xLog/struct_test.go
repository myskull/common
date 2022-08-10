package xLog

import (
	"fmt"
	"testing"
)

func TestSetLevel(t *testing.T) {
	fmt.Println(logLevel)
}

func TestInfo(t *testing.T) {
	//SetLevel(LOG_DEBUG)
	//SetLevel(LOG_INFO + LOG_DEBUG)
	Info("蓝色%v", "红红")
	Debug("蓝色%v", "红红")
	Warning("蓝色%v", "红红")
	Error("蓝色%v", "红红")
}
