package message

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	templateSendURL = "https://api.weixin.qq.com/cgi-bin/message/template/send"
	templateListURL = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	templateAddURL  = "https://api.weixin.qq.com/cgi-bin/template/api_add_template"
	templateDelURL  = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
)

//Template 模板消息
type Template struct {
	*context.Context
}

//NewTemplate 实例化
func NewTemplate(context *context.Context) *Template {
	tpl := new(Template)
	tpl.Context = context
	return tpl
}

//TemplateMessage 发送的模板消息内容
type TemplateMessage struct {
	ToUser     string                       `json:"touser"`          // 必须, 接受者OpenID
	TemplateID string                       `json:"template_id"`     // 必须, 模版ID
	URL        string                       `json:"url,omitempty"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color      string                       `json:"color,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       map[string]*TemplateDataItem `json:"data"`            // 必须, 模板数据

	MiniProgram struct {
		AppID    string `json:"appid"`    //所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` //可选,跳转至小程序地址
}

//TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type resTemplateSend struct {
	util.CommonError

	MsgID int64 `json:"msgid"`
}

//ResTemplateDel 模版消息删除结果
type ResTemplateDel struct {
	util.CommonError
}

//Send 发送模板消息
func (tpl *Template) Send(msg *TemplateMessage) (msgID int64, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	var result resTemplateSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	msgID = result.MsgID
	return
}

// TemplateItem 模板消息.
type TemplateItem struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type resTemplateList struct {
	util.CommonError

	TemplateList []*TemplateItem `json:"template_list"`
}
type ResTemplateAdd struct {
	util.CommonError

	TemplateId string `json:"template_id"`
}

//List 获取模板列表
func (tpl *Template) List() (templateList []*TemplateItem, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resTemplateList
	err = util.DecodeWithError(response, &res, "ListTemplate")
	if err != nil {
		return
	}
	templateList = res.TemplateList
	return
}

//Add 获得模板ID
func (tpl *Template) Add(templateShortId string) (templateId string, err error) {
	var accessToken string
	var templateAdd ResTemplateAdd
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateAddURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, struct {
		TemplateIdShort string `json:"template_id_short"`
	}{
		TemplateIdShort: templateShortId,
	})

	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &templateAdd, "ListTemplate")
	if err != nil {
		return
	}
	return templateAdd.TemplateId, nil
}

//Del 删除模版ID
func (tpl *Template) Del(templateId string) (err error) {
	var accessToken string
	var templateDel ResTemplateDel
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateDelURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, struct {
		TemplateId string `json:"template_id"`
	}{
		TemplateId: templateId,
	})

	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &templateDel, "Del")
	if err != nil {
		return
	}
	return
}
