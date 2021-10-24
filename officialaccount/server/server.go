package server

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"
)

// Server interface
type Server interface {
	SetRequestAndResponse(req *http.Request, writer http.ResponseWriter)
	SkipValidate(skip bool)                                             // SkipValidate set skip validate
	Serve() error                                                       // Serve 处理微信的请求消息
	Validate() bool                                                     // Validate 校验请求是否合法
	GetOpenID() string                                                  // GetOpenID return openID
	SetMessageHandler(handler func(*message.MixMessage) *message.Reply) // SetMessageHandler 设置用户自定义的回调方法
	Send() (err error)                                                  // Send 将自定义的消息发送

	Render(bytes []byte)                // Render render from bytes
	String(str string)                  // String render from string
	XML(obj interface{})                // XML render to xml
	Query(key string) string            // Query returns the keyed url query value if it exists
	GetQuery(key string) (string, bool) // GetQuery is like Query(), it returns the keyed url query value
	Close()                             // Close or Recycle this server
}
