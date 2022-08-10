package wxsocket

import (
	"github.com/myskull/common/socketServer/wxauth"
	"github.com/myskull/common/xconfig"
	"golang.org/x/net/websocket"
	"net/url"
	"strings"
)

type XSocket struct {
	ws        *websocket.Conn `json:"-"`
	pingTime  int64           `json:"-"` // 上一次心跳时间
	auth      *wxauth.WxAuth  `json:"-"`
	uid       uint32          `json:"uid"`
	urlParam  url.Values
	UserAgent string
	IP        string
	id        string //链接ID
}

var connect_type = 0 // 默认方式，明文链接
func New(ws *websocket.Conn) *XSocket {
	socket := &XSocket{
		ws: ws,
	}
	socket.initRequest()
	return socket
}

func (this *XSocket) initRequest() {
	connect_type = xconfig.GetInt("websocket", "connect_type", 0)
	if connect_type == 0 {
		query_param_str := ""
		index := strings.Index(this.ws.Request().RequestURI, "?")
		if index >= 0 {
			query_param_str = this.ws.Request().RequestURI[index+1:]
		}
		url_param, err := url.ParseQuery(query_param_str)
		if err != nil {
			return
		}
		this.urlParam = url_param
	}
	this.UserAgent = this.ws.Request().UserAgent()
	this.IP = this.ws.Request().RemoteAddr
}

// 登录认证
func (this *XSocket) Auth() bool {
	auth := wxauth.Get(this.uid)
	if auth == nil {
		return false
	}
	this.auth = auth
	return true
}
