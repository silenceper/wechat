package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	getcategoryURL      = "https://api.weixin.qq.com/cgi-bin/wxopen/getcategory"
	getAuditCategoryURL = "https://api.weixin.qq.com/wxa/get_category"
)

// Category 小程序类目
type Category struct {
	util.CommonError
	Categories    []CategoryInfo `json:"categories"`
	Limit         int            `json:"limit"`
	Quota         int            `json:"quota"`
	CategoryLimit int            `json:"category_limit"`
}

// CategoryInfo 小程序类目列表
type CategoryInfo struct {
	First       int    `json:"first"`
	FirstName   string `json:"first_name"`
	Second      int    `json:"second"`
	SecondName  string `json:"second_name"`
	AuditStatus int    `json:"audit_status"`
	AuditReason string `json:"audit_reason"`
}

// AuditCategory 提审类目列表
type AuditCategory struct {
	util.CommonError
	CategoryList []AuditCategoryInfo `json:"category_list"`
}

// AuditCategoryInfo 审核时的类目信息
type AuditCategoryInfo struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	ThirdClass  string `json:"third_class"`
	FirstID     int    `json:"first_id"`
	SecondID    int    `json:"second_id"`
	ThirdID     int    `json:"third_id"`
}

// GetCategory 小程序类目
// 这个好像只支持通过接口创建的小程序才能调用
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

// GetAuditCategory 获取审核时可填写的类目信息
// 本接口接口可获取已设置的二级类目及用于代码审核的可选三级类目。
func (m *MiniPrograms) GetAuditCategory() (result []AuditCategoryInfo, err error) {
	var body []byte
	body, err = m.get(getAuditCategoryURL, nil)
	if err != nil {
		return
	}
	var ret AuditCategory
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	result = ret.CategoryList
	return
}
