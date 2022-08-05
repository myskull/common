package xjson

import (
	"fmt"
	"github.com/myskull/common/xdate"
	"github.com/myskull/common/xparse"
	"strings"
	"time"
)

type A []M
type M map[string]interface{}

// 读取字符串
func (this M) Get(key string, _def ...string) string {
	def := ""
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return fmt.Sprint(val)
	} else {
		return def
	}
}

// 整形
func (this M) Int(key string, _def ...int) int {
	def := int(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Int(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Int8(key string, _def ...int8) int8 {
	def := int8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Int8(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Int32(key string, _def ...int32) int32 {
	def := int32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Int32(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Int64(key string, _def ...int64) int64 {
	def := int64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Int64(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Uint(key string, _def ...uint) uint {
	def := uint(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Uint(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Uint8(key string, _def ...uint8) uint8 {
	def := uint8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Uint8(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Uint32(key string, _def ...uint32) uint32 {
	def := uint32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Uint32(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Uint64(key string, _def ...uint64) uint64 {
	def := uint64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Uint64(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Bool(key string, _def ...bool) bool {
	def := false
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Bool(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Float32(key string, _def ...float32) float32 {
	def := float32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Float32(fmt.Sprint(val))
	} else {
		return def
	}
}

// 整形
func (this M) Float64(key string, _def ...float64) float64 {
	def := float64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	val, ok := this[key]
	if ok {
		return xparse.Float64(fmt.Sprint(val))
	} else {
		return def
	}
}

func (this M) Exists(key string) bool {
	_, ok := this[key]
	return ok
}

// 时间戳转换日期格式
func (this M) Date(key string, _format ...string) string {
	format := "2006-01-02 15:04:05"
	if len(_format) > 0 {
		format = _format[0]
	}
	timestamp := this.Int64(key)
	return xdate.ToDate(timestamp, format)
}

// 日期格式转换时间戳
func (this M) Timestamp(key string, _def ...int64) int64 {
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
func (this M) NumberList(key string, _sep ...string) []int32 {
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
func (this M) Splits(key string, sep string) []string {
	str := this.Get(key)
	if str != "" {
		return strings.Split(str, sep)
	} else {
		return []string{}
	}
}
