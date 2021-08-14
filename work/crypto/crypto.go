package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	// ValidateSignatureError 无效的签名
	ValidateSignatureError int = -40001
	// ParseXMLError 解析XML失败
	ParseXMLError int = -40002
	// ValidateCorpIDError 校验CorpID错误
	ValidateCorpIDError int = -40005
	// EncryptAESError 加密失败
	EncryptAESError int = -40006
	// DecryptAESError 解密失败
	DecryptAESError int = -40007
	// IllegalBuffer 无效的Buffer
	IllegalBuffer int = -40008
	// DecodeBase64Error Base64解码失败
	DecodeBase64Error int = -40010
	// GenXMLError 生成XML失败
	GenXMLError int = -40010
)

// ProtocolType 编码类型
type ProtocolType int

const (
	// XMLType XML类型
	XMLType ProtocolType = 1
)

// CryptError 错误
type CryptError struct {
	ErrCode int
	ErrMsg  string
}

// NewCryptError 初始化新的错误
func NewCryptError(errCode int, errMsg string) *CryptError {
	return &CryptError{ErrCode: errCode, ErrMsg: errMsg}
}

// WXBizMsg4Recv 微信接受消息
type WXBizMsg4Recv struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
	AgentID    string `xml:"AgentID"`
}

// CDATA 数据
type CDATA struct {
	Value string `xml:",cdata"`
}

// WXBizMsg4Send 发送微信消息
type WXBizMsg4Send struct {
	XMLName   xml.Name `xml:"xml"`
	Encrypt   CDATA    `xml:"Encrypt"`
	Signature CDATA    `xml:"MsgSignature"`
	Timestamp string   `xml:"TimeStamp"`
	Nonce     CDATA    `xml:"Nonce"`
}

// NewWXBizMsg4Send 初始化新的消息
func NewWXBizMsg4Send(encrypt, signature, timestamp, nonce string) *WXBizMsg4Send {
	return &WXBizMsg4Send{Encrypt: CDATA{Value: encrypt}, Signature: CDATA{Value: signature}, Timestamp: timestamp, Nonce: CDATA{Value: nonce}}
}

// ProtocolProcessor 编码处理
type ProtocolProcessor interface {
	parse(srcData []byte) (*WXBizMsg4Recv, *CryptError)
	serialize(msgSend *WXBizMsg4Send) ([]byte, *CryptError)
}

// WXBizMsgCrypt 消息加密
type WXBizMsgCrypt struct {
	token             string
	encodingAesKey    string
	receiverID        string
	protocolProcessor ProtocolProcessor
}

// XMLProcessor XML处理
type XMLProcessor struct {
}

func (r *XMLProcessor) parse(srcData []byte) (*WXBizMsg4Recv, *CryptError) {
	var msg4Recv WXBizMsg4Recv
	err := xml.Unmarshal(srcData, &msg4Recv)
	if nil != err {
		return nil, NewCryptError(ParseXMLError, "xml to msg fail")
	}
	return &msg4Recv, nil
}

func (r *XMLProcessor) serialize(msg4Send *WXBizMsg4Send) ([]byte, *CryptError) {
	xmlMsg, err := xml.Marshal(msg4Send)
	if nil != err {
		return nil, NewCryptError(GenXMLError, err.Error())
	}
	return xmlMsg, nil
}

// NewWXBizMsgCrypt 初始化
func NewWXBizMsgCrypt(token, encodingAesKey, receiverID string, protocolType ProtocolType) *WXBizMsgCrypt {
	var protocolProcessor ProtocolProcessor
	if protocolType != XMLType {
		panic("unSupport protocol")
	} else {
		protocolProcessor = new(XMLProcessor)
	}
	return &WXBizMsgCrypt{token: token, encodingAesKey: encodingAesKey + "=", receiverID: receiverID, protocolProcessor: protocolProcessor}
}

func (r *WXBizMsgCrypt) randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (r *WXBizMsgCrypt) pKCS7Padding(plaintext string, blockSize int) []byte {
	padding := blockSize - (len(plaintext) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	var buffer bytes.Buffer
	buffer.WriteString(plaintext)
	buffer.Write(padText)
	return buffer.Bytes()
}

func (r *WXBizMsgCrypt) pKCS7UnPadding(plaintext []byte, blockSize int) ([]byte, *CryptError) {
	plaintextLen := len(plaintext)
	if nil == plaintext || plaintextLen == 0 {
		return nil, NewCryptError(DecryptAESError, "pKCS7UnPadding error nil or zero")
	}
	if plaintextLen%blockSize != 0 {
		return nil, NewCryptError(DecryptAESError, "pKCS7UnPadding text not a multiple of the block size")
	}
	paddingLen := int(plaintext[plaintextLen-1])
	return plaintext[:plaintextLen-paddingLen], nil
}

func (r *WXBizMsgCrypt) cbcEncryptor(plaintext string) ([]byte, *CryptError) {
	aesKey, err := base64.StdEncoding.DecodeString(r.encodingAesKey)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}
	const blockSize = 32
	padMsg := r.pKCS7Padding(plaintext, blockSize)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, NewCryptError(EncryptAESError, err.Error())
	}

	ciphertext := make([]byte, len(padMsg))
	iv := aesKey[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext, padMsg)
	base64Msg := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(base64Msg, ciphertext)

	return base64Msg, nil
}

