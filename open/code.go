package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	// GetCodePageURL 获取已上传的代码的页面列表
	GetCodePageURL = "https://api.weixin.qq.com/wxa/get_page"
)

// CodePageList 已上传的代码的页面列表
type CodePageList struct {
	util.CommonError
	PageList []string `json:"page_list"`
}

// GetCodePage 获取已上传的代码的页面列表
func (m *MiniPrograms) GetCodePage() (ret CodePageList, err error) {
	var body []byte
	body, err = m.get(GetCodePageURL, nil)
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
