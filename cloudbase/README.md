# 小程序-云开发 SDK

[云开发（CloudBase）](https://tencentcloudbase.github.io/)是基于Serverless架构构建的一站式后端云服务，涵盖函数、数据库、存储、CDN等服务，免后端运维，支持小程序、Web和APP开发。 其中，小程序·云开发是微信和腾讯云联合推出的云端一体化解决方案，基于云开发可以免鉴权调用微信所有开放能力，在微信开发者工具中即可开通使用。



## 使用说明
**引入依赖**
>推荐使用go module 进行管理

```
go get github.com/silenceper/wechat@v1.2.3
```

**初始化配置**

```golang
//使用memcache保存access_token，也可选择redis或自定义cache
memCache=cache.NewMemcache("127.0.0.1:11211")

//配置小程序参数
config := &wechat.Config{
    AppID:     "your app id",
    AppSecret: "your app secret",
    Cache:     memCache,
}
wc := wechat.NewWechat(config)
wcTcb := wc.GetTcb()
```

### 使用API

#### 触发云函数
```golang
res, err := wcTcb.InvokeCloudFunction("test-xxxx", "add", `{"a":1,"b":2}`)
if err != nil {
    panic(err)
}
```

更多使用方法参考[pkg.go.dev](https://pkg.go.dev/github.com/silenceper/wechat@v1.2.3/tcb?tab=doc#Tcb)

## Demo
### 使用wechat sdk开发一个留言板

这是一个使用wechat sdk来完成一个留言板的例子，使用到了云开发中的云函数，数据库，存储API：

- [起步：项目搭建](./guestbook-demo/start.md)
- [数据库：调用云开发数据库实现文本保存](./guestbook-demo/database.md)
- [云函数：调用云函数实现文本过滤](./guestbook-demo/cloudfunctions.md)
- [云开发存储：实现留言本附件上传](./guestbook-demo/storage.md)

以上文中的所有代码都上传在 [https://github.com/go-demo/guestbook](https://github.com/go-demo/guestbook)