package wxauth

type WxAuth struct {
	Uid  uint32 `json:"uid"`
	Data interface{}
}

var Get = func(uid uint32) *WxAuth {
	return nil
}
