package externalcontact

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// AddMsgTemplateURL 创建企业群发
	AddMsgTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=%s"
	// GetGroupMsgListV2URL 获取群发记录列表
	GetGroupMsgListV2URL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2?access_token=%s"
	// GetGroupMsgTaskURL 获取群发成员发送任务列表
	GetGroupMsgTaskURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task?access_token=%s"
	// GetGroupMsgSendResultURL 获取企业群发成员执行结果
	GetGroupMsgSendResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=%s"
	// SendWelcomeMsgURL 发送新客户欢迎语
	SendWelcomeMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=%s"
	// GroupWelcomeTemplateAddURL 添加入群欢迎语素材
	GroupWelcomeTemplateAddURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token=%s"
	// GroupWelcomeTemplateEditURL 编辑入群欢迎语素材
	GroupWelcomeTemplateEditURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token=%s"
	// GroupWelcomeTemplateGetURL 获取入群欢迎语素材
	GroupWelcomeTemplateGetURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token=%s"
	// GroupWelcomeTemplateDelURL 删除入群欢迎语素材
	GroupWelcomeTemplateDelURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token=%s"
)

// AddMsgTemplateRequest 创建企业群发请求
type AddMsgTemplateRequest struct {
	ChatType       string        `json:"chat_type"`
	ExternalUserID []string      `json:"external_userid"`
	Sender         string        `json:"sender,omitempty"`
	Text           MsgText       `json:"text"`
	Attachments    []*Attachment `json:"attachments"`
}

// MsgText 文本消息
type MsgText struct {
	Content string `json:"content"`
}

type (
	// Attachment 附件
	Attachment struct {
		MsgType     string                `json:"msgtype"`
		Image       AttachmentImg         `json:"image,omitempty"`
		Link        AttachmentLink        `json:"link,omitempty"`
		MiniProgram AttachmentMiniProgram `json:"miniprogram,omitempty"`
		Video       AttachmentVideo       `json:"video,omitempty"`
		File        AttachmentFile        `json:"file,omitempty"`
	}
	// AttachmentImg 图片消息
	AttachmentImg struct {
		MediaID string `json:"media_id"`
		PicURL  string `json:"pic_url"`
	}
	// AttachmentLink 图文消息
	AttachmentLink struct {
		Title  string `json:"title"`
		PicURL string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	}
	// AttachmentMiniProgram 小程序消息
	AttachmentMiniProgram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		AppID      string `json:"appid"`
		Page       string `json:"page"`
	}
	// AttachmentVideo 视频消息
	AttachmentVideo struct {
		MediaID string `json:"media_id"`
	}
	// AttachmentFile 文件消息
	AttachmentFile struct {
		MediaID string `json:"media_id"`
	}
)

// AddMsgTemplateResponse 创建企业群发响应
type AddMsgTemplateResponse struct {
	util.CommonError
	FailList []string `json:"fail_list"`
	MsgID    string   `json:"msgid"`
}

// AddMsgTemplate 创建企业群发
// see https://developer.work.weixin.qq.com/document/path/92135
func (r *Client) AddMsgTemplate(req *AddMsgTemplateRequest) (*AddMsgTemplateResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(AddMsgTemplateURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &AddMsgTemplateResponse{}
	if err = util.DecodeWithError(response, result, "AddMsgTemplate"); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupMsgListV2Request 获取群发记录列表请求
type GetGroupMsgListV2Request struct {
	ChatType   string `json:"chat_type"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type"`
	Limit      int    `json:"limit"`
	Cursor     string `json:"cursor"`
}

// GetGroupMsgListV2Response 获取群发记录列表响应
type GetGroupMsgListV2Response struct {
	util.CommonError
	NextCursor   string      `json:"next_cursor"`
	GroupMsgList []*GroupMsg `json:"group_msg_list"`
}

// GroupMsg 群发消息
type GroupMsg struct {
	MsgID       string        `json:"msgid"`
	Creator     string        `json:"creator"`
	CreateTime  int           `json:"create_time"`
	CreateType  int           `json:"create_type"`
	Text        MsgText       `json:"text"`
	Attachments []*Attachment `json:"attachments"`
}

// GetGroupMsgListV2 获取群发记录列表
// see https://developer.work.weixin.qq.com/document/path/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8
func (r *Client) GetGroupMsgListV2(req *GetGroupMsgListV2Request) (*GetGroupMsgListV2Response, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GetGroupMsgListV2URL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetGroupMsgListV2Response{}
	if err = util.DecodeWithError(response, result, "GetGroupMsgListV2"); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupMsgTaskRequest 获取群发成员发送任务列表请求
type GetGroupMsgTaskRequest struct {
	MsgID  string `json:"msgid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

// GetGroupMsgTaskResponse 获取群发成员发送任务列表响应
type GetGroupMsgTaskResponse struct {
	util.CommonError
	NextCursor string  `json:"next_cursor"`
	TaskList   []*Task `json:"task_list"`
}

// Task 获取群发成员发送任务列表任务
type Task struct {
	UserID   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int    `json:"send_time"`
}

// GetGroupMsgTask 获取群发成员发送任务列表
// see https://developer.work.weixin.qq.com/document/path/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E5%8F%91%E9%80%81%E4%BB%BB%E5%8A%A1%E5%88%97%E8%A1%A8
func (r *Client) GetGroupMsgTask(req *GetGroupMsgTaskRequest) (*GetGroupMsgTaskResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GetGroupMsgTaskURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetGroupMsgTaskResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupMsgTask"); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupMsgSendResultRequest 获取企业群发成员执行结果请求
type GetGroupMsgSendResultRequest struct {
	MsgID  string `json:"msgid"`
	UserID string `json:"userid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

// GetGroupMsgSendResultResponse 获取企业群发成员执行结果响应
type GetGroupMsgSendResultResponse struct {
	util.CommonError
	NextCursor string  `json:"next_cursor"`
	SendList   []*Send `json:"send_list"`
}

// Send 企业群发成员执行结果
type Send struct {
	ExternalUserID string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int    `json:"send_time"`
}

// GetGroupMsgSendResult 获取企业群发成员执行结果
// see https://developer.work.weixin.qq.com/document/path/93338#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E6%89%A7%E8%A1%8C%E7%BB%93%E6%9E%9C
func (r *Client) GetGroupMsgSendResult(req *GetGroupMsgSendResultRequest) (*GetGroupMsgSendResultResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GetGroupMsgSendResultURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetGroupMsgSendResultResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupMsgSendResult"); err != nil {
		return nil, err
	}
	return result, nil
}

// SendWelcomeMsgRequest 发送新客户欢迎语请求
type SendWelcomeMsgRequest struct {
	WelcomeCode string        `json:"welcome_code"`
	Text        MsgText       `json:"text"`
	Attachments []*Attachment `json:"attachments"`
}

// SendWelcomeMsgResponse 发送新客户欢迎语响应
type SendWelcomeMsgResponse struct {
	util.CommonError
}

// SendWelcomeMsg 发送新客户欢迎语
// see https://developer.work.weixin.qq.com/document/path/92137
func (r *Client) SendWelcomeMsg(req *SendWelcomeMsgRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(SendWelcomeMsgURL, accessToken), req); err != nil {
		return err
	}
	result := &SendWelcomeMsgResponse{}
	if err = util.DecodeWithError(response, result, "SendWelcomeMsg"); err != nil {
		return err
	}
	return nil
}

// GroupWelcomeTemplateAddRequest 添加入群欢迎语素材请求
type GroupWelcomeTemplateAddRequest struct {
	Text        MsgText               `json:"text"`
	Image       AttachmentImg         `json:"image"`
	Link        AttachmentLink        `json:"link"`
	MiniProgram AttachmentMiniProgram `json:"miniprogram"`
	File        AttachmentFile        `json:"file"`
	Video       AttachmentVideo       `json:"video"`
	AgentID     int                   `json:"agentid"`
	Notify      int                   `json:"notify"`
}

// GroupWelcomeTemplateAddResponse 添加入群欢迎语素材响应
type GroupWelcomeTemplateAddResponse struct {
	util.CommonError
	TemplateID string `json:"template_id"`
}

// GroupWelcomeTemplateAdd 添加入群欢迎语素材
// see https://developer.work.weixin.qq.com/document/path/92366#%E6%B7%BB%E5%8A%A0%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (r *Client) GroupWelcomeTemplateAdd(req *GroupWelcomeTemplateAddRequest) (*GroupWelcomeTemplateAddResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GroupWelcomeTemplateAddURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GroupWelcomeTemplateAddResponse{}
	if err = util.DecodeWithError(response, result, "GroupWelcomeTemplateAdd"); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupWelcomeTemplateEditRequest 编辑入群欢迎语素材请求
type GroupWelcomeTemplateEditRequest struct {
	TemplateID  string                `json:"template_id"`
	Text        MsgText               `json:"text"`
	Image       AttachmentImg         `json:"image"`
	Link        AttachmentLink        `json:"link"`
	MiniProgram AttachmentMiniProgram `json:"miniprogram"`
	File        AttachmentFile        `json:"file"`
	Video       AttachmentVideo       `json:"video"`
	AgentID     int                   `json:"agentid"`
}

// GroupWelcomeTemplateEditResponse 编辑入群欢迎语素材响应
type GroupWelcomeTemplateEditResponse struct {
	util.CommonError
}

// GroupWelcomeTemplateEdit 编辑入群欢迎语素材
// see https://developer.work.weixin.qq.com/document/path/92366#%E7%BC%96%E8%BE%91%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (r *Client) GroupWelcomeTemplateEdit(req *GroupWelcomeTemplateEditRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GroupWelcomeTemplateEditURL, accessToken), req); err != nil {
		return err
	}
	result := &GroupWelcomeTemplateEditResponse{}
	if err = util.DecodeWithError(response, result, "GroupWelcomeTemplateEdit"); err != nil {
		return err
	}
	return nil
}

