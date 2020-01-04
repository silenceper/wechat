package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	commitURL                 = "https://api.weixin.qq.com/wxa/commit"
	getCodePageURL            = "https://api.weixin.qq.com/wxa/get_page"
	getTestQrcodeURL          = "https://api.weixin.qq.com/wxa/get_qrcode"
	submitAuditURL            = "https://api.weixin.qq.com/wxa/submit_audit"
	getAuditStatusURL         = "https://api.weixin.qq.com/wxa/get_auditstatus"
	getLatestAuditStatusURL   = "https://api.weixin.qq.com/wxa/get_latest_auditstatus"
	undoCodeAuditURL          = "https://api.weixin.qq.com/wxa/undocodeaudit"
	releaseURL                = "https://api.weixin.qq.com/wxa/release"
	revertCodeReleaseURL      = "https://api.weixin.qq.com/wxa/revertcoderelease"
	grayReleaseURL            = "https://api.weixin.qq.com/wxa/grayrelease"
	getGrayReleasePlanURL     = "https://api.weixin.qq.com/wxa/getgrayreleaseplan"
	revertGrayReleaseURL      = "https://api.weixin.qq.com/wxa/revertgrayrelease"
	changeVisitStatusURL      = "https://api.weixin.qq.com/wxa/change_visitstatus"
	getWeappSupportVersionURL = "https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion"
	setWeappSupportVersionURL = "https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion"
	queryQuotaURL             = "https://api.weixin.qq.com/wxa/queryquota"
	speedUpAuditURL           = "https://api.weixin.qq.com/wxa/speedupaudit"
)

// CommitParam 提交代码参数
type CommitParam struct {
	TemplateID  int            `json:"template_id"`  // 模板编号
	Ext         CommitParamExt `json:"-"`            // 扩展
	ExtJSON     string         `json:"ext_json"`     // 扩展
	UserVersion string         `json:"user_version"` // 提交版本
	UserDesc    string         `json:"user_desc"`    // 版本说明
}

// CommitParamExt 此处还能支持更多，不过貌似没啥用
type CommitParamExt struct {
	ExtAppID string            `json:"extAppid"` // appid
	Ext      map[string]string `json:"ext"`      // 附加扩展配置
}

// CodePageList 已上传的代码的页面列表
type CodePageList struct {
	util.CommonError
	PageList []string `json:"page_list"`
}

// SubmitAuditParam 提审参数
type SubmitAuditParam struct {
	ItemList      []SubmitAuditItem `json:"item_list"` // 审核项列表，不知道微信要这玩意干啥，sdk自动注入一条
	FeedbackInfo  string            `json:"feedback_info"`
	FeedbackStuff string            `json:"feedback_stuff"`
}
type SubmitAuditItem struct {
	Address     string `json:"address"`
	Tag         string `json:"tag"`
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	ThirdClass  string `json:"third_class"`
	FirstID     int    `json:"first_id"`
	SecondID    int    `json:"second_id"`
	ThirdID     int    `json:"third_id"`
	Title       string `json:"title"`
}

// SubmitAuditResponse 提审返回
type submitAuditResponse struct {
	util.CommonError
	AuditID uint64 `json:"auditid"`
}

// AuditStatusResponse 获取审核结果
type AuditStatusResponse struct {
	util.CommonError
	AuditID    uint64 `json:"auditid"`
	Status     int    `json:"status"` // 0-审核成功 1-审核被拒绝 2-审核中 3-已撤回
	Reason     string `json:"reason"`
	ScreenShot string `json:"ScreenShot"`
}

// GrayReleasePlanResponse 分阶段发布计划结果
type GrayReleasePlanResponse struct {
	util.CommonError
	GrayReleasePlan struct {
		Status          int `json:"status"`           // 0:初始状态 1:执行中 2:暂停中 3:执行完毕 4:被删除
		CreateTimestamp int `json:"create_timestamp"` // 分阶段发布计划的创建时间
		GrayPercentage  int `json:"gray_percentage"`  // 当前灰度比例
	} `json:"gray_release_plan"`
}

// WeappSupportVersionResponse 最低基础库版本及各版本用户占比
type WeappSupportVersionResponse struct {
	util.CommonError
	NowVersion string `json:"now_version"`
	UvInfo     struct {
		Items []struct {
			Percentage float32 `json:"percentage"`
			Version    string  `json:"version"`
		}
	} `json:"uv_info"`
}

// QueryQuotaResponse 当月提审限额（quota）和加急次数
type QueryQuotaResponse struct {
	util.CommonError
	Rest         int `json:"rest"`          // quota剩余值
	Limit        int `json:"limit"`         // 当月分配quota
	SpeedupRest  int `json:"speedup_rest"`  // 剩余加急次数
	SpeedupLimit int `json:"speedup_limit"` // 当月分配加急次数
}

// Commit 上传小程序代码
func (m *MiniPrograms) Commit(param CommitParam) (err error) {
	var body []byte
	ret := util.CommonError{}
	if param.Ext.ExtAppID == "" {
		param.Ext.ExtAppID = m.AuthAppID
	}
	if param.ExtJSON == "" {
		var extJsonByte []byte
		extJsonByte, err = json.Marshal(param.Ext)
		if err != nil {
			return
		}
		param.ExtJSON = string(extJsonByte)
	}
	body, err = m.post(commitURL, param)
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

// GetCodePage 获取已上传的代码的页面列表
func (m *MiniPrograms) GetCodePage() (ret CodePageList, err error) {
	var body []byte
	body, err = m.get(getCodePageURL, nil)
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

// GetTestQrcode 获取体验二维码
func (m *MiniPrograms) GetTestQrcode(path string) (ret []byte, err error) {
	rmap := map[string]string{
		"path": path,
	}
	ret, err = m.getBinary(getTestQrcodeURL, rmap)
	return
}

// SubmitAudit 提审
func (m *MiniPrograms) SubmitAudit(param SubmitAuditParam) (auditID uint64, err error) {
	ret := submitAuditResponse{}
	body, err := m.post(submitAuditURL, param)
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
	auditID = ret.AuditID
	return
}

// GetAuditStatus 查询指定发布审核单的审核状态
func (m *MiniPrograms) GetAuditStatus(auditID uint64) (ret AuditStatusResponse, err error) {
	rmap := map[string]uint64{
		"auditid": auditID,
	}
	body, err := m.post(getAuditStatusURL, rmap)
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

// GetLatestAuditStatus 查询最新一次提交的审核状态
func (m *MiniPrograms) GetLatestAuditStatus() (ret AuditStatusResponse, err error) {
	body, err := m.get(getLatestAuditStatusURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// UndoCodeAudit 小程序审核撤回
// 调用本接口可以撤回当前的代码审核单
// 注意： 单个帐号每天审核撤回次数最多不超过 1 次，一个月不超过 10 次。
func (m *MiniPrograms) UndoCodeAudit() (err error) {
	body, err := m.get(undoCodeAuditURL, nil)
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// Release 发布已通过审核的小程序
// 调用本接口可以发布最后一个审核通过的小程序代码版本
func (m *MiniPrograms) Release() (err error) {
	body, err := m.post(releaseURL, nil)
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// RevertCodeRelease 版本回退
// 调用本接口可以将小程序的线上版本进行回退
// 如果没有上一个线上版本，将无法回退
// 只能向上回退一个版本，即当前版本回退后，不能再调用版本回退接口
func (m *MiniPrograms) RevertCodeRelease() (err error) {
	body, err := m.get(revertCodeReleaseURL, nil)
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// GrayRelease 分阶段发布
// gray 灰度的百分比 1 ~ 100 的整数
func (m *MiniPrograms) GrayRelease(gray int) (err error) {
	body, err := m.post(grayReleaseURL, map[string]int{
		"gray_percentage": gray,
	})
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// GetGrayReleasePlan 查询当前分阶段发布详情
func (m *MiniPrograms) GetGrayReleasePlan() (ret GrayReleasePlanResponse, err error) {
	body, err := m.get(getGrayReleasePlanURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// RevertGrayRelease 取消分阶段发布
// 在小程序分阶段发布期间，可以随时调用本接口取消分阶段发布。
// 取消分阶段发布后，受影响的微信用户（即被灰度升级的微信用户）的小程序版本将回退到分阶段发布前的版本
func (m *MiniPrograms) RevertGrayRelease() (err error) {
	body, err := m.get(revertGrayReleaseURL, nil)
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// ChangeVisitStatus 修改小程序线上代码的可见状态（仅供第三方代小程序调用）
func (m *MiniPrograms) ChangeVisitStatus(visit bool) (err error) {
	var param = make(map[string]Action)
	if visit {
		param["action"] = ActionOpen
	} else {
		param["action"] = ActionClose
	}
	body, err := m.post(changeVisitStatusURL, param)
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// GetWeappSupportVersion 查询当前设置的最低基础库版本及各版本用户占比
// 调用本接口可以查询小程序当前设置的最低基础库版本，以及小程序在各个基础库版本的用户占比
func (m *MiniPrograms) GetWeappSupportVersion() (ret WeappSupportVersionResponse, err error) {
	body, err := m.post(getWeappSupportVersionURL, nil)
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

// SetWeappSupportVersion 设置最低基础库版本
// 调用本接口可以设置小程序的最低基础库支持版本，可以先查询当前小程序在各个基础库的用户占比来辅助进行决策
func (m *MiniPrograms) SetWeappSupportVersion(version string) (err error) {
	body, err := m.post(setWeappSupportVersionURL, map[string]string{
		"version": version,
	})
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// QueryQuota 查询服务商的当月提审限额（quota）和加急次数
// 服务商可以调用该接口，查询当月平台分配的提审限额和剩余可提审次数，以及当月分配的审核加急次数和剩余加急次数。（所有旗下小程序共用该额度）
func (m *MiniPrograms) QueryQuota() (ret QueryQuotaResponse, err error) {
	body, err := m.get(queryQuotaURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
		return
	}
	return
}

// SpeedUpAudit 加急审核申请
// 有加急次数的第三方可以通过该接口，对已经提审的小程序进行加急操作，加急后的小程序预计2-12小时内审完
func (m *MiniPrograms) SpeedUpAudit(auditID uint64) (err error) {
	body, err := m.post(speedUpAuditURL, map[string]uint64{
		"auditid": auditID,
	})
	if err != nil {
		return
	}
	var ret util.CommonError
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}
