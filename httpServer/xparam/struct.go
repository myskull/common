package xparam

import (
	"encoding/json"
	"fmt"
	"gitee.com/myskull/common/xdate"
	"gitee.com/myskull/common/xparse"
	"net/http"
	"strings"
	"time"
)

type XParam struct {
	request   *http.Request
	data      map[string]string
	isParse   bool // 是否已经解析完成
	isBody    bool // 是否从body读取数据
	maxMemory int64
	body      []byte
}

func New(request *http.Request) *XParam {
	param := &XParam{
		request:   request,
		data:      map[string]string{},
		maxMemory: 32,
	}
	param.parse()
	return param
}

// 解析请求数据
func (this *XParam) parse() {
	if this.isParse {
		return
	}
	this.request.ParseForm()

	if len(this.request.Form) > 0 {
		for key, val := range this.request.Form {
			this.data[key] = val[0]
		}
	}
	if len(this.request.PostForm) > 0 {
		for key, val := range this.request.PostForm {
			this.data[key] = val[0]
		}
	}
	if this.request.Method != "GET" && len(this.request.PostForm) == 0 {
		// 啥都没有，估计要从body读取数据
		this.request.ParseMultipartForm(1 << this.maxMemory)
		//fmt.Println("文件上传", this.request.Header.Get("Content-type"))
		if this.request.MultipartForm != nil {
			for key, val := range this.request.MultipartForm.Value {
				this.data[key] = val[0]
			}
		} else {
			// 从body读取
			this.isBody = true
		}
	}
	this.isParse = true
}

// 通过参数序列化到对象
func (this *XParam) Unmarshal(v interface{}) error {
	if this.request.Method == "GET" {
	} else if len(this.request.PostForm) > 0 {
		//form=xxx
	} else {
		return json.Unmarshal(this.body, v)
	}
	return nil
}

// 读取字符串
func (this *XParam) Get(key string, _def ...string) string {
	def := ""
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return val
	} else {
		return def
	}
}

// 整形
func (this *XParam) Int(key string, _def ...int) int {
	def := int(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Int(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Int8(key string, _def ...int8) int8 {
	def := int8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Int8(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Int32(key string, _def ...int32) int32 {
	def := int32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Int32(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Int64(key string, _def ...int64) int64 {
	def := int64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Int64(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Uint(key string, _def ...uint) uint {
	def := uint(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Uint(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Uint8(key string, _def ...uint8) uint8 {
	def := uint8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Uint8(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Uint32(key string, _def ...uint32) uint32 {
	def := uint32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Uint32(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Uint64(key string, _def ...uint64) uint64 {
	def := uint64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Uint64(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Bool(key string, _def ...bool) bool {
	def := false
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Bool(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Float32(key string, _def ...float32) float32 {
	def := float32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Float32(val)
	} else {
		return def
	}
}

// 整形
func (this *XParam) Float64(key string, _def ...float64) float64 {
	def := float64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this.data[key]
	if ok {
		return xparse.Float64(val)
	} else {
		return def
	}
}

func (this *XParam) Exists(key string) bool {
	_, ok := this.data[key]
	return ok
}

// 时间戳转换日期格式
func (this *XParam) Date(key string, _format ...string) string {
	format := "2006-01-02 15:04:05"
	if len(_format) > 0 {
		format = _format[0]
	}
	timestamp := this.Int64(key)
	return xdate.ToDate(timestamp, format)
}

// 日期格式转换时间戳
func (this *XParam) Timestamp(key string, _def ...int64) int64 {
	def := time.Now().Unix()
	if len(_def) > 0 {
		def = _def[0]
	}
	if this.Exists(key) {
		return xdate.ToTime(this.Get(key))
	} else {
		return def
	}
}

// 数字列表
func (this *XParam) NumberList(key string, _sep ...string) []int32 {
	sep := ","
	if len(_sep) > 0 {
		sep = _sep[0]
	}
	list := []int32{}
	if !this.Exists(key) {
		return list
	} else {
		numbers := this.Get(key)
		if numbers != "" {
			rows := strings.Split(numbers, sep)
			for _, row := range rows {
				list = append(list, xparse.Int32(row))
			}
		}
		return list
	}
}

// 分割列表
func (this *XParam) Splits(key string, sep string) []string {
	str := this.Get(key)
	if str != "" {
		return strings.Split(str, sep)
	} else {
		return []string{}
	}
}

// 读取头
func (this *XParam) Header() http.Header {
	return this.request.Header
}

func (this *XParam) XPath() XPath {
	return parseXPath(this.request.URL.Path)
}

// 读取ip信息
func (this *XParam) IP() string {
	//fmt.Println(this.request.Header) // 还未从header读取
	fmt.Println(this.request.RemoteAddr)
	index := strings.LastIndex(this.request.RemoteAddr, ":")
	return this.request.RemoteAddr[0:index]
}

func (this *XParam) Host() string {
	return this.request.Host
}

func (this *XParam) Request() *http.Request {
	return this.request
}
