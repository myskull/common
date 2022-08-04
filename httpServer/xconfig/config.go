package xconfig

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func New(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("加载配置文件失败:%v\n", err)
		return
	}
	reg, _ := regexp.Compile(`(\n\r|\n|\r)`)
	list := reg.Split(string(b), -1)
	var section = ""
	for _, row := range list {
		row = strings.ReplaceAll(row, " ", "")
		if row == "" || strings.HasPrefix(row, "#") {
			continue
		}
		if strings.HasPrefix(row, "[") && strings.HasSuffix(row, "]") {
			// 区间配置开始
			section = row[1 : len(row)-1]
			continue
		}
		// 分割
		index := strings.Index(row, "=")
		if index < 0 {
			continue
		}
		Set(section, row[0:index], row[index+1:])
	}
}
