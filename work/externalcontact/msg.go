package externalcontact

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// AddMsgTemplateURL 创建企业群发
	AddMsgTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=%s"
	// GetGroupMsgTaskURL 获取群发成员发送任务列表
	GetGroupMsgTaskURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task?access_token=%s"
	// GetGroupMsgSendResultURL 获取企业群发成员执行结果
	GetGroupMsgSendResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=%s"
	// SendWelcomeMsgURL 发送新客户欢迎语
	SendWelcomeMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=%s"
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
// see https://developer.work.weixin.qq.com/document/path/93338
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
// see https://developer.work.weixin.qq.com/document/path/93338
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
