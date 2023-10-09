package checkin

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getCheckinDataURL 获取打卡记录数据
	getCheckinDataURL = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=%s"
	// getDayDataURL 获取打卡日报数据
	getDayDataURL = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata?access_token=%s"
	// getMonthDataURL 获取打卡月报数据
	getMonthDataURL = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_monthdata?access_token=%s"
)

type (
	// GetCheckinDataRequest 获取打卡记录数据请求
	GetCheckinDataRequest struct {
		OpenCheckinDataType int64    `json:"opencheckindatatype"`
		StartTime           int64    `json:"starttime"`
		EndTime             int64    `json:"endtime"`
		UserIDList          []string `json:"useridlist"`
	}
	// GetCheckinDataResponse 获取打卡记录数据响应
	GetCheckinDataResponse struct {
		util.CommonError
		CheckinData []*GetCheckinDataItem `json:"checkindata"`
	}
	// GetCheckinDataItem 打卡记录数据
	GetCheckinDataItem struct {
		UserID         string   `json:"userid"`
		GroupName      string   `json:"groupname"`
		CheckinType    string   `json:"checkin_type"`
		ExceptionType  string   `json:"exception_type"`
		CheckinTime    int64    `json:"checkin_time"`
		LocationTitle  string   `json:"location_title"`
		LocationDetail string   `json:"location_detail"`
		WifiName       string   `json:"wifiname"`
		Notes          string   `json:"notes"`
		WifiMac        string   `json:"wifimac"`
		MediaIDs       []string `json:"mediaids"`
		SchCheckinTime int64    `json:"sch_checkin_time"`
		GroupID        int64    `json:"groupid"`
		ScheduleID     int64    `json:"schedule_id"`
		TimelineID     int64    `json:"timeline_id"`
		Lat            int64    `json:"lat,omitempty"`
		Lng            int64    `json:"lng,omitempty"`
		DeviceID       string   `json:"deviceid,omitempty"`
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

type (
	// GetDayDataResponse 获取打卡日报数据
	GetDayDataResponse struct {
		util.CommonError
		Datas []DayDataItem `json:"datas"`
	}

	DayDataItem struct {
		BaseInfo       DayBaseInfo     `json:"base_info"`
		SummaryInfo    DaySummaryInfo  `json:"summary_info"`
		HolidayInfos   []HolidayInfo   `json:"holiday_infos"`
		ExceptionInfos []ExceptionInfo `json:"exception_infos"`
		OtInfo         OtInfo          `json:"ot_info"`
		SpItems        []SpItem        `json:"sp_items"`
	}

	// DayBaseInfo 基础信息
	DayBaseInfo struct {
		Date        int64       `json:"date"`
		RecordType  int64       `json:"record_type"`
		Name        string      `json:"name"`
		NameEx      string      `json:"name_ex"`
		DepartsName string      `json:"departs_name"`
		AcctId      string      `json:"acctid"`
		DayType     int64       `json:"day_type"`
		RuleInfo    DayRuleInfo `json:"rule_info"`
	}

	// CheckInTime 当日打卡时间
	CheckInTime struct {
		WorkSec    int64 `json:"work_sec"`
		OffWorkSec int64 `json:"off_work_sec"`
	}

	// DayRuleInfo 打卡人员所属规则信息
	DayRuleInfo struct {
		GroupId      int64         `json:"groupid"`
		GroupName    string        `json:"groupname"`
		ScheduleId   int64         `json:"scheduleid"`
		ScheduleName string        `json:"schedulename"`
		CheckInTimes []CheckInTime `json:"checkintime"`
	}

	// DaySummaryInfo 汇总信息
	DaySummaryInfo struct {
		CheckinCount    int64 `json:"checkin_count"`
		RegularWorkSec  int64 `json:"regular_work_sec"`
		StandardWorkSec int64 `json:"standard_work_sec"`
		EarliestTime    int64 `json:"earliest_time"`
		LastestTime     int64 `json:"lastest_time"`
	}

	// HolidayInfo 假勤相关信息
	HolidayInfo struct {
		SpNumber      string        `json:"sp_number"`
		SpTitle       SpTitle       `json:"sp_title"`
		SpDescription SpDescription `json:"sp_description"`
	}

	// SpTitle 假勤信息摘要-标题信息
	SpTitle struct {
		Data []SpData `json:"data"`
	}

	// SpDescription 假勤信息摘要-描述信息
	SpDescription struct {
		Data []SpData `json:"data"`
	}

	// SpData 假勤信息(多种语言描述，目前只有中文一种)
	SpData struct {
		Lang string `json:"lang"`
		Text string `json:"text"`
	}

	// SpItem 假勤统计信息
	SpItem struct {
		Count      int64  `json:"count"`
		Duration   int64  `json:"duration"`
		TimeType   int64  `json:"time_type"`
		Type       int64  `json:"type"`
		VacationId int64  `json:"vacation_id"`
		Name       string `json:"name"`
	}

	// ExceptionInfo 校准状态信息
	ExceptionInfo struct {
		Count     int64 `json:"count"`
		Duration  int64 `json:"duration"`
		Exception int64 `json:"exception"`
	}

	// OtInfo 加班信息
	OtInfo struct {
		OtStatus          int64    `json:"ot_status"`
		OtDuration        int64    `json:"ot_duration"`
		ExceptionDuration []uint64 `json:"exception_duration"`
	}
)

// GetDayData 获取打卡日报数据
// @see https://developer.work.weixin.qq.com/document/path/96498
func (r *Client) GetDayData(req *GetCheckinDataRequest) (result *GetDayDataResponse, err error) {
	var (
		response    []byte
		accessToken string
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return
	}
	if response, err = util.PostJSON(fmt.Sprintf(getDayDataURL, accessToken), req); err != nil {
		return
	}

	result = new(GetDayDataResponse)
	if err = util.DecodeWithError(response, result, "GetDayData"); err != nil {
		return
	}
	return
}

type (
	// GetMonthDataResponse 获取打卡月报数据
	GetMonthDataResponse struct {
		util.CommonError
		Datas []MonthDataItem `json:"datas"`
	}

	MonthDataItem struct {
		BaseInfo       MonthBaseInfo    `json:"base_info"`
		SummaryInfo    MonthSummaryInfo `json:"summary_info"`
		ExceptionInfos []ExceptionInfo  `json:"exception_infos"`
		SpItems        []SpItem         `json:"sp_items"`
		OverWorkInfo   OverWorkInfo     `json:"overwork_info"`
	}

	// MonthBaseInfo 基础信息
	MonthBaseInfo struct {
		RecordType  int64         `json:"record_type"`
		Name        string        `json:"name"`
		NameEx      string        `json:"name_ex"`
		DepartsName string        `json:"departs_name"`
		AcctId      string        `json:"acctid"`
		RuleInfo    MonthRuleInfo `json:"rule_info"`
	}

	// MonthRuleInfo 打卡人员所属规则信息
	MonthRuleInfo struct {
		GroupId   int64  `json:"groupid"`
		GroupName string `json:"groupname"`
	}

	// MonthSummaryInfo 汇总信息
	MonthSummaryInfo struct {
		WorkDays        int64 `json:"work_days"`
		ExceptDays      int64 `json:"except_days"`
		RegularDays     int64 `json:"regular_days"`
		RegularWorkSec  int64 `json:"regular_work_sec"`
		StandardWorkSec int64 `json:"standard_work_sec"`
	}

	// OverWorkInfo 加班情况
	OverWorkInfo struct {
		WorkdayOverSec int64 `json:"workday_over_sec"`
		HolidayOverSec int64 `json:"holidays_over_sec"`
		RestDayOverSec int64 `json:"restdays_over_sec"`
	}
)

// GetMonthData 获取打卡月报数据
// @see https://developer.work.weixin.qq.com/document/path/96499
func (r *Client) GetMonthData(req *GetCheckinDataRequest) (result *GetMonthDataResponse, err error) {
	var (
		response    []byte
		accessToken string
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return
	}
	if response, err = util.PostJSON(fmt.Sprintf(getMonthDataURL, accessToken), req); err != nil {
		return
	}

	result = new(GetMonthDataResponse)
	if err = util.DecodeWithError(response, result, "GetMonthData"); err != nil {
		return
	}
	return
}
