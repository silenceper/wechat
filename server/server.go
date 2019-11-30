package server

import (
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/context"
	"gitee.com/zhimiao/wechat-sdk/message"
	"gitee.com/zhimiao/wechat-sdk/pay"
	"github.com/siddontang/go/log"
	"runtime/debug"
)

// Server struct
type Server struct {
	*context.Context
	debug          bool                                    // 是否调试模式
	openID         string                                  // 用户唯一openid
	messageHandler func(message.MixMessage) *message.Reply // 消息钩子
	payHandler     func(pay.NotifyResult) *message.Reply   // 消息钩子
	requestRaw     []byte                                  // 原始数据
	requestMsg     message.MixMessage                      // 消息类型数据
	requestPayMsg  pay.NotifyResult                        // 支付消息类型数据
	responseType   message.ResponseType                    // 返回类型 string xml json
	responseMsg    interface{}                             // 响应数据
	isSafeMode     bool                                    // 是否是加密模式
	random         []byte
	nonce          string
	timestamp      int64
}

// NewServer init
func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

// Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	echostr, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(echostr)
		return nil
	}
	response, err := srv.handleRequest()
	if err != nil {
		return err
	}
	// debug
	if srv.debug {
		log.Info("request msg = ", string(srv.requestRaw))
	}
	return srv.buildResponse(response)
}

// GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

// SetDebug set debug field
func (srv *Server) SetDebug(debug bool) {
	srv.debug = debug
}

// SetMessageHandler 常规消息钩子
func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *message.Reply) {
	srv.messageHandler = handler
}

// SetPayHandler 支付消息钩子
func (srv *Server) SetPayHandler(handler func(pay.NotifyResult) *message.Reply) {
	srv.payHandler = handler
}

// 组装返回数据
func (srv *Server) buildResponse(reply *message.Reply) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		// do nothing
		return nil
	}
	switch reply.ReplyScene {
	case message.ReplySceneKefu:
		srv.kefu(reply)
	case message.ReplySceneOpen:
		srv.open(reply)
	}
	return
}

// Send 将自定义的消息发送
func (srv *Server) Send() (err error) {
	if srv.debug {
		fmt.Printf("server send => %#v", srv)
	}
	if srv.responseMsg == nil {
		return
	}
	// 检测消息类型
	switch srv.responseType {
	case message.ResponseTypeXML:
		srv.XML(srv.responseMsg)
		return
	case message.ResponseTypeString:
		if v, ok := srv.responseMsg.(string); ok {
			srv.String(v)
		}
		return
	case message.ResponseTypeJSON:

	}
	return
}
