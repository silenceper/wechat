# WeChat SDK for Go

使用Golang开发的微信SDK，简单、易用。

## 快速开始

以下是一个处理消息接收以及回复的例子：

```go

	//配置微信参数
	config := &wechat.Config{
		AppID:          "xxxx",
		AppSecret:      "xxxx",
		Token:          "xxxx",
		EncodingAESKey: "xxxx",
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

## 更多API使用

[文档地址](https://github.com/gowechat/docs)

## 已实现的微信平台接口

- 消息接收（包含事件的推送）
- 消息回复（文本,图片，语音，视频，音乐，图文）
- 网页授权Oauth2 
- 网页Js-SDK
- 素材管理

持续完善中...


## License

Apache License, Version 2.0
