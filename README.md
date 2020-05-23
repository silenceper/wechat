# WeChat SDK for Go
[![Build Status](https://travis-ci.org/silenceper/wechat.svg?branch=release-2.0)](https://travis-ci.org/silenceper/wechat)
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/wechat)](https://goreportcard.com/report/github.com/silenceper/wechat)
[![GoDoc](http://godoc.org/github.com/silenceper/wechat?status.svg)](http://godoc.org/github.com/silenceper/wechat)

使用Golang开发的微信SDK，简单、易用。


## 快速开始

以下是一个微信公众号处理消息接收以及回复的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
wc := wechat.NewWechat()
memory := cache.NewMemory()
cfg := &offConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Token:     "xxx",
    //EncodingAESKey: "xxxx",
    Cache: memory,
}
officialAccount := wc.GetOfficialAccount(cfg)

// 传入request和responseWriter
server := officialAccount.GetServer(req, rw)
//设置接收消息的处理方法
server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

    //回复消息：演示回复用户发送的消息
    text := message.NewText(msg.Content)
    return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
})

//处理消息接收以及回复
err := server.Serve()
if err != nil {
    fmt.Println(err)
    return
}
//发送回复的消息
server.Send()

```

## 文档
[Wechat SDK 2.0 文档](http://silenceper.com/wechat)


## 目录说明
- officialaccount: 微信公众号API
- miniprogram: 小程序API
- minigame:小游戏API
- pay:微信支付API
- opernplatform:开放平台API
- work:企业微信
- aispeech:智能对话

## 如何贡献
- 提交issue，描述需要贡献的内容
- 完成更改后，提交PR


## License

Apache License, Version 2.0
