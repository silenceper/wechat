package server

import (
	"reflect"
	"strings"
)

// Error 错误
type Error string

const (
	SDKValidateSignatureError Error = "签名验证错误"         //-40001
	SDKParseJsonError         Error = "xml/json解析失败"   //-40002
	SDKComputeSignatureError  Error = "sha加密生成签名失败"    //-40003
	SDKIllegalAesKey          Error = "AESKey 非法"      //-40004
	SDKValidateCorpidError    Error = "ReceiveId 校验错误" //-40005
	SDKEncryptAESError        Error = "AES 加密失败"       //-40006
	SDKDecryptAESError        Error = "AES 解密失败"       //-40007
	SDKIllegalBuffer          Error = "解密后得到的buffer非法" //-40008
	SDKEncodeBase64Error      Error = "base64加密失败"     //-40009
	SDKDecodeBase64Error      Error = "base64解密失败"     //-40010
	SDKGenJsonError           Error = "生成xml/json失败"   //-40011
	SDKIllegalProtocolType    Error = "协议类型非法"         //-40012
	SDKUnknownError           Error = "未知错误"
)

// Error 输出错误信息
func (r Error) Error() string {
	return reflect.ValueOf(r).String()
}

// NewSDKErr 初始化SDK实例错误信息
func NewSDKErr(code int64, msgList ...string) Error {
	switch code {
	case 40001:
		return SDKValidateSignatureError
	case 40002:
		return SDKParseJsonError
	case 40003:
		return SDKComputeSignatureError
	case 40004:
		return SDKIllegalAesKey
	case 40005:
		return SDKValidateCorpidError
	case 40006:
		return SDKEncryptAESError
	case 40007:
		return SDKDecryptAESError
	case 40008:
		return SDKIllegalBuffer
	case 40009:
		return SDKEncodeBase64Error
	case 40010:
		return SDKDecodeBase64Error
	case 40011:
		return SDKGenJsonError
	case 40012:
		return SDKIllegalProtocolType
	default:
		//返回未知的自定义错误
		if len(msgList) > 0 {
			return Error(strings.Join(msgList, ","))
		}
		return SDKUnknownError
	}
}
