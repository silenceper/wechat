package externalcontact

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

// GroupChatURL 客户群
const GroupChatURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat"

type (
	AddJoinWayRequest struct {
		Scene          int      `json:"scene"`            // 必填 1 - 群的小程序插件,2 - 群的二维码插件
		Remark         string   `json:"remark"`           //非必填	联系方式的备注信息，用于助记，超过30个字符将被截断
		AutoCreateRoom int      `json:"auto_create_room"` //非必填	当群满了后，是否自动新建群。0-否；1-是。 默认为1
		RoomBaseName   string   `json:"room_base_name"`   //非必填	自动建群的群名前缀，当auto_create_room为1时有效。最长40个utf8字符
		RoomBaseId     int      `json:"room_base_id"`     //非必填	自动建群的群起始序号，当auto_create_room为1时有效
		ChatIdList     []string `json:"chat_id_list"`     //必填	使用该配置的客户群ID列表，支持5个。见客户群ID获取方法
		State          string   `json:"state"`            //非必填	企业自定义的state参数，用于区分不同的入群渠道。不超过30个UTF-8字符
	}

	AddJoinWayResponse struct {
		util.CommonError
		ConfigId string `json:"config_id"`
	}
)

// AddJoinWay 加入群聊
// @see https://developer.work.weixin.qq.com/document/path/92229
func (r *Client) AddJoinWay(req *AddJoinWayRequest) (*AddJoinWayResponse, error) {
	var (
		accessToken string
		err         error
		response    []byte
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	response, err = util.PostJSON(fmt.Sprintf("%s/add_join_way?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &AddJoinWayResponse{}
	if err = util.DecodeWithError(response, result, "AddJoinWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	JoinWayConfigRequest struct {
		ConfigId string `json:"config_id"`
	}

	JoinWay struct {
		ConfigId       string   `json:"config_id"`
		Scene          int      `json:"scene"`
		Remark         string   `json:"remark"`
		AutoCreateRoom int      `json:"auto_create_room"`
		RoomBaseName   string   `json:"room_base_name"`
		RoomBaseId     int      `json:"room_base_id"`
		ChatIdList     []string `json:"chat_id_list"`
		QrCode         string   `json:"qr_code"`
		State          string   `json:"state"`
	}

	GetJoinWayResponse struct {
		util.CommonError
		JoinWay JoinWay `json:"join_way"`
	}
)

// GetJoinWay 获取客户群进群方式配置
// @see https://developer.work.weixin.qq.com/document/path/92229
func (r *Client) GetJoinWay(req *JoinWayConfigRequest) (*GetJoinWayResponse, error) {
	var (
		accessToken string
		err         error
		response    []byte
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	response, err = util.PostJSON(fmt.Sprintf("%s/get_join_way?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &GetJoinWayResponse{}
	if err = util.DecodeWithError(response, result, "GetJoinWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type UpdateJoinWayRequest struct {
	ConfigId       string   `json:"config_id"`
	Scene          int      `json:"scene"`            // 必填 1 - 群的小程序插件,2 - 群的二维码插件
	Remark         string   `json:"remark"`           //非必填	联系方式的备注信息，用于助记，超过30个字符将被截断
	AutoCreateRoom int      `json:"auto_create_room"` //非必填	当群满了后，是否自动新建群。0-否；1-是。 默认为1
	RoomBaseName   string   `json:"room_base_name"`   //非必填	自动建群的群名前缀，当auto_create_room为1时有效。最长40个utf8字符
	RoomBaseId     int      `json:"room_base_id"`     //非必填	自动建群的群起始序号，当auto_create_room为1时有效
	ChatIdList     []string `json:"chat_id_list"`     //必填	使用该配置的客户群ID列表，支持5个。见客户群ID获取方法
	State          string   `json:"state"`            //非必填	企业自定义的state参数，用于区分不同的入群渠道。不超过30个UTF-8字符
}

// UpdateJoinWay 更新客户群进群方式配置
// @see https://developer.work.weixin.qq.com/document/path/92229
func (r *Client) UpdateJoinWay(req *UpdateJoinWayRequest) (*util.CommonError, error) {
	var (
		accessToken string
		err         error
		response    []byte
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	response, err = util.PostJSON(fmt.Sprintf("%s/update_join_way?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &util.CommonError{}
	if err = util.DecodeWithError(response, result, "UpdateJoinWay"); err != nil {
		return nil, err
	}
	return result, nil
}

// DelJoinWay 删除客户群进群方式配置
// @see https://developer.work.weixin.qq.com/document/path/92229
func (r *Client) DelJoinWay(req *JoinWayConfigRequest) (any, error) {
	var (
		accessToken string
		err         error
		response    []byte
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	response, err = util.PostJSON(fmt.Sprintf("%s/del_join_way?access_token=%s", GroupChatURL, accessToken), req)
	if err != nil {
		return nil, err
	}
	result := &util.CommonError{}
	if err = util.DecodeWithError(response, result, "DelJoinWay"); err != nil {
		return nil, err
	}
	return result, nil
}
