# 微信SDK

此工程为 [https://github.com/silenceper/wechat](https://github.com/silenceper/wechat)的二开版本

主场开发场景是多第三方平台为起点控制多小程序包含支付全套流程中间件

因此本工程开发线路也是围绕着主场工程线路展开，有坑排坑，有缺补缺

### 主场工程已开源

[![纸喵软件/wechat](https://gitee.com/zhimiao/wechat/widgets/widget_card.svg?colors=4183c4,ffffff,ffffff,e3e9ed,666666,9b9b9b)](https://gitee.com/zhimiao/wechat)

## 快速开始

sdk实例获取

```go

// memcache := cache.NewMemcache("127.0.0.1:11211")
memcache := chache.NewMemory()

wcConfig := &wechat.Config{
	AppID:          cfg.AppID,
	AppSecret:      cfg.AppSecret,
	Token:          cfg.Token,
	EncodingAESKey: cfg.EncodingAESKey,//消息加解密时用到
	Cache:          memcache,
}
```

微信通知接收

```go

//配置微信参数
config := &wechat.Config{
	AppID:          "xxxx",
	AppSecret:      "xxxx",
	Token:          "xxxx",
	EncodingAESKey: "xxxx",
	Cache:          cache.NewMemory(), // 使用memory保存access_token，也可选择redis或自定义cache
}
wc := wechat.NewWechat(config)

// 传入request和responseWriter
server := wc.GetServer(request, responseWriter)
server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

	//回复消息：演示回复用户发送的消息
	text := message.NewText(msg.Content)
	return &message.Reply{message.MsgTypeText, text}
})

server.Serve()
server.Send()

```

#### 和主流框架配合使用

主要是request和responseWriter在不同框架中获取方式可能不一样：

- Gin Framework: [./examples/gin/gin.go](./examples/gin/gin.go)

**Cache 设置**

Cache主要用来保存全局access_token以及js-sdk中的ticket：
默认采用memcache存储。当然也可以直接实现`cache/cache.go`中的接口


> 缓存字典

| key | 备注 |
|:------|:-------|
| qy_access_token_${小程序APPID} | 小程序token |
| authorizer_access_token_${小程序APPID} | 代小程序accesstoken |
| component_access_token_${平台APPID} | 代小程序accesstoken |
| component_verify_ticket_${平台APPID} | 第三方平台票据 |


更多API使用请参考 godoc ：
[https://godoc.org/gitee.com/zhimiao/wechat-sdk](https://godoc.org/gitee.com/zhimiao/wechat-sdk)

## License

Apache License, Version 2.0

## Third Party Softwares

This software uses the following third party open source components.  
The third party licensors of these components may provide additional license rights,  
terms and conditions and/or require certain notices as described below.

* [wechat](https://github.com/silenceper/wechat), licensed under the [Apache License 2.0](https://github.com/silenceper/wechat/blob/master/LICENSE)  
Copyright (c) silenceper 