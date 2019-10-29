package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"strconv"

	"gitee.com/zhimiao/wechat-sdk/context"
	"gitee.com/zhimiao/wechat-sdk/message"
	"gitee.com/zhimiao/wechat-sdk/util"
)

//Server struct
type Server struct {
	*context.Context

	debug bool

	openID string

	messageHandler func(message.MixMessage) *message.Reply

	requestRawXMLMsg  []byte
	requestMsg        message.MixMessage

	responseType message.ResponseType // 返回类型 string xml json
	responseMsg interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
}

//NewServer init
func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

// SetDebug set debug field
func (srv *Server) SetDebug(debug bool) {
	srv.debug = debug
}

//Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	if !srv.Validate() {
		return fmt.Errorf("请求校验失败")
	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(echostr)
		return nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		return err
	}

	//debug
	if srv.debug {
		fmt.Println("request msg = ", string(srv.requestRawXMLMsg))
	}

	return srv.buildResponse(response)
}

//Validate 校验请求是否合法
func (srv *Server) Validate() bool {
	if srv.debug {
		return true
	}
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

//HandleRequest 处理微信的请求
func (srv *Server) handleRequest() (reply *message.Reply, err error) {
	//set isSafeMode
	srv.isSafeMode = false
	encryptType := srv.Query("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}

	//set openID
	srv.openID = srv.Query("openid")

	var msg interface{}
	msg, err = srv.getMessage()
	if err != nil {
		return
	}
	mixMessage, success := msg.(message.MixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	srv.requestMsg = mixMessage
	reply = srv.messageHandler(mixMessage)
	return
}

//GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

//getMessage 解析微信返回的消息
func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	var err error
	if srv.isSafeMode {
		var encryptedXMLMsg message.EncryptedXMLMsg
		if err := xml.NewDecoder(srv.Request.Body).Decode(&encryptedXMLMsg); err != nil {
			return nil, fmt.Errorf("从body中解析xml失败,err=%v", err)
		}

		//验证消息签名
		timestamp := srv.Query("timestamp")
		srv.timestamp, err = strconv.ParseInt(timestamp, 10, 32)
		if err != nil {
			return nil, err
		}
		nonce := srv.Query("nonce")
		srv.nonce = nonce
		msgSignature := srv.Query("msg_signature")
		msgSignatureGen := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureGen {
			return nil, fmt.Errorf("消息不合法，验证签名失败")
		}

		//解密
		srv.random, rawXMLMsgBytes, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			return nil, fmt.Errorf("消息解密失败, err=%v", err)
		}
	} else {
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			return nil, fmt.Errorf("从body中解析xml失败, err=%v", err)
		}
	}

	srv.requestRawXMLMsg = rawXMLMsgBytes

	return srv.parseRequestMessage(rawXMLMsgBytes)
}

func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg message.MixMessage, err error) {
	msg = message.MixMessage{}
	err = xml.Unmarshal(rawXMLMsgBytes, &msg)
	return
}

//SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *message.Reply) {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(reply *message.Reply) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		//do nothing
		return nil
	}
	switch reply.ReplyScene {
	case message.ReplyTypeKefu:
		srv.kefu(reply)
	case message.ReplyTypeOpen:
		srv.open(reply)
	}
	return
}

//Send 将自定义的消息发送
func (srv *Server) Send() (err error) {
	if srv.responseMsg == nil {
		return
	}
	// 检测消息类型
	switch srv.responseType {
	case message.ResponseTypeXml:
		srv.XML(srv.responseMsg)
		return
	case message.ResponseTypeString:
		if v, ok := srv.responseMsg.(string); ok {
			srv.String(v)
		}
		return
	case message.ResponseTypeJson:

	}
	return
}