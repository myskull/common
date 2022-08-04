package xauth

import (
	"encoding/json"
	"fmt"
	"gitee.com/myskull/common/httpServer/xparam"
	"gitee.com/myskull/common/httpServer/xredis"
)

const (
	XAUTH_REDIS_KEY = "xauth_redis_key_%v"
)

type XAuth struct {
	ID     uint32      `json:"id"`
	Extend interface{} // 扩展数据
}

var Check = func(param *xparam.XParam) *XAuth {
	return Get(0)
}

var Set = func(auth *XAuth) error {
	b, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	err = xredis.Set(fmt.Sprintf(XAUTH_REDIS_KEY, auth.ID), string(b), 60*5)
	if err != nil {
		fmt.Println("保存用户信息失败:", err.Error())
	}
	return err
}
var Get = func(id uint32) *XAuth {
	auth_str := xredis.Get(fmt.Sprintf(XAUTH_REDIS_KEY, id))
	if auth_str == "" {
		return nil
	}
	var auth = &XAuth{}
	err := json.Unmarshal([]byte(auth_str), auth)
	if err != nil {
		return nil
	}
	return auth
}
var Del = func(id uint32) {
	xredis.Del(fmt.Sprintf(XAUTH_REDIS_KEY, id))
}
