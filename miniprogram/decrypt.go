package miniprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var (
	// ErrAppIDNotMatch appid不匹配
	ErrAppIDNotMatch = errors.New("app id not match")
	// ErrInvalidBlockSize block size不合法
	ErrInvalidBlockSize = errors.New("invalid block size")
	// ErrInvalidPKCS7Data PKCS7数据不合法
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data")
	// ErrInvalidPKCS7Padding 输入padding失败
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// DecryptedData 解密信息
type DecryptedData struct {
	// 用户授权
	OpenID    string `json:"openId,omitempty"`
	UnionID   string `json:"unionId,omitempty"`
	NickName  string `json:"nickName,omitempty"`
	Gender    int    `json:"gender,omitempty"`
	City      string `json:"city,omitempty"`
	Province  string `json:"province,omitempty"`
	Country   string `json:"country,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty"`
	Language  string `json:"language,omitempty"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark,omitempty"`
	// 手机授权
	PhoneNumber     	string 			`json:"phoneNumber,omitempty"`
	PurePhoneNumber 	string 			`json:"purePhoneNumber,omitempty"`
	CountryCode     	string 			`json:"countryCode,omitempty"`
	// 运动步数
	StepInfoList 		[]StepInfo 		`json:"stepInfoList,omitempty"`
}

// StepInfo 用户微信运动步数(单条)
type StepInfo struct {
	Timestamp 	int64 	`json:"timestamp"`
	Step 		int 	`json:"step"`
}
	
// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

// Decrypt 解密数据
func (wxa *MiniProgram) Decrypt(sessionKey, encryptedData, iv string) (*DecryptedData, error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	var userInfo DecryptedData
	err = json.Unmarshal(cipherText, &userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo.Watermark.AppID != "" && userInfo.Watermark.AppID != wxa.AppID {
		return nil, ErrAppIDNotMatch
	}
	return &userInfo, nil
}
