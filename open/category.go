package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	getcategoryURL = "https://api.weixin.qq.com/cgi-bin/wxopen/getcategory"
)

// Category 小程序类目
type Category struct {
	util.CommonError
	Categories    []CategoryInfo `json:"categories"`
	Limit         int            `json:"limit"`
	Quota         int            `json:"quota"`
	CategoryLimit int            `json:"category_limit"`
}
type CategoryInfo struct {
	First       int    `json:"first"`
	FirstName   string `json:"first_name"`
	Second      int    `json:"second"`
	SecondName  string `json:"second_name"`
	AuditStatus int    `json:"audit_status"`
	AuditReason string `json:"audit_reason"`
}

// 这个好像只支持通过接口创建的小程序才能调用
// GetCategory 小程序类目
func (m *MiniPrograms) GetCategory() (ret Category, err error) {
	var body []byte
	body, err = m.get(getcategoryURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}
