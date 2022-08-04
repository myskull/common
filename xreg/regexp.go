package xreg

import "regexp"

func IsNumber(param string) bool {
	reg, _ := regexp.Compile(`^[\-0-9]+$`)
	return reg.MatchString(param)
}

func IsId(param string) bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(param)
}

func IsFloat(param string) bool {
	reg, _ := regexp.Compile(`^[\-0-9.]+$`)
	return reg.MatchString(param)
}
