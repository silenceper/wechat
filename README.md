## 📢 注意： 此分支为v1版本，已不再维护更新，请切换至 [v2](https://github.com/silenceper/wechat/tree/release-2.0)！

# WeChat SDK for Go
[![Build Status](https://travis-ci.org/silenceper/wechat.svg?branch=master)](https://travis-ci.org/silenceper/wechat)
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/wechat)](https://goreportcard.com/report/github.com/silenceper/wechat)
[![GoDoc](http://godoc.org/github.com/silenceper/wechat?status.svg)](http://godoc.org/github.com/silenceper/wechat)

使用Golang开发的微信SDK，简单、易用。

## 快速开始

以下是一个处理消息接收以及回复的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
memCache=cache.NewMemcache("127.0.0.1:11211")

//配置微信参数
config := &wechat.Config{
	AppID:          "xxxx",
	AppSecret:      "xxxx",
	Token:          "xxxx",
	EncodingAESKey: "xxxx",
	Cache:          memCache
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
完整代码：[examples/http/http.go](./examples/http/http.go)

#### 和主流框架配合使用

主要是request和responseWriter在不同框架中获取方式可能不一样：

- Beego: [./examples/beego/beego.go](./examples/beego/beego.go)
- Gin Framework: [./examples/gin/gin.go](./examples/gin/gin.go)

## 基本配置

```go
memcache := cache.NewMemcache("127.0.0.1:11211")

wcConfig := &wechat.Config{
	AppID:          cfg.AppID,
	AppSecret:      cfg.AppSecret,
	Token:          cfg.Token,
	EncodingAESKey: cfg.EncodingAESKey,//消息加解密时用到
	Cache:          memcache,
}
```

**Cache 设置**

Cache主要用来保存全局access_token以及js-sdk中的ticket：
默认采用memcache存储。当然也可以直接实现`cache/cache.go`中的接口


## 基本API使用

- [消息管理](#消息管理)
	- 接收普通消息
	- 接收事件推送
	- 被动回复消息
		- 回复文本消息
		- 回复图片消息
		- 回复视频消息
		- 回复音乐消息
		- 回复图文消息
- [自定义菜单](#自定义菜单)
	- 自定义菜单创建接口
	- 自定义菜单查询接口
	- 自定义菜单删除接口
	- 自定义菜单事件推送
	- 个性化菜单接口
		- 添加个性化菜单
		- 删除个性化菜单
		- 测试个性化菜单匹配结果
	- 获取公众号菜单配置
- [微信网页开发](#微信网页开发)
	- Oauth2 授权
		- 发起授权
		- 通过code换取access_token
		- 拉取用户信息
		- 刷新access_token
		- 检验access_token是否有效
	- 获取js-sdk配置
- [素材管理](#素材管理)
- [小程序开发](#小程序开发)
- [小程序-云开发](./tcb)

## 消息管理

通过`wechat.GetServer(request,responseWriter)`获取到server对象之后

调用`SetMessageHandler(func(msg message.MixMessage){})`设置消息的处理函数，函数参数为message.MixMessage 结构如下：

```go
//MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken

	//基本消息
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        float64 `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`

	//事件相关
	Event     string `xml:"Event"`
	EventKey  string `xml:"EventKey"`
	Ticket    string `xml:"Ticket"`
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`

	MenuID    string `xml:"MenuId"`

	//扫码事件
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	//发图事件
	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	//发送地理位置事件
	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}
}

```

具体参数请参考微信文档：[接收普通消息
](http://mp.weixin.qq.com/wiki/17/f298879f8fb29ab98b2f2971d42552fd.html)

### 接收普通消息
```go
server.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		switch v.MsgType {
		//文本消息
		case message.MsgTypeText:
			//do something

			//图片消息
		case message.MsgTypeImage:
			//do something

			//语音消息
		case message.MsgTypeVoice:
			//do something

			//视频消息
		case message.MsgTypeVideo:
			//do something

			//小视频消息
		case message.MsgTypeShortVideo:
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something

			//事件推送消息
		case message.MsgTypeEvent:

		}
}


```


### 接收事件推送
```go
//事件推送消息
case message.MsgTypeEvent:
	switch v.Event {
		//EventSubscribe 订阅
		case message.EventSubscribe:
			//do something

			//取消订阅
		case message.EventUnsubscribe:
			//do something

			//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
		case message.EventScan:
			//do something

			// 上报地理位置事件
		case message.EventLocation:
			//do something

			// 点击菜单拉取消息时的事件推送
		case message.EventClick:
			//do something

			// 点击菜单跳转链接时的事件推送
		case message.EventView:
			//do something

			// 扫码推事件的事件推送
		case message.EventScancodePush:
			//do something

			// 扫码推事件且弹出“消息接收中”提示框的事件推送
		case message.EventScancodeWaitmsg:
			//do something

			// 弹出系统拍照发图的事件推送
		case message.EventPicSysphoto:
			//do something

			// 弹出拍照或者相册发图的事件推送
		case message.EventPicPhotoOrAlbum:
			//do something

			// 弹出微信相册发图器的事件推送
		case message.EventPicWeixin:
			//do something

			// 弹出地理位置选择器的事件推送
		case message.EventLocationSelect:
			//do something

	}


```


### 被动回复消息

回复消息需要返回 `*message.Reply` 对象结构体如下：

```go
type Reply struct {
	MsgType MsgType  //消息类型
	MsgData interface{}  //消息结构
}

```
注意：`return nil`表示什么也不做

####  回复文本消息
```go
	text := message.NewText("回复文本消息")
	return &message.Reply{message.MsgTypeText, text}
```
####  回复图片消息
```go
//mediaID 可通过素材管理-上上传多媒体文件获得
image :=message.NewImage("mediaID")
return &message.Reply{message.MsgTypeImage, image}
```
####  回复视频消息
```go
video := message.NewVideo("mediaID", "视频标题", "视频描述")
return &message.Reply{message.MsgTypeVideo, video}
```
####  回复音乐消息
```go
music := message.NewMusic("title", "description", "musicURL", "hQMusicURL", "thumbMediaID")
return &message.Reply{message.MsgTypeMusic,music}
```
**字段说明：**

Title:音乐标题

Description:音乐描述

MusicURL:音乐链接

HQMusicUrl：高质量音乐链接，WIFI环境优先使用该链接播放音乐

ThumbMediaId：缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id

####  回复图文消息

```go
articles := make([]*message.Article, 1)

article := new(message.Article)
article.Title = "标题"
article.Description = "描述信息信息信息"
article.PicURL = "http://ww1.sinaimg.cn/large/65209136gw1f7vhjw95eqj20wt0zk40z.jpg"
article.URL = "https://github.com/silenceper/wechat"
articles[0] = article

news := message.NewNews(articles)
return &message.Reply{message.MsgTypeNews,news}
```
**字段说明：**

Title：图文消息标题

Description：图文消息描述

PicUrl	：图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200

Url	：点击图文消息跳转链接


## 自定义菜单

通过` wechat.GetMenu()`获取menu的实例

### 自定义菜单创建接口

以下是一个创建二级菜单的例子

```go
mu := wc.GetMenu()

buttons := make([]*menu.Button, 1)
btn := new(menu.Button)

//创建click类型菜单
btn.SetClickButton("name", "key123")
buttons[0] = btn

//设置btn为二级菜单
btn2 := new(menu.Button)
btn2.SetSubButton("subButton", buttons)

buttons2 := make([]*menu.Button, 1)
buttons2[0] = btn2

//发送请求
err := mu.SetMenu(buttons2)
if err != nil {
	fmt.Printf("err= %v", err)
	return
}

```

**创建其他类型的菜单：**

```go
//SetViewButton view类型
func (btn *Button) SetViewButton(name, url string)

// SetScanCodePushButton 扫码推事件
func (btn *Button) SetScanCodePushButton(name, key string)

//SetScanCodeWaitMsgButton 设置 扫码推事件且弹出"消息接收中"提示框
func (btn *Button) SetScanCodeWaitMsgButton(name, key string)

//SetPicSysPhotoButton 设置弹出系统拍照发图按钮
func (btn *Button) SetPicSysPhotoButton(name, key string)

//SetPicPhotoOrAlbumButton 设置弹出拍照或者相册发图类型按钮
func (btn *Button) SetPicPhotoOrAlbumButton(name, key string) {

// SetPicWeixinButton 设置弹出微信相册发图器类型按钮
func (btn *Button) SetPicWeixinButton(name, key string)

// SetLocationSelectButton 设置 弹出地理位置选择器 类型按钮
func (btn *Button) SetLocationSelectButton(name, key string)

//SetMediaIDButton  设置 下发消息(除文本消息) 类型按钮
func (btn *Button) SetMediaIDButton(name, mediaID string)

//SetViewLimitedButton  设置 跳转图文消息URL 类型按钮
func (btn *Button) SetViewLimitedButton(name, mediaID string) {

```

### 自定义菜单查询接口

```go
mu := wc.GetMenu()
resMenu,err:=mu.GetMenu()
```
>返回结果 resMenu 结构参考 ./menu/menu.go 中ResMenu 结构体

### 自定义菜单删除接口

```go
mu := wc.GetMenu()
err:=mu.DeleteMenu()
```

### 自定义菜单事件推送

 请参考 消息管理 - 事件推送

### 个性化菜单接口
**添加个性化菜单**

```go

func (menu *Menu) AddConditional(buttons []*Button, matchRule *MatchRule) error
```

**删除个性化菜单**

```go
//删除个性化菜单
func (menu *Menu) DeleteConditional(menuID int64) error

```
**测试个性化菜单匹配结果**

```go
//菜单匹配
func (menu *Menu) MenuTryMatch(userID string) (buttons []Button, err error) {

```

### 获取公众号菜单配置

```go
//获取自定义菜单配置接口
func (menu *Menu) GetCurrentSelfMenuInfo() (resSelfMenuInfo ResSelfMenuInfo, err error)

```

## 微信网页开发

### Oauth2 授权

具体授权流程请参考微信文档：[网页授权](http://mp.weixin.qq.com/wiki/4/9ac2e7b1f1d22e9e57260f6553822520.html)

**1.发起授权**

```go
oauth := wc.GetOauth()
err := oauth.Redirect("跳转的绝对地址", "snsapi_userinfo", "123dd123")
if err != nil {
	fmt.Println(err)
}

```
> 如果不希望直接跳转，可通过 oauth.GetRedirectURL 获取跳转的url

**2.通过code换取access_token**

```go
code := c.Query("code")
resToken, err := oauth.GetUserAccessToken(code)
if err != nil {
	fmt.Println(err)
	return
}

```
**3.拉取用户信息(需scope为 snsapi_userinfo)**

```go
//getUserInfo
userInfo, err := oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(userInfo)

```
**刷新access_token**

```go
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error)

```
**检验access_token是否有效**

```go
func (oauth *Oauth) CheckAccessToken(accessToken, openID string) (b bool, err error)
```

### 获取js-sdk配置

```go
js := wc.GetJs()
cfg, err := js.GetConfig("传入需要的调用js-sdk的url地址")
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(cfg)
```
其中返回的cfg结构体如下：

```go
type Config struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

```

## 素材管理

[素材管理API](https://godoc.org/github.com/silenceper/wechat/material#Material)

### 批量获取永久素材

```go
list, err := wc.GetMaterial().BatchGetMaterial(material.PermanentMaterialTypeNews, 0, 10)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(list)
```

## 小程序开发

获取小程序操作对象

``` go
memCache=cache.NewMemcache("127.0.0.1:11211")
config := &wechat.Config{
	AppID:     "xxx",
	AppSecret: "xxx",
	Cache:     memCache=cache.NewMemcache("127.0.0.1:11211"),
}
wc := wechat.NewWechat(config)

wxa := wc.GetMiniProgram()
```

### 小程序登录凭证校验

``` go
func (wxa *MiniProgram) Code2Session(jsCode string) (result ResCode2Session, err error)
```

### 小程序数据统计

**获取用户访问小程序日留存**

``` go
func (wxa *MiniProgram) GetAnalysisDailyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error)
```

**获取用户访问小程序月留存**

``` go
func (wxa *MiniProgram) GetAnalysisMonthlyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error)
```

**获取用户访问小程序周留存**

``` go
func (wxa *MiniProgram) GetAnalysisWeeklyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error)
```

**获取用户访问小程序数据概况**

``` go
func (wxa *MiniProgram) GetAnalysisDailySummary(beginDate, endDate string) (result ResAnalysisDailySummary, err error)
```

**获取用户访问小程序数据日趋势**

``` go
func (wxa *MiniProgram) GetAnalysisDailyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error)
```

**获取用户访问小程序数据月趋势**

``` go
func (wxa *MiniProgram) GetAnalysisMonthlyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error)
```

**获取用户访问小程序数据周趋势**

``` go
func (wxa *MiniProgram) GetAnalysisWeeklyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error)
```

**获取小程序新增或活跃用户的画像分布数据**

``` go
func (wxa *MiniProgram) GetAnalysisUserPortrait(beginDate, endDate string) (result ResAnalysisUserPortrait, err error)
```

**获取用户小程序访问分布数据**

``` go
func (wxa *MiniProgram) GetAnalysisVisitDistribution(beginDate, endDate string) (result ResAnalysisVisitDistribution, err error)
```

**获取小程序页面访问数据**

``` go
func (wxa *MiniProgram) GetAnalysisVisitPage(beginDate, endDate string) (result ResAnalysisVisitPage, err error)
```

### 小程序二维码生成

**获取小程序二维码，适用于需要的码数量较少的业务场景**

``` go
func (wxa *MiniProgram) CreateWXAQRCode(coderParams QRCoder) (response []byte, err error)
```

**获取小程序码，适用于需要的码数量较少的业务场景**

``` go
func (wxa *MiniProgram) GetWXACode(coderParams QRCoder) (response []byte, err error)
```

**获取小程序码，适用于需要的码数量极多的业务场景**

``` go
func (wxa *MiniProgram) GetWXACodeUnlimit(coderParams QRCoder) (response []byte, err error)
```


更多API使用请参考 godoc ：
[https://godoc.org/github.com/silenceper/wechat](https://godoc.org/github.com/silenceper/wechat)

## License

Apache License, Version 2.0
