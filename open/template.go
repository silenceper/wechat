package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
	"strconv"
)

const (
	// TemplateDraftListURL 草稿箱列表
	TemplateDraftListURL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?"
	// AddDraftToTemplateURL 添加草稿到模板
	AddDraftToTemplateURL = "https://api.weixin.qq.com/wxa/addtotemplate?"
	// TemplateListURL 获取模板列表
	TemplateListURL = "https://api.weixin.qq.com/wxa/gettemplatelist?"
	// DeleteTemplateURL 删除模板
	DeleteTemplateURL = "https://api.weixin.qq.com/wxa/deletetemplate?"
)

// TplDetail 模板详情
type TplDetail struct {
	CreateTime             int64  `json:"create_time"`              // 1571730935
	UserVersion            string `json:"user_version"`             // "1.1.3"
	UserDesc               string `json:"user_desc"`                // "小子(LT) 在 2019年10月22日下午3点55分 提交上传"
	DraftID                int    `json:"draft_id"`                 // 145
	TemplateID             int    `json:"template_id"`              // 145
	SourceMiniprogramAppid string `json:"source_miniprogram_appid"` // "wx37625cd1b23423432"
	SourceMiniprogram      string `json:"source_miniprogram"`       // "纸喵软件"
	Developer              string `json:"developer"`                // "小子(LT)"
}

// TplResponse 模板返回体
type TplResponse struct {
	util.CommonError
	DraftList    []TplDetail `json:"draft_list"`
	TemplateList []TplDetail `json:"template_list"`
}

// DeleteTpl 删除模板
func (o *Open) DeleteTpl(TemplateID int) (err error) {
	body, err := o.post(DeleteTemplateURL, map[string]string{
		"template_id": strconv.Itoa(TemplateID),
	})
	if err != nil {
		return
	}
	ret := &TplResponse{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// TplList 获取模板列表
func (o *Open) TplList() (ret TplResponse, err error) {
	var body []byte
	body, err = o.get(TemplateListURL, nil)
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

// AddDrafToTpl 添加草稿到模板
func (o *Open) AddDrafToTpl(draftID int) (err error) {
	body, err := o.post(AddDraftToTemplateURL, map[string]string{
		"draft_id": strconv.Itoa(draftID),
	})
	if err != nil {
		return
	}
	ret := &TplResponse{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// TplDraftList 草稿列表
func (o *Open) TplDraftList() (ret TplResponse, err error) {
	var body []byte
	body, err = o.get(TemplateDraftListURL, nil)
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
