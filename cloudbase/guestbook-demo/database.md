# 数据库：调用云开发数据库实现文本保存

在这一节，我们主要描述如何利用`wechat sdk`将留言的内容保存在云开发数据库中。

<a name="RjeGP"></a>
## API说明
参考微信云开发文档 [数据库篇](https://developers.weixin.qq.com/minigame/dev/wxcloud/reference-http-api/database/#%E6%95%B0%E6%8D%AE%E5%BA%93)，可以先阅读其原始的api提供的方法和说明，在[SDK DOC](https://pkg.go.dev/github.com/silenceper/wechat/tcb?tab=doc#Tcb.DatabaseAdd)中都可以找到对应的方法以及参数。

主要利用到[SDK DOC](https://pkg.go.dev/github.com/silenceper/wechat/tcb?tab=doc#Tcb.DatabaseAdd)中的如下方法，其他方法可在文档中找到，sdk文档中以 `Database` 开头的方法即为数据库相关的方法调用。
```go
func (tcb *Tcb) DatabaseAdd(env, query string) (*DatabaseAddRes, error) //数据库内容保存
func (tcb *Tcb) DatabaseCount(env, query string) (*DatabaseCountRes, error)//数据库计数
func (tcb *Tcb) DatabaseQuery(env, query string) (*DatabaseQueryRes, error)//数据库内容查询
```
返回结果对应字段说明： [`DatabaseAddRes`](https://pkg.go.dev/github.com/silenceper/wechat/tcb?tab=doc#DatabaseAddRes) ， [`DatabaseCountRes`](https://pkg.go.dev/github.com/silenceper/wechat/tcb?tab=doc#DatabaseCountRes) , [`DatabaseQueryRes`](https://pkg.go.dev/github.com/silenceper/wechat/tcb?tab=doc#DatabaseQueryRes) 

<a name="VKuDQ"></a>
## 包引入
本例中引入的WeChat sdk版本为v1.2.3版本，通过如下方法引入
```bash
go get github.com/silenceper/wechat@v1.2.3
```
可以在go.mod文件中看到引入的包以及对应的版本：
```go
module github.com/go-demo/guestbook

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/silenceper/wechat v1.2.3 // indirect
)
```

<a name="m7jnm"></a>
## 保存至云开发数据库
<a name="a095b5de"></a>
### WeChat SDK配置
为了方便在其他方法中调用<br />创建config.go用于解析云开发对应的配置参数，appkey，app_secret等：

```go
package main

import (
	"io/ioutil"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/tcb"
	"gopkg.in/yaml.v2"
)

//Config 配置信息
type Config struct {
	TcbEnv    string `yaml:"tcb_env"`
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

var cfg *Config
var _ = getConfig()

//通过getConfig方法获取配置参数
func getConfig() *Config {
	if cfg != nil {
		return cfg
	}
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	cfg = &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

var wechatTcb *tcb.Tcb
var _ = getTcb()
//通过getTcb获取wechat sdk的配置参数
func getTcb() *tcb.Tcb {
	if wechatTcb != nil {
		return wechatTcb
	}
	memCache := cache.NewMemory()

	//配置小程序参数
	config := &wechat.Config{
		AppID:     getConfig().AppID,
		AppSecret: getConfig().AppSecret,
		Cache:     memCache,
	}
	wc := wechat.NewWechat(config)
	wechatTcb = wc.GetTcb()
	return wechatTcb
}
```

其中config.yaml写入三个配置参数：

```yaml
tcb_env: test-6ku2s //云开发环境
app_id: xxxxxx //云开发appid
app_secret: xxxxxxxxx //云开发对应的app secret
```

<a name="VbDg9"></a>
### 调用API
为了方便在其他方法中方便调用sdk中的方法，这里新建一个 `feedbackService` struct，创建对应的save方法用于保存留言, `feedback.go` :

```go
package main

import (
	"fmt"
	"time"
)

//FeedbackService service
type FeedbackService struct {
}

//NewFeedbackService new
func NewFeedbackService() *FeedbackService {
	return &FeedbackService{}
}

//Feedback 留言记录
type Feedback struct {
	Username   string `form:"username",json:"username"`
	Content    string `form:"content",json:"content"`
	FilePath   string `json:"filePath"`//文件路径
	FileID     string `json:"fileId"` //存放文件
	CreateTime string `json:"createTime"`
}

func (svc *FeedbackService) Save(feedback *Feedback) error {
	if feedback.Username == "" || feedback.Content == "" {
		return fmt.Errorf("用户名或留言内容不能为空")
	}
	query := `db.collection(\"%s\").add({
      data: [{
        username: \"%s\",
        content: \"%s\",
		filePath: \"%s\",
		fileId: \"%s\",
        createTime: \"%s\",
      }]
      })`
	feedback.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	query = fmt.Sprintf(query, "guestbook", feedback.Username, feedback.Content, feedback.FilePath, feedback.FileID, feedback.CreateTime)
	_, err := getTcb().DatabaseAdd(getConfig().TcbEnv, query)
	if err != nil {
		return err
	}
	return nil
}
```

其中对于数据库中的query语句可以参考[https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/database/add.html](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/database/add.html)

这里调用 `DatabaseAdd` 方法实现内容的保存。

<a name="8db0f827"></a>
## 接收表单提交内容
将表单提交的内容提交到 `/feedback` 路由中，并创建 `feedback`方法接收表单提交的参数，

```go
func main() {
	r := gin.Default()

	//包含html模板
	r.LoadHTMLGlob("./template/*")
	//渲染留言页面
	r.GET("/",index)
	//提交留言
	r.POST("/feedback", feedback)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

其中feedback方法如下：

```go
//接收提交的内容
func feedback(c *gin.Context) {
	//1、接收提交参数
	feedback := &Feedback{}
	err := c.Bind(feedback)
	if err != nil {
		showError(c, err)
		return
	}
	feedbackService := NewFeedbackService()
	//保存内容
	err = feedbackService.Save(feedback)
	if err != nil {
		showError(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
```

这里通过c.Bind方法将form表单中的内容绑定到 `Feedback` stuct中，再通过调用feedbackService中的Save方法对文本内容进行保存。
<a name="DlBj4"></a>
## 展示留言内容
留言内容的展示主要分为两步，一先从数据库展示出来，二是将留言内容展示在页面：<br />查询的sql语句为：
```go
db.collection("guestbook").orderBy('createTime','desc').skip(0).limit(10).get()
```

在 `feedback.go` 中新增List方法，其中参数传入skip和limit参数用于分页

```go
//List 文本列表
func (svc *FeedbackService) List(skip, limit int) ([]*Feedback, error) {
	query := fmt.Sprintf("db.collection(\"guestbook\").orderBy('createTime','desc').skip(%d).limit(%d).get()", skip, limit)
	data, err := getTcb().DatabaseQuery(getConfig().TcbEnv, query)
	if err != nil {
		return nil, err
	}
	feedbacks := make([]*Feedback, 0, len(data.Data))
	for _, v := range data.Data {
		feedbackItem := &Feedback{}
		err := json.Unmarshal([]byte(v), feedbackItem)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedbackItem)

	}
	//fmt.Println(data.Pager)
	return feedbacks, nil
}
```

这里主要调用 `DatabaseQuery` 方法对db进行查询。

在main.go中的index方法在从数据中获取的数据取出并渲染在index.html中：

```go
//首页
func index(c *gin.Context) {
	page := c.Query("page")
	//获取记录数量
	feedbackService := NewFeedbackService()
	count, err := feedbackService.Count()
	if err != nil {
		showError(c, err)
		return
	}
	limit := 10
	totalPage := math.Ceil(float64(count) / float64(limit))
	totalPageInt := int(totalPage)
	pageInt, _ := strconv.Atoi(page)
	if pageInt > totalPageInt {
		pageInt = totalPageInt
	}
	if pageInt < 1 {
		pageInt = 1
	}

	//展示留言列表
	skip := (pageInt - 1) * limit
	list, err := feedbackService.List(skip, limit)
	if err != nil {
		showError(c, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "留言板",
		"list":      list,
		"prevPage":  pageInt - 1,
		"nextPage":  pageInt + 1,
		"page":      pageInt,
		"totalPage": totalPageInt,
	})
}
```

其中index.html通过go template语法对内容进行渲染

```html
    <div class="list-group list-group-flush">
        {{range .list}}
        <div class="list-group-item">
            <div><span><b>{{.Username}} 在 {{.CreateTime}} 说：</b></span></div>
            <div><p>{{.Content}}</p></div>
        </div>
        {{end}}
    </div>
    <div>
        <span>第{{.page}}页</span>
        {{if gt .page 1}}<a href="/?page={{.prevPage}}">上一页</a>{{end}}
       {{if lt .page .totalPage}} <a href="/?page={{.nextPage}}">下一页</a>{{end}}
    </div>
```

这样就实现了对文本内容的保存

