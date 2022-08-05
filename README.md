```aidl
路由注册 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
```
```aidl
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
```
```aidl
配置文件内容:
```
```aidl
[redis]
address=127.0.0.1:6379
[mysql]
database=数据库名字
hostname=数据库地址:数据库端口
username=数据库账号
password=数据库密码
[system]
# 日志等级备注: 2INFO  4DEBUG  8WARNING  16:ERROR
# INFO+DEBUG+WARNING+ERROR   表示全部输出
# INFO+DEBUG 表示只输出INFO跟ERRPR
# 0或则不填，表示默认
logLevel=
```
```aidl
服务启动
httpServer.Start(8000,"配置文件路径")
mysql操作:  xmysql.NewBuilder() 
redis: xredis.Get()
日志: xLog.Info
配置文件读取: xconfig.Get()
API返回: xresp.Success()
读取请求数据: xparam.XParam
用户登录检验: xauth.Check(param *xparam.XParam)
用户信息读取: xauth.Get()
用户信息保存: xauth.Set()
用户信息删除: xauth.Del()
重定义用户的存储:
    xauth.Get = function(id int64)*xauth.XAuth{
        return &xauth.XAuth{}
    }
    
    xauth.Set = function(xauth.XAuth)error{
        return nil
    }
    
    xauth.Del = function(id int64){
    }
    xauth.Check = func(param *xparam.XParam)*xauth.XAuth{
        return nil // nil表示登录失败
    }
```
```aidl
param xparam.XParam
读取路由地址的参数: 
    Router: /users/10/21
    param.XPath().Int()     // 读取到 21  自动读取到最后一个参数
    param.Varable(0).Get()  // 读取到 users
    param.Varable(1).Int()  // 读取到 10
    param.Varable(2).Int()  // 读取到 21
    param.Varable(3).Int()  // 读取到 0  没有参数对应，返回Int的默认值0
```