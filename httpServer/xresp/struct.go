package xresp

type XResp struct {
	Code       int         `json:"code"`    // 状态码
	Message    string      `json:"message"` // 提示信息
	Data       interface{} `json:"data"`    // 扩展数据
	IsLocation bool        `json:"-"`
	Url        string      `json:"-"` // 重定向的地址
}

// 404 not found
func X404() XResp {
	return XResp{
		Code: 404,
	}
}

func SystemError() XResp {
	return XResp{
		Code:    50000,
		Message: "系统异常,请联系管理员处理!",
	}
}

func ServerError() XResp {
	return XResp{
		Code:    50001,
		Message: "服务器繁忙,请稍后再试!",
	}
}

func ParamError() XResp {
	return XResp{
		Code:    50002,
		Message: "参数不合法!",
	}
}

func NoLogin() XResp {
	return XResp{
		Code:    40000,
		Message: "请先登录!",
	}
}

func Success(_data ...interface{}) XResp {
	resp := XResp{
		Code:    0,
		Message: "ok!",
	}
	if len(_data) > 0 {
		resp.Data = _data[0]
	}
	return resp
}

func Failed(message string, _code ...int) XResp {
	code := 1
	if len(_code) > 0 {
		code = _code[0]
	}
	return XResp{
		Code:    code,
		Message: message,
	}
}

func JCode(code int, message string, data interface{}) XResp {
	return XResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Location(url string) XResp {
	return XResp{
		IsLocation: true,
		Url:        url,
	}
}
