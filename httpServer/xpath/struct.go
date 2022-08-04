package xpath

import (
	"github.com/myskull/common/xparse"
	"strings"
)

type XPath struct {
	path    string // 路径
	varable string // 变量值
	isParse bool   // 是否已经解析
}

func Parse(path string) *XPath {
	result := &XPath{
		path: path,
	}
	return result

}

func (this *XPath) parse() {
	if this.isParse {
		return
	}
	index := strings.LastIndex(this.path, "/")
	if index > -1 {
		this.varable = this.path[index+1:]
	}
	this.isParse = true
}

func (this *XPath) Int() int {
	this.parse()
	return xparse.Int(this.varable)
}

func (this *XPath) Int8() int8 {
	this.parse()
	return xparse.Int8(this.varable)
}

func (this *XPath) Int32() int32 {
	this.parse()
	return xparse.Int32(this.varable)
}

func (this *XPath) Int64() int64 {
	this.parse()
	return xparse.Int64(this.varable)
}

func (this *XPath) Uint() uint {
	this.parse()
	return xparse.Uint(this.varable)
}

func (this *XPath) Uint8() uint8 {
	this.parse()
	return xparse.Uint8(this.varable)
}

func (this *XPath) Uint32() uint32 {
	this.parse()
	return xparse.Uint32(this.varable)
}

func (this *XPath) Uint64() uint64 {
	this.parse()
	return xparse.Uint64(this.varable)
}

func (this *XPath) Str() string {
	this.parse()
	return this.varable
}
