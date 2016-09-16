/*
Package wechat provide wechat sdk for go

使用Golang开发的微信SDK，简单、易用。

以下是一个处理消息接收以及回复的例子：

```
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
    return &message.Reply{message.MsgText, text}
})

server.Serve()
server.Send()

```

更多信息：https://github.com/silenceper/wechat

*/
package wechat
