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

// UserInfo 用户信息
type UserInfo struct {
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

// 用户手机号
type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
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

// get cipherText
func getCipherText(sessionKey, encryptedData, iv string) ([]byte, error) {
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
	return cipherText, err
}

// Decrypt 解密(用户)数据
func (wxa *MiniProgram) Decrypt(sessionKey, encryptedData, iv string) (*UserInfo, error) {
	// 拿到 cipherText
	cipherText,err := getCipherText(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var userInfo UserInfo
	err = json.Unmarshal(cipherText, &userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo.Watermark.AppID != wxa.AppID {
		return nil, ErrAppIDNotMatch
	}
	return &userInfo, nil
}

// Decrypt 解密(手机)数据
func (wxa *MiniProgram) DecryptPhone(sessionKey, encryptedData, iv string) (*PhoneInfo, error) {
	// 拿到 cipherText
	cipherText,err := getCipherText(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var phoneInfo PhoneInfo
	err = json.Unmarshal(cipherText, &phoneInfo)
	if err != nil {
		return nil, err
	}
	if phoneInfo.Watermark.AppID != wxa.AppID {
		return nil, ErrAppIDNotMatch
	}
	return &phoneInfo, nil
}
