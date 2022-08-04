package xrouter

import (
	"fmt"
	"gitee.com/myskull/common/httpServer/xauth"
	"gitee.com/myskull/common/httpServer/xparam"
	"gitee.com/myskull/common/httpServer/xresp"
	"regexp"
	"sync"
)

type XRouter struct {
	Title    string                                                     // 路由的名字
	Router   string                                                     // 路由
	Callback func(xparam *xparam.XParam, auth *xauth.XAuth) xresp.XResp // 回调函数
	Method   string
	IsLogin  bool    // 是否验证登录
	Params   []Param // 自动校验参数列表, 不校验也可以拿到其他的参数. 这里主要是为了校验其他数据格式
}

var RouterPools = make(map[string]*XRouter) // 路由函数
var locker sync.RWMutex

// 根据正则匹配，检测出符合的一个路由
func Get(method, path string) *XRouter {
	locker.RLock()
	defer locker.RUnlock()
	for _, router := range RouterPools {
		if method != router.Method && router.Method != "" {
			continue
		}
		reg, err := regexp.Compile(fmt.Sprintf(`^%v$`, router.Router))
		if err != nil {
			return nil
		}
		if reg.MatchString(path) {
			// 最后一个参数就是pathvarable的值
			return router
		}
	}
	return nil
}

func Register(router XRouter) {
	locker.Lock()
	defer locker.Unlock()
	RouterPools[router.Method+"_"+router.Router] = &router
}
