package context

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/minishop/config"
	"github.com/silenceper/wechat/v2/util"
)

// Context struct
type Context struct {
	Config            *config.Config
	AccessTokenHandle credential.AccessTokenHandle
}

// FetchData 数据请求
func (c *Context) FetchData(urlStr string, body interface{}) (response []byte, err error) {
	accessToken, err := c.AccessTokenHandle.GetAccessToken()
	if err != nil {
		return nil, err
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)

	v := url.Values{}

	if c.Config.ServiceID != "" {
		v.Add("service_id", c.Config.ServiceID)
	}
	if c.Config.SpecificationID != "" {
		v.Add("specification_id", c.Config.SpecificationID)
	}
	encode := v.Encode()
	if encode != "" {
		urlStr = urlStr + "&" + encode
	}
	response, err = util.PostJSON(urlStr, body)
	if err != nil {
		return
	}
	// 返回错误信息
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err == nil && result.ErrCode != 0 {
		err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return response, err
}
