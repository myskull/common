package xconfig

import (
	"github.com/myskull/common/xparse"
	"sync"
)

/**
专门读取配置的文件
*/
type XConfig struct {
	data map[string]map[string]string
}

var config = XConfig{
	data: map[string]map[string]string{},
}
var locker sync.RWMutex

func Set(section, key, value string) {
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		config.data[section] = map[string]string{}
	}
	config.data[section][key] = value
}

// 读取字符串
func GetStr(section, key string, _def ...string) string {
	def := ""
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return config.data[section][key]
}

func GetInt(section, key string, _def ...int) int {
	def := int(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Int(config.data[section][key])
}

func GetInt8(section, key string, _def ...int8) int8 {
	def := int8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Int8(config.data[section][key])
}

func GetInt32(section, key string, _def ...int32) int32 {
	def := int32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Int32(config.data[section][key])
}

func GetInt64(section, key string, _def ...int64) int64 {
	def := int64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Int64(config.data[section][key])
}

func GetUint(section, key string, _def ...uint) uint {
	def := uint(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Uint(config.data[section][key])
}

func GetUint8(section, key string, _def ...uint8) uint8 {
	def := uint8(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Uint8(config.data[section][key])
}

func GetUint32(section, key string, _def ...uint32) uint32 {
	def := uint32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Uint32(config.data[section][key])
}

func GetUint64(section, key string, _def ...uint64) uint64 {
	def := uint64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Uint64(config.data[section][key])
}

func GetFloat32(section, key string, _def ...float32) float32 {
	def := float32(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Float32(config.data[section][key])
}

func GetFloat64(section, key string, _def ...float64) float64 {
	def := float64(0)
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Float64(config.data[section][key])
}

func GetBool(section, key string, _def ...bool) bool {
	def := false
	if len(_def) > 0 {
		def = _def[0]
	}
	locker.Lock()
	defer locker.Unlock()
	if _, ok := config.data[section]; !ok {
		return def
	}
	return xparse.Bool(config.data[section][key])
}
