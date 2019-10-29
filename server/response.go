package server

import (
	"encoding/xml"
	"gitee.com/zhimiao/wechat-sdk/message"
	"gitee.com/zhimiao/wechat-sdk/open"
	"gitee.com/zhimiao/wechat-sdk/util"
	"reflect"
	"strconv"
)

// 开放平台
func (srv *Server) open(reply *message.Reply) (err error) {
	if reply.ResponseType == "" {
		reply.ResponseType = message.ResponseTypeString
	}
	if reply.ReplyScene == "" {
		reply.ReplyScene = message.ReplyTypeOpen
	}
	if reply.MsgData == "" {
		reply.MsgData = open.SUCCESS
	}

	// 验证票据 /10min通知
	if srv.requestMsg.InfoType == message.InfoTypeVerifyTicket {
		open.NewOpen(srv.Context).SetComponentVerifyTicket(srv.requestMsg.ComponentVerifyTicket)
	}
	return nil
}

// 客服消息
func (srv *Server) kefu(reply *message.Reply) (err error) {
	srv.responseType = reply.ResponseType
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

	params[0] = reflect.ValueOf(util.GetCurrTs())
	value.MethodByName("SetCreateTime").Call(params)

	srv.responseMsg = msgData
	if srv.isSafeMode {
		raw, err := xml.Marshal(srv.responseMsg)
		//安全模式下对消息进行加密
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, raw, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return
		}
		//TODO 如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		srv.responseMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
	}
	return
}
