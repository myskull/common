package httpServer

import (
	"encoding/json"
	"fmt"
	"github.com/myskull/common/httpServer/xLog"
	"github.com/myskull/common/httpServer/xauth"
	"github.com/myskull/common/httpServer/xconfig"
	"github.com/myskull/common/httpServer/xparam"
	"github.com/myskull/common/httpServer/xresp"
	"github.com/myskull/common/httpServer/xrouter"
	"net/http"
)

func Start(port int, file string) {
	if file != "" {
		xconfig.New(file)
	}
	if xconfig.KeyExists("system", "logLevel") {
		logLevel := xconfig.GetInt("system", "logLevel", xLog.LOG_ALL)
		if logLevel > 0 {
			xLog.SetLevel(logLevel)
		}
	}
	http.HandleFunc("/", xHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func xHandler(w http.ResponseWriter, r *http.Request) {
	pathURL := r.URL.Path
	if pathURL == "/favicon.ico" {
		return
	}
	router := xrouter.Get(r.Method, pathURL)
	if router == nil {
		output(w, xresp.X404())
		return
	}
	param := xparam.New(r)
	if router.MaxMemory > 0 {
		param.SetMaxMemory(router.MaxMemory) // 通过路由设定文件上传
	}
	if len(router.Params) > 0 {
		// 需要自动校验参数
		for _, p := range router.Params {
			val := param.Get(p.Key, p.Default)
			if val != p.Default && !p.Check(val) {
				// 默认值不需要校验
				output(w, xresp.NoLogin())
				return
			}
		}
	}
	var auth *xauth.XAuth
	if router.IsLogin {
		auth = xauth.Check(param)
		if auth == nil {
			output(w, xresp.NoLogin())
			return
		}
	}
	var resp = xresp.XResp{}
	if router.Callback != nil {
		resp = router.Callback(param, auth)
	}
	output(w, resp)
}

func output(w http.ResponseWriter, resp xresp.XResp) {
	if resp.IsLocation {
		w.Header().Set("Cache-control", "must-revalidate,no-store")
		w.Header().Set("Content-Type", "text/html;charset=UTF-8")
		w.Header().Set("Location", resp.Url)
		w.WriteHeader(307)
	} else {
		b, err := json.Marshal(resp)
		if err != nil {
			xLog.Error("json序列化失败:", err.Error())
			w.WriteHeader(502)
		} else {
			w.Write(b)
		}
	}
}
