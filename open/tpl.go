package open

import (
	"encoding/json"
	"fmt"
)

const (
	// 草稿箱列表
	TEMPLATE_DRAFT_LIST_URL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?"
	// 添加草稿到模板
	DRAFT_ADD_TO_TEMPLATE_URL = "https://api.weixin.qq.com/wxa/addtotemplate?"
	// 获取模板列表
	TEMPLATE_LIST_URL = "https://api.weixin.qq.com/wxa/gettemplatelist?"
	// 删除模板
	DELETE_TEMPLATE_URL = "https://api.weixin.qq.com/wxa/deletetemplate?"
)

type TplDetail struct {
	CreateTime             int64  `json:"create_time"`              // 1571730935
	UserVersion            string `json:"user_version"`             // "1.1.3"
	UserDesc               string `json:"user_desc"`                // "小子(LT) 在 2019年10月22日下午3点55分 提交上传"
	DraftId                int    `json:"draft_id"`                 // 145
	SourceMiniprogramAppid string `json:"source_miniprogram_appid"` // "wx37625cd1bad0aaa0"
	SourceMiniprogram      string `json:"source_miniprogram"`       // "易零售"
	Developer              string `json:"developer"`                // "小子(LT)"
}

type TplResponse struct {
	common
	DraftList []TplDetail `json:"draft_list"`
}

func (o *Open) TplDraftList() (ret TplResponse, err error) {
	var body []byte
	body, err = o.get(TEMPLATE_DRAFT_LIST_URL, map[string]string{})
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
