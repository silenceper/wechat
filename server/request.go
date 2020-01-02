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
	if m.ReturnCode != "" && m.MchID != "" {
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
	err = xml.Unmarshal(srv.requestRaw, &srv.requestPayMsg)
	if err != nil {
		return
	}
	// 解析结果非正确的，直接跳出
	if srv.requestPayMsg.ReturnCode != "SUCCESS" {
		log.Info(srv.requestRaw)
		return
	}
	// 含有加密数据
	if srv.requestPayMsg.ReqInfo != "" {
		var rawXMLMsg, encryptData []byte
		key2 := util.MD5(srv.PayKey)
		encryptData, err = base64.StdEncoding.DecodeString(srv.requestPayMsg.ReqInfo)
		if err != nil {
			if srv.debug {
				log.Warn("返回数据无法识别", srv.requestPayMsg)
			}
			return
		}
		rawXMLMsg, err = util.ECBDecrypt(encryptData, []byte(key2))
		if err != nil || len(rawXMLMsg) == 0 {
			if srv.debug {
				log.Warn(srv.random, rawXMLMsg, err)
			}
			return
		}
		err = xml.Unmarshal(rawXMLMsg, &srv.requestPayMsg)
		if err != nil {
			return
		}

	} else if !pay.VerifySign(srv.PayKey, srv.requestPayMsg) {
		log.Warn("验签失败", srv.PayKey, srv.requestPayMsg)
		return
	}
	// 判断支付返回类型
	if srv.requestPayMsg.RefundFee > 0 {
		srv.requestPayMsg.PayNotifyInfo = pay.PayTypeRefund
	} else if srv.requestPayMsg.TotalFee > 0 {
		srv.requestPayMsg.PayNotifyInfo = pay.PayTypePay
	}
	reply = srv.payHandler(srv.requestPayMsg)
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
		err = xml.Unmarshal(srv.requestRaw, &encryptedXMLMsg)
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