func (r *WXBizMsgCrypt) cbcDecipher(base64EncryptMsg string) ([]byte, *CryptError) {
	aesKey, err := base64.StdEncoding.DecodeString(r.encodingAesKey)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}

	encryptMsg, err := base64.StdEncoding.DecodeString(base64EncryptMsg)
	if nil != err {
		return nil, NewCryptError(DecodeBase64Error, err.Error())
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, NewCryptError(DecryptAESError, err.Error())
	}

	if len(encryptMsg) < aes.BlockSize {
		return nil, NewCryptError(DecryptAESError, "encrypt_msg size is not valid")
	}

	iv := aesKey[:aes.BlockSize]

	if len(encryptMsg)%aes.BlockSize != 0 {
		return nil, NewCryptError(DecryptAESError, "encrypt_msg not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptMsg, encryptMsg)

	return encryptMsg, nil
}

func (r *WXBizMsgCrypt) calSignature(timestamp, nonce, data string) string {
	sortArr := []string{r.token, timestamp, nonce, data}
	sort.Strings(sortArr)
	var buffer bytes.Buffer
	for _, value := range sortArr {
		buffer.WriteString(value)
	}

	sha := sha1.New()
	sha.Write(buffer.Bytes())
	signature := fmt.Sprintf("%x", sha.Sum(nil))
	return signature
}

// ParsePlainText 解析原始文本内容
func (r *WXBizMsgCrypt) ParsePlainText(plaintext []byte) ([]byte, uint32, []byte, []byte, *CryptError) {
	const blockSize = 32
	plaintext, err := r.pKCS7UnPadding(plaintext, blockSize)
	if nil != err {
		return nil, 0, nil, nil, err
	}

	textLen := uint32(len(plaintext))
	if textLen < 20 {
		return nil, 0, nil, nil, NewCryptError(IllegalBuffer, "plain is to small 1")
	}
	random := plaintext[:16]
	msgLen := binary.BigEndian.Uint32(plaintext[16:20])
	if textLen < (20 + msgLen) {
		return nil, 0, nil, nil, NewCryptError(IllegalBuffer, "plain is to small 2")
	}

	msg := plaintext[20 : 20+msgLen]
	receiverID := plaintext[20+msgLen:]

	return random, msgLen, msg, receiverID, nil
}

// VerifyURL 校验请求参数是否合法
func (r *WXBizMsgCrypt) VerifyURL(msgSignature, timestamp, nonce, echoStr string) ([]byte, *CryptError) {
	signature := r.calSignature(timestamp, nonce, echoStr)

	if strings.Compare(signature, msgSignature) != 0 {
		return nil, NewCryptError(ValidateSignatureError, "signature not equal")
	}

	plaintext, err := r.cbcDecipher(echoStr)
	if nil != err {
		return nil, err
	}

	_, _, msg, receiverID, err := r.ParsePlainText(plaintext)
	if nil != err {
		return nil, err
	}

	if len(r.receiverID) > 0 && strings.Compare(string(receiverID), r.receiverID) != 0 {
		fmt.Println(string(receiverID), r.receiverID, len(receiverID), len(r.receiverID))
		return nil, NewCryptError(ValidateCorpIDError, "receiverId is not eQuil")
	}

	return msg, nil
}

// EncryptMsg 消息加密
func (r *WXBizMsgCrypt) EncryptMsg(replyMsg, timestamp, nonce string) ([]byte, *CryptError) {
	randStr := r.randString(16)
	var buffer bytes.Buffer
	buffer.WriteString(randStr)

	msgLenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLenBuf, uint32(len(replyMsg)))
	buffer.Write(msgLenBuf)
	buffer.WriteString(replyMsg)
	buffer.WriteString(r.receiverID)

	tmpCipherText, err := r.cbcEncryptor(buffer.String())
	if nil != err {
		return nil, err
	}
	ciphertext := string(tmpCipherText)

	signature := r.calSignature(timestamp, nonce, ciphertext)

	msg4Send := NewWXBizMsg4Send(ciphertext, signature, timestamp, nonce)
	return r.protocolProcessor.serialize(msg4Send)
}

// DecryptMsg 消息解密
func (r *WXBizMsgCrypt) DecryptMsg(msgSignature, timestamp, nonce string, postData []byte) ([]byte, *CryptError) {
	msg4Recv, cryptErr := r.protocolProcessor.parse(postData)
	if nil != cryptErr {
		return nil, cryptErr
	}

	signature := r.calSignature(timestamp, nonce, msg4Recv.Encrypt)

	if strings.Compare(signature, msgSignature) != 0 {
		return nil, NewCryptError(ValidateSignatureError, "signature not equal")
	}

	plaintext, cryptErr := r.cbcDecipher(msg4Recv.Encrypt)
	if nil != cryptErr {
		return nil, cryptErr
	}

	_, _, msg, receiverID, cryptErr := r.ParsePlainText(plaintext)
	if nil != cryptErr {
		return nil, cryptErr
	}

	if len(r.receiverID) > 0 && strings.Compare(string(receiverID), r.receiverID) != 0 {
		return nil, NewCryptError(ValidateCorpIDError, "receiver_id is not e_quil")
	}

	return msg, nil
}