// GroupWelcomeTemplateGetRequest 获取入群欢迎语素材请求
type GroupWelcomeTemplateGetRequest struct {
	TemplateID string `json:"template_id"`
}

// GroupWelcomeTemplateGetResponse 获取入群欢迎语素材响应
type GroupWelcomeTemplateGetResponse struct {
	util.CommonError
	Text        MsgText               `json:"text"`
	Image       AttachmentImg         `json:"image"`
	Link        AttachmentLink        `json:"link"`
	MiniProgram AttachmentMiniProgram `json:"miniprogram"`
	File        AttachmentFile        `json:"file"`
	Video       AttachmentVideo       `json:"video"`
}

// GroupWelcomeTemplateGet 获取入群欢迎语素材
// see https://developer.work.weixin.qq.com/document/path/92366#%E8%8E%B7%E5%8F%96%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (r *Client) GroupWelcomeTemplateGet(req *GroupWelcomeTemplateGetRequest) (*GroupWelcomeTemplateGetResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GroupWelcomeTemplateGetURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GroupWelcomeTemplateGetResponse{}
	if err = util.DecodeWithError(response, result, "GroupWelcomeTemplateGet"); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupWelcomeTemplateDelRequest 删除入群欢迎语素材请求
type GroupWelcomeTemplateDelRequest struct {
	TemplateID string `json:"template_id"`
	AgentID    int    `json:"agentid"`
}

// GroupWelcomeTemplateDelResponse 删除入群欢迎语素材响应
type GroupWelcomeTemplateDelResponse struct {
	util.CommonError
}

// GroupWelcomeTemplateDel 删除入群欢迎语素材
// see https://developer.work.weixin.qq.com/document/path/92366#%E5%88%A0%E9%99%A4%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (r *Client) GroupWelcomeTemplateDel(req *GroupWelcomeTemplateDelRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GroupWelcomeTemplateDelURL, accessToken), req); err != nil {
		return err
	}
	result := &GroupWelcomeTemplateDelResponse{}
	if err = util.DecodeWithError(response, result, "GroupWelcomeTemplateDel"); err != nil {
		return err
	}
	return nil
}
