package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime/debug"
	"strconv"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/util"
)

//Server struct
type Server struct {
	*context.Context

	openID string

	messageHandler func(message.MixMessage) *message.Reply

	requestRawXMLMsg  []byte
	requestMsg        message.MixMessage
	responseRawXMLMsg []byte
	responseMsg       interface{}

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

//Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	if !srv.Validate() {
		//return fmt.Errorf("请求校验失败")
	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(echostr)
		return nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		fmt.Println("handleRequest err:", err.Error())
		return err
	}

	//debug
	//fmt.Println("request msg = ", string(srv.requestRawXMLMsg))

	return srv.buildResponse(response)
}

//Validate 校验请求是否合法
func (srv *Server) Validate() bool {
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
	msgType := reply.MsgType
	switch msgType {
	case message.MsgTypeText:
	case message.MsgTypeImage:
	case message.MsgTypeVoice:
	case message.MsgTypeVideo:
	case message.MsgTypeMusic:
	case message.MsgTypeNews:
	case message.MsgTypeTransfer:
	default:
		err = message.ErrUnsupportReply
		return
	}

	msgData := reply.MsgData
	value := reflect.ValueOf(msgData)
	//msgData must be a ptr
	kind := value.Kind().String()
	if "ptr" != kind {
		return message.ErrUnsupportReply
	}

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.requestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.requestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(msgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurrTs())
	value.MethodByName("SetCreateTime").Call(params)

	srv.responseMsg = msgData
	srv.responseRawXMLMsg, err = xml.Marshal(msgData)
	return
}

//Send 将自定义的消息发送
func (srv *Server) Send() (err error) {
	replyMsg := srv.responseMsg
	if srv.isSafeMode {
		//安全模式下对消息进行加密
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, srv.responseRawXMLMsg, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return
		}
		//TODO 如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		replyMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
	}
	if replyMsg != nil {
		srv.XML(replyMsg)
	}
	return
}
