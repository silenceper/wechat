package encryptor

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/miniprogram/context"
)

// WeRun 微信运动
type WeRun struct {
	*context.Context
}

// WeRunData 微信运动数据
type WeRunData struct {
	StepInfoList []struct {
		Timestamp int `json:"timestamp"`
		Step      int `json:"step"`
	} `json:"stepInfoList"`
}

// NewWeRun 实例化
func NewWeRun(ctx *context.Context) *WeRun {
	return &WeRun{Context: ctx}
}

// GetWeRunData 解密数据
func (werun *WeRun) GetWeRunData(sessionKey, encryptedData, iv string) (*WeRunData, error) {
	cipherText, err := getCipherText(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var weRunData WeRunData
	err = json.Unmarshal(cipherText, &weRunData)
	if err != nil {
		return nil, err
	}
	return &weRunData, nil
}
