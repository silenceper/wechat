package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/context"
	"gitee.com/zhimiao/wechat-sdk/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// SUCCESS 成功
	SUCCESS string = "success"
)

// Action 操作项
type Action string

const (
	// ActionAdd 添加
	ActionAdd Action = "add" // ActionAdd 添加
	// ActionDelete 删除
	ActionDelete = "delete" // 删除
	// ActionSet 设置
	ActionSet = "set" // 设置
	// ActionGet 获取
	ActionGet = "get" // 获取
	// ActionOpen 开启
	ActionOpen = "open" // 开启
	// ActionClose 关闭
	ActionClose = "close" // 关闭
)

var actionMap = []Action{
	ActionAdd,
	ActionDelete,
	ActionSet,
	ActionGet,
	ActionOpen,
	ActionClose,
}

// CheckAction 校验Action合法性
func CheckAction(action Action) bool {
	for _, v := range actionMap {
		if v == action {
			return true
		}
	}
	return false
}

// Open struct extends context
type Open struct {
	*context.Context
}

// MiniPrograms 代小程序
type MiniPrograms struct {
	Open
	AuthAppID        string
	AuthRefreshToken string
}

// NewOpen 创建开放平台句柄
func NewOpen(ctx *context.Context) *Open {
	open := &Open{Context: ctx}
	return open
}

// NewMiniPrograms 创建开放平台代小程序句柄
func (o *Open) NewMiniPrograms(appid string, refrshToken string) *MiniPrograms {
	if appid == "" || refrshToken == "" {
		return nil
	}
	miniPrograms := &MiniPrograms{
		Open:             *o,
		AuthAppID:        appid,
		AuthRefreshToken: refrshToken,
	}
	return miniPrograms
}

func (o *Open) buildRequest(urlStr string, param map[string]string) (requestURL string, err error) {
	accessToken, err := o.GetComponentAccessToken()
	if err != nil {
		return
	}
	u, err := url.Parse(urlStr)
	qs := u.Query()
	qs.Add("access_token", accessToken)
	if param != nil {
		for k, v := range param {
			qs.Set(k, v)
		}
	}
	u.RawQuery = qs.Encode()
	requestURL = u.String()
	return
}

// fetchData 拉取统计数据
func (o *Open) post(urlStr string, body interface{}) (response []byte, err error) {
	sendURL, err := o.buildRequest(urlStr, nil)
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendURL, body)
	return
}

// fetchData 拉取统计数据
func (o *Open) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendURL, err := o.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendURL)
	return
}

func (m *MiniPrograms) buildRequest(urlStr string, param map[string]string) (requestURL string, err error) {
	accessToken, err := m.GetAuthrAccessToken(m.AuthAppID)
	if err != nil {
		var ret *context.AuthrAccessToken
		ret, err = m.RefreshAuthrToken(m.AuthAppID, m.AuthRefreshToken)
		if err != nil {
			return
		}
		accessToken = ret.AccessToken
	}
	u, err := url.Parse(urlStr)
	qs := u.Query()
	qs.Add("access_token", accessToken)
	if param != nil {
		for k, v := range param {
			qs.Set(k, v)
		}
	}
	u.RawQuery = qs.Encode()
	requestURL = u.String()
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) post(urlStr string, body interface{}) (response []byte, err error) {
	sendURL, err := m.buildRequest(urlStr, nil)
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendURL, body)
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendURL, err := m.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendURL)
	return
}

// getBinary 拉取二进制数据
func (m *MiniPrograms) getBinary(urlStr string, param map[string]string) (ret []byte, err error) {
	sendURL, err := m.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err := http.Get(sendURL)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http get error : uri=%v , statusCode=%v", sendURL, response.StatusCode)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var jsonRet util.CommonError
		err = json.Unmarshal(responseData, &jsonRet)
		if err == nil && jsonRet.ErrCode != 0 {
			err = fmt.Errorf("[%d]: %s", jsonRet.ErrCode, jsonRet.ErrMsg)
			return
		}
	} else if contentType == "image/jpeg" {
		// 返回文件
		ret = responseData
		return
	} else {
		err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
		return
	}
	return
}
