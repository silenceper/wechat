package message

import (
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	setIndustryURL = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"
)

//Industry 行业信息
type Industry struct {
	*context.Context
}

//NewIndustry 实例化
func NewIndustry(context *context.Context) *Industry {
	industry := new(Industry)
	industry.Context = context
	return industry
}

type SetIndustryResult struct {
	util.CommonError
}

//SetIndustry 设置行业
func (industry *Industry) SetIndustry(industryId1, industryId2 int64) (result *SetIndustryResult, err error) {
	var accessToken string
	var data = struct {
		IndustryId1 int64 `json:"industry_id1"`
		IndustryId2 int64 `json:"industry_id2"`
	}{IndustryId1: industryId1, IndustryId2: industryId2}
	accessToken, err = industry.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(setIndustryURL, accessToken)
	result = &SetIndustryResult{}
	response, err := util.PostJSON(uri, data)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, result, "SetIndustry")

	//msgID = result.MsgID
	return
}
