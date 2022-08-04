package xrouter

import "regexp"

const (
	PTYPE_DIY      = 0 // 自定义校验
	PTYPE_USERNAME = 1 // 账号校验
	PTYPE_PASSWORD = 2 // 密码校验
	PTYPE_INT      = 3 // 整数校验
	PTYPE_UINT     = 4 // 正整数校验
	PTYPE_NUMBERS  = 5 // 数字列表 1,2,3
	PTYPE_ID       = 6 // ID校验
)

func checkID(value string) bool {
	reg, _ := regexp.Compile(`^[1-9][0-9]*$`)
	return reg.MatchString(value)
}

func checkNumbers(value string) bool {
	reg, _ := regexp.Compile(`^([0-9]+(,)?)+$`)
	return reg.MatchString(value)
}

func checkInt(value string) bool {
	reg, _ := regexp.Compile(`^(\-)?[0-9]+$`)
	return reg.MatchString(value)
}

func checkUint(value string) bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(value)
}

func checkUsername(value string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9@_]{6,18}$`)
	return reg.MatchString(value)
}

func checkPassword(value string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9@_\!#\$%\^\&\*\(\)\~\+\,\.\<\>]{6,18}$`)
	return reg.MatchString(value)
}
