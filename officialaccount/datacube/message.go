package datacube

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	getUpstreamMsg          = "https://api.weixin.qq.com/datacube/getupstreammsg"
	getUpstreamMsgHour      = "https://api.weixin.qq.com/datacube/getupstreammsghour"
	getUpstreamMsgWeek      = "https://api.weixin.qq.com/datacube/getupstreammsgweek"
	getUpstreamMsgMonth     = "https://api.weixin.qq.com/datacube/getupstreammsgmonth"
	getUpstreamMsgDist      = "https://api.weixin.qq.com/datacube/getupstreammsgdist"
	getUpstreamMsgDistWeek  = "https://api.weixin.qq.com/datacube/getupstreammsgdistweek"
	getUpstreamMsgDistMonth = "https://api.weixin.qq.com/datacube/getupstreammsgdistmonth"
)

//ResUpstreamMsg 获取消息发送概况数据响应
type ResUpstreamMsg struct {
	util.CommonError

	List []struct {
		RefDate  string `json:"ref_date"`
		MsgType  int    `json:"msg_type"`
		MsgUser  int    `json:"msg_user"`
		MsgCount int    `json:"msg_count"`
	} `json:"list"`
}

//ResUpstreamMsgHour 获取消息分送分时数据响应
type ResUpstreamMsgHour struct {
	util.CommonError

	List []struct {
		RefDate  string `json:"ref_date"`
		RefHour  int    `json:"ref_hour"`
		MsgType  int    `json:"msg_type"`
		MsgUser  int    `json:"msg_user"`
		MsgCount int    `json:"msg_count"`
	} `json:"list"`
}

//ResUpstreamMsgWeek 获取消息发送周数据响应
type ResUpstreamMsgWeek struct {
	util.CommonError

	List []struct {
		RefDate  string `json:"ref_date"`
		MsgType  int    `json:"msg_type"`
		MsgUser  int    `json:"msg_user"`
		MsgCount int    `json:"msg_count"`
	} `json:"list"`
}

//ResUpstreamMsgMonth 获取消息发送月数据响应
type ResUpstreamMsgMonth struct {
	util.CommonError

	List []struct {
		RefDate  string `json:"ref_date"`
		MsgType  int    `json:"msg_type"`
		MsgUser  int    `json:"msg_user"`
		MsgCount int    `json:"msg_count"`
	} `json:"list"`
}

//ResUpstreamMsgDist 获取消息发送分布数据响应
type ResUpstreamMsgDist struct {
	util.CommonError

	List []struct {
		RefDate       string `json:"ref_date"`
		CountInterval int    `json:"count_interval"`
		MsgUser       int    `json:"msg_user"`
	} `json:"list"`
}

//ResUpstreamMsgDistWeek 获取消息发送分布周数据响应
type ResUpstreamMsgDistWeek struct {
	util.CommonError

	List []struct {
		RefDate       string `json:"ref_date"`
		CountInterval int    `json:"count_interval"`
		MsgUser       int    `json:"msg_user"`
	} `json:"list"`
}

//ResUpstreamMsgDistMonth 获取消息发送分布月数据响应
type ResUpstreamMsgDistMonth struct {
	util.CommonError

	List []struct {
		RefDate       string `json:"ref_date"`
		CountInterval int    `json:"count_interval"`
		MsgUser       int    `json:"msg_user"`
	} `json:"list"`
}

//GetUpstreamMsg 获取消息发送概况数据
func (cube *DataCube) GetUpstreamMsg(s string, e string) (resUpstreamMsg ResUpstreamMsg, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsg, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsg)
	if err != nil {
		return
	}
	if resUpstreamMsg.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsg Error , errcode=%d , errmsg=%s", resUpstreamMsg.ErrCode, resUpstreamMsg.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgHour 获取消息分送分时数据
func (cube *DataCube) GetUpstreamMsgHour(s string, e string) (resUpstreamMsgHour ResUpstreamMsgHour, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgHour, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgHour)
	if err != nil {
		return
	}
	if resUpstreamMsgHour.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgHour Error , errcode=%d , errmsg=%s", resUpstreamMsgHour.ErrCode, resUpstreamMsgHour.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgWeek 获取消息发送周数据
func (cube *DataCube) GetUpstreamMsgWeek(s string, e string) (resUpstreamMsgWeek ResUpstreamMsgWeek, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgWeek, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgWeek)
	if err != nil {
		return
	}
	if resUpstreamMsgWeek.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgWeek Error , errcode=%d , errmsg=%s", resUpstreamMsgWeek.ErrCode, resUpstreamMsgWeek.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgMonth 获取消息发送月数据
func (cube *DataCube) GetUpstreamMsgMonth(s string, e string) (resUpstreamMsgMonth ResUpstreamMsgMonth, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgMonth, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgMonth)
	if err != nil {
		return
	}
	if resUpstreamMsgMonth.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgMonth Error , errcode=%d , errmsg=%s", resUpstreamMsgMonth.ErrCode, resUpstreamMsgMonth.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgDist 获取消息发送分布数据
func (cube *DataCube) GetUpstreamMsgDist(s string, e string) (resUpstreamMsgDist ResUpstreamMsgDist, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgDist, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgDist)
	if err != nil {
		return
	}
	if resUpstreamMsgDist.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgDist Error , errcode=%d , errmsg=%s", resUpstreamMsgDist.ErrCode, resUpstreamMsgDist.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgDistWeek 获取消息发送分布周数据
func (cube *DataCube) GetUpstreamMsgDistWeek(s string, e string) (resUpstreamMsgDistWeek ResUpstreamMsgDistWeek, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgDistWeek, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgDistWeek)
	if err != nil {
		return
	}
	if resUpstreamMsgDistWeek.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgDistWeek Error , errcode=%d , errmsg=%s", resUpstreamMsgDistWeek.ErrCode, resUpstreamMsgDistWeek.ErrMsg)
		return
	}
	return
}

//GetUpstreamMsgDistMonth 获取消息发送分布月数据
func (cube *DataCube) GetUpstreamMsgDistMonth(s string, e string) (resUpstreamMsgDistMonth ResUpstreamMsgDistMonth, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUpstreamMsgDistMonth, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resUpstreamMsgDistMonth)
	if err != nil {
		return
	}
	if resUpstreamMsgDistMonth.ErrCode != 0 {
		err = fmt.Errorf("GetUpstreamMsgDistMonth Error , errcode=%d , errmsg=%s", resUpstreamMsgDistMonth.ErrCode, resUpstreamMsgDistMonth.ErrMsg)
		return
	}
	return
}
