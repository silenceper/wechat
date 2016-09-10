package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/util"
)

//Server struct
type Server struct {
	*context.Context
	isSafeMode bool
	rawXMLMsg  string
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
		return fmt.Errorf("请求校验失败")
	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		return srv.String(echostr)
	}

	srv.handleRequest()

	return nil
}

//Validate 校验请求是否合法
func (srv *Server) Validate() bool {
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

//HandleRequest 处理微信的请求
func (srv *Server) handleRequest() {
	srv.isSafeMode = false
	encryptType := srv.Query("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}

	_, err := srv.getMessage()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

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
		nonce := srv.Query("nonce")
		msgSignature := srv.Query("msg_signature")
		msgSignatureCreate := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureCreate {
			return nil, fmt.Errorf("消息不合法，验证签名失败")
		}

		//解密
		rawXMLMsgBytes, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			return nil, fmt.Errorf("消息解密失败,err=%v", err)
		}
	} else {
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			return nil, fmt.Errorf("从body中解析xml失败,err=%v", err)
		}
	}

	srv.rawXMLMsg = string(rawXMLMsgBytes)
	fmt.Println(srv.rawXMLMsg)
	return nil, nil
}
