package xparse

import "strconv"

/**
数据格式化
*/
func Int(param string) int {
	i, _ := strconv.Atoi(param)
	return i
}

func Int8(param string) int8 {
	i, _ := strconv.ParseInt(param, 10, 8)
	return int8(i)
}

func Int16(param string) int16 {
	i, _ := strconv.ParseInt(param, 10, 16)
	return int16(i)
}

func Int32(param string) int32 {
	i, _ := strconv.ParseInt(param, 10, 32)
	return int32(i)
}

func Int64(param string) int64 {
	i, _ := strconv.ParseInt(param, 10, 32)
	return i
}

func Uint(param string) uint {
	i, _ := strconv.Atoi(param)
	return uint(i)
}

func Uint8(param string) uint8 {
	i, _ := strconv.ParseUint(param, 10, 8)
	return uint8(i)
}

func Uint16(param string) uint16 {
	i, _ := strconv.ParseUint(param, 10, 16)
	return uint16(i)
}

func Uint32(param string) uint32 {
	i, _ := strconv.ParseUint(param, 10, 32)
	return uint32(i)
}

func Uint64(param string) uint64 {
	i, _ := strconv.ParseUint(param, 10, 32)
	return i
}

func Bool(param string) bool {
	i, _ := strconv.ParseBool(param)
	return i
}

func Float32(param string) float32 {
	i, _ := strconv.ParseFloat(param, 32)
	return float32(i)
}

func Float64(param string) float64 {
	i, _ := strconv.ParseFloat(param, 64)
	return i
}
