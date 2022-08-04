package xparam

import (
	"fmt"
	"gitee.com/myskull/common/xparse"
	"strings"
)

type XPath struct {
	path string
	data string
	list []string // 分割列表
}

// 专门读取转后一位参数的结构体
func parseXPath(path string) XPath {
	xpath := XPath{
		path: path,
		data: "",
		list: []string{},
	}
	index := strings.LastIndex(path, "/")
	xpath.data = path[index+1:]
	xpath.list = strings.Split(path, "/")
	return xpath
}

func (this XPath) Get() string {
	return this.data
}

func (this XPath) Uint() uint {
	return xparse.Uint(this.data)
}
func (this XPath) Uint8() uint8 {
	return xparse.Uint8(this.data)
}

func (this XPath) Uint32() uint32 {
	return xparse.Uint32(this.data)
}

func (this XPath) Uint64() uint64 {
	return xparse.Uint64(this.data)
}

func (this XPath) Int() int {
	return xparse.Int(this.data)
}

func (this XPath) Int8() int8 {
	return xparse.Int8(this.data)
}

func (this XPath) Int32() int32 {
	return xparse.Int32(this.data)
}

func (this XPath) Int64() int64 {
	return xparse.Int64(this.data)
}

func (this XPath) Float32() float32 {
	return xparse.Float32(this.data)
}

func (this XPath) Float64() float64 {
	return xparse.Float64(this.data)
}

// 0不变，
func (this XPath) Varable(index int) XPath {
	// /users/1/2  index 从后面数
	path := "/"
	i := 0
	for _, row := range this.list {
		if i <= index {
			i++
			continue
		}
		path += row
		break
	}
	return parseXPath(fmt.Sprintf("/%v", path))
}
