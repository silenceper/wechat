#使用

1.代码移入GOPATH/src下

2.services/weixinService.go中添加微信公众号配置信息
```
config = &wechat.Config{
		AppID:          "your app id",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
}
```
3.运行代码