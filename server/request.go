package server

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/message"
	"gitee.com/zhimiao/wechat-sdk/pay"
	"gitee.com/zhimiao/wechat-sdk/util"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"strconv"
)

type chooseModel struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
}

// IsWechatPay 是否是微信支付
func (m *chooseModel) IsPay() bool {
	if m.ReturnCode != "" && m.ReturnMsg != "" && m.MchID != "" {
		return true
	}
	return false
}

// IsMessage 是否是常规消息体
func (m *chooseModel) IsMessage() bool {
	// 当前非支付类型的都是消息
	return !m.IsPay()
}

//HandleRequest 处理微信的请求
func (srv *Server) handleRequest() (reply *message.Reply, err error) {
	srv.requestRaw, err = ioutil.ReadAll(srv.Request.Body)
	if err != nil {
		err = fmt.Errorf("从body中解析xml失败, err=%v", err)
		return
	}
	choose := chooseModel{}
	err = xml.Unmarshal(srv.requestRaw, &choose)
	if err != nil {
		err = fmt.Errorf("无法识别响应数据, data=%s, err=%v", srv.requestRaw, err)
		return
	}
	if choose.IsPay() {
		reply, err = srv.getPay()
	} else {
		reply, err = srv.getMessage()
	}
	return
}

// getPay 解析支付消息结构
func (srv *Server) getPay() (reply *message.Reply, err error) {
	var notifyResult pay.NotifyResult
	err = xml.Unmarshal(srv.requestRaw, &notifyResult)
	if err != nil {
		return
	}
	// 解析结果非正确的，直接跳出
	if notifyResult.ReturnCode != "SUCCESS" {
		log.Info(srv.requestRaw)
		return
	}
	// 含有加密数据
	if notifyResult.ReqInfo != "" {
		var rawXMLMsg, appID, decryptData []byte
		decryptData, err = base64.StdEncoding.DecodeString(notifyResult.ReqInfo)
		if err != nil {
			return nil, err
		}
		key2 := util.MD5(srv.PayKey)
		srv.random, rawXMLMsg, appID, err = util.AESDecryptMsg(decryptData, []byte(key2))
		if err != nil || len(rawXMLMsg) == 0 {
			if srv.debug {
				log.Warn(srv.random, rawXMLMsg, appID, err)
			}
			return
		}
		err = xml.Unmarshal(rawXMLMsg, &notifyResult)
		if err != nil {
			return
		}
	} else if !pay.VerifySign(srv.PayKey, notifyResult) {
		log.Warn("验签失败", srv.PayKey, notifyResult)
		return
	}
	return
}

//getMessage 解析常规消息结构
func (srv *Server) getMessage() (reply *message.Reply, err error) {
	// 接收OpenId
	srv.openID = srv.Query("openid")
	// 检测数据是否加密
	srv.isSafeMode = srv.Query("encrypt_type") == "aes"
	// 检测数据签名
	if !srv.debug && srv.Query("signature") == util.Signature(srv.Token, srv.Query("timestamp"), srv.Query("nonce")) {
		err = fmt.Errorf("请求校验失败")
		return
	}
	if srv.isSafeMode {
		var encryptedXMLMsg message.EncryptedXMLMsg
		err = xml.NewDecoder(srv.Request.Body).Decode(&encryptedXMLMsg)
		if err != nil {
			err = fmt.Errorf("从body中解析xml失败,err=%v", err)
			return
		}
		//验证消息签名
		timestamp := srv.Query("timestamp")
		srv.timestamp, err = strconv.ParseInt(timestamp, 10, 32)
		if err != nil {
			return
		}
		nonce := srv.Query("nonce")
		srv.nonce = nonce
		msgSignature := srv.Query("msg_signature")
		msgSignatureGen := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureGen {
			err = fmt.Errorf("消息不合法，验证签名失败")
			return
		}
		//解密
		srv.random, srv.requestRaw, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			err = fmt.Errorf("消息解密失败, err=%v", err)
			return
		}
	}
	err = xml.Unmarshal(srv.requestRaw, &srv.requestMsg)
	reply = srv.messageHandler(srv.requestMsg)
	return
}
