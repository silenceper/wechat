package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/miniprogram/message"
	"github.com/silenceper/wechat/v2/util"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
)

// Server struct
type Server struct {
	*context.Context
	Write        http.ResponseWriter
	Request      *http.Request
	skipValidate bool
	openID       string

	messageHandler func(mixMessage *message.MiniProgramMixMessage) *message.Reply

	RequestRawXMLMsg []byte
	RequestMsg       *message.MiniProgramMixMessage

	ResponseRawXMLMsg []byte
	ResponseMsg       interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
}

func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

func (srv *Server) Server() error {
	if !srv.Validate() {
		return fmt.Errorf("请求签名校验失败")
	}
	echoStr := srv.Query("echostr")
	if echoStr != "" {
		srv.SetResponseWrite(echoStr)
		return nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		return err
	}

	return srv.buildResponse(response)

}

// SkipValidate 设置跳过签名校验
func (srv *Server) SkipValidate(skip bool) {
	srv.skipValidate = skip
}

// Validate 校验请求是否合法
func (srv *Server) Validate() bool {
	if srv.skipValidate {
		return true
	}
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

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
	mixMessage, success := msg.(*message.MiniProgramMixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	srv.RequestMsg = mixMessage
	reply = srv.messageHandler(mixMessage)
	return
}

//GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	var err error
	if srv.isSafeMode {
		var encryptedXMLMsg message.EncryptedXMLMsg
		if err := xml.NewDecoder(srv.Request.Body).Decode(&encryptedXMLMsg); err != nil {
			return nil, fmt.Errorf("从body中解析xml失败，err=%v", err)
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

		// 解密
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
	srv.RequestRawXMLMsg = rawXMLMsgBytes
	return srv.parseRequestMessage(rawXMLMsgBytes)
}

func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg *message.MiniProgramMixMessage, err error) {
	msg = &message.MiniProgramMixMessage{}
	err = xml.Unmarshal(rawXMLMsgBytes, msg)
	return
}

func (srv *Server) SetMessageHandler(handler func(*message.MiniProgramMixMessage) *message.Reply) {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(reply *message.Reply) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		return nil
	}
	msgType := reply.MsgType
	switch msgType {
	case message.MsgTypeEvent:
	case message.MsgTypeImage:
	case message.MsgTypeLink:
	case message.MsgTypeText:
	case message.MsgTypeMiniProgramPage:
	default:
		err = message.ErrUnsupportedReply
		return
	}
	msgData := reply.MsgData
	value := reflect.ValueOf(msgData)
	//msgData must be a ptr
	kind := value.Kind().String()
	if kind != "ptr" {
		return message.ErrUnsupportedReply
	}
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.RequestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.RequestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(srv.RequestMsg.MsgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurrTS())
	value.MethodByName("SetCreateTime").Call(params)

	srv.ResponseMsg = msgData
	srv.ResponseRawXMLMsg, err = xml.Marshal(msgData)
	return
}
