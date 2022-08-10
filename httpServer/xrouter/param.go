package xrouter

/**
参数自动校验
*/
type Param struct {
	Title   string // 标题
	Key     string // 参数的名字
	Comment string // 参数备注  生成API用到
	Type    int    // 参数类型
	Default string // 默认值
	Require bool   // 是否是必填
	Func    func(value string) bool
}

// 参数校验
func (this Param) Check(value string) bool {
	switch this.Type {
	case PTYPE_USERNAME:
		return checkUsername(value)
	case PTYPE_PASSWORD:
		return checkPassword(value)
	case PTYPE_INT:
		return checkInt(value)
	case PTYPE_UINT:
		return checkUint(value)
	case PTYPE_ID:
		return checkID(value)
	case PTYPE_NUMBERS:
		return checkNumbers(value)
	case PTYPE_DIY:
		if this.Func == nil {
			return false
		}
		return this.Func(value)
	}
	return false
}
