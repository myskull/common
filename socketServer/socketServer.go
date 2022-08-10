package socketServer

import (
	"fmt"
	"github.com/myskull/common/xLog"
	"golang.org/x/net/websocket"
	"net/http"
)

func Start(post uint32) {
	http.Handle("/websocket", websocket.Handler(xHandle))
	http.ListenAndServe(fmt.Sprintf(":%v", post), nil)
}

func xHandle(ws *websocket.Conn) {
	fmt.Println(ws.Request().RequestURI)
	fmt.Println(ws.Request().URL)
	fmt.Println(ws.Request().UserAgent())
	defer ws.Close()
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			xLog.Error("读取消息失败:%v", err)
			break
		}
	}
}
