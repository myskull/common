package httpServer

import (
	"fmt"
	"github.com/myskull/common/httpServer/xauth"
	"github.com/myskull/common/httpServer/xparam"
	"github.com/myskull/common/httpServer/xresp"
	"github.com/myskull/common/httpServer/xrouter"
	"testing"
)

var routes = []string{
	//"/\\w",
	"/ww/[\\d\\w]+",
	"/request",
}

type XRoute struct {
	route    string // 路由
	callback func() // 回调函数
}

// 请求参数
type XParam struct {
}

type PathVarable struct {
	Path string // 路径
}

func parsePathVarable(path string) *PathVarable {
	pathVarable := &PathVarable{
		Path: path,
	}
	return pathVarable
}

func TestAA(test *testing.T) {

	xrouter.Register(xrouter.XRouter{
		Router: "/request",
		Callback: func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp {
			fmt.Println("老了老弟1")
			return xresp.Success()
		},
		Method:  "",
		IsLogin: true,
	})
	xrouter.Register(xrouter.XRouter{
		Router: "/location",
		Callback: func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp {
			ip := xparam.IP()
			fmt.Println(ip)
			return xresp.Success()
		},
		Method: "",
	})
	xrouter.Register(xrouter.XRouter{
		Title:  "用户列表",
		Router: "/users/[\\d]+/[\\d]+",
		Callback: func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp {
			fmt.Printf("老了老弟:[%v]\n", xparam.XPath().Varable(0).Get())
			fmt.Printf("老了老弟:[%v]\n", xparam.XPath().Varable(1).Get())
			fmt.Printf("老了老弟:[%v]\n", xparam.XPath().Varable(2).Get())
			fmt.Printf("老了老弟:[%v]\n", xparam.XPath().Varable(3).Get())
			fmt.Printf("老了老弟:[%v]\n", xparam.XPath().Get())
			return xresp.Success()
		},
		Method: "",
		//IsLogin: true,
	})
	xrouter.Register(xrouter.XRouter{
		Router: "/request/[\\d]+",
		Callback: func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp {
			fmt.Println("老了老弟2", xparam.XPath().Uint())
			return xresp.Success()
		},
		Method:  "",
		IsLogin: true,
	})
	xrouter.Register(xrouter.XRouter{
		Title:  "登录",
		Router: "/login",
		Callback: func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp {
			uid := xparam.Uint32("uid")
			xauth.Set(&xauth.XAuth{ID: uid})
			return xresp.Success()
		},
		Params: []xrouter.Param{
			{Title: "账号", Key: "username", Comment: "用户的账号", Type: xrouter.PTYPE_USERNAME},
			{Title: "密码", Key: "password", Comment: "用户的密码", Type: xrouter.PTYPE_PASSWORD},
		},
		Method: "",
	})
	xauth.Check = func(param *xparam.XParam) *xauth.XAuth {
		uid := param.Uint32("uid")
		// 数据库， redis的操作
		return xauth.Get(uid)
	}
	Start(8000)
}
