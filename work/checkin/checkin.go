package checkin

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getCheckinDataURL 获取打卡记录数据
	getCheckinDataURL = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=%s"
)

type (
	// GetCheckinDataRequest 获取打卡记录数据请求
	GetCheckinDataRequest struct {
		Opencheckindatatype int64    `json:"opencheckindatatype"`
		Starttime           int64    `json:"starttime"`
		Endtime             int64    `json:"endtime"`
		Useridlist          []string `json:"useridlist"`
	}
	// GetCheckinDataResponse 获取打卡记录数据响应
	GetCheckinDataResponse struct {
		util.CommonError
		CheckinData []*CheckinData `json:"checkindata"`
	}
	// CheckinData 打卡记录数据
	CheckinData struct {
		Userid         string   `json:"userid"`
		Groupname      string   `json:"groupname"`
		CheckinType    string   `json:"checkin_type"`
		ExceptionType  string   `json:"exception_type"`
		CheckinTime    int64    `json:"checkin_time"`
		LocationTitle  string   `json:"location_title"`
		LocationDetail string   `json:"location_detail"`
		Wifiname       string   `json:"wifiname"`
		Notes          string   `json:"notes"`
		Wifimac        string   `json:"wifimac"`
		Mediaids       []string `json:"mediaids"`
		SchCheckinTime int64    `json:"sch_checkin_time"`
		Groupid        int64    `json:"groupid"`
		ScheduleId     int64    `json:"schedule_id"`
		TimelineId     int64    `json:"timeline_id"`
		Lat            int64    `json:"lat,omitempty"`
		Lng            int64    `json:"lng,omitempty"`
		Deviceid       string   `json:"deviceid,omitempty"`
	}
)

// GetCheckinData 获取打卡记录数据
// @see https://developer.work.weixin.qq.com/document/path/90262
func (r *Client) GetCheckinData(req *GetCheckinDataRequest) (*GetCheckinDataResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(getCheckinDataURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetCheckinDataResponse{}
	if err = util.DecodeWithError(response, result, "GetCheckinData"); err != nil {
		return nil, err
	}
	return result, nil
}
