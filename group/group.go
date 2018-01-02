package group

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	baseApi        = "https://api.weixin.qq.com/cgi-bin"
	groupCreateURL = "/groups/create"
	groupGetURL    = "/groups/get"
	groupUpdateURL = "/groups/members/update"
	groupUserURL   = "/groups/getid"
)

type Group struct {
	*context.Context
}

type ResGroupList struct {
	util.CommonError
	Groups []SelfGroup `json:"groups"`
}

type ResCreateGroup struct {
	util.CommonError
	Group struct {
		Id int64 `json:"id"`
	}
}

type ResUserGroup struct {
	util.CommonError
	GroupId int64 `json:"groupid"`
}

type SelfGroup struct {
	Name    string `json:"name,omitempty"`
	GroupId int64  `json:"id,omitempty"`
	Count   int64  `json:"count,omitempty"`
}

type reqCreateGroup struct {
	Group SelfGroup `json:"group"`
}

type reqUpdateGroup struct {
	Openid    string `json:"openid"`
	ToGroupid string `json:"to_groupid"`
}

type reqUserGroup struct {
	Openid string `json:"openid"`
}

//NewGroup 实例
func NewGroup(context *context.Context) *Group {
	group := new(Group)
	group.Context = context
	return group
}

//获取分组列表
func (group *Group) GetGroup() (resGroup ResGroupList, err error) {
	var accessToken string
	accessToken, err = group.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, groupGetURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &resGroup)
	if err != nil {
		return
	}
	if resGroup.ErrCode != 0 {
		err = fmt.Errorf("GetGroup Error , errcode=%d , errmsg=%s", resGroup.ErrCode, resGroup.ErrMsg)
		return
	}
	return
}

//更新用户所在分组
func (group *Group) UpdateUserGroup(openid, toGroupId string) error {
	accessToken, err := group.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, groupUpdateURL, accessToken)
	updateGroup := reqUpdateGroup{
		Openid:    openid,
		ToGroupid: toGroupId,
	}
	var response []byte
	response, err = util.PostJSON(uri, updateGroup)
	if err != nil {
		return err
	}
	var commError util.CommonError
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return err
	}
	if commError.ErrCode != 0 {
		return fmt.Errorf("AddGroup Error , errcode=%d , errmsg=%s", commError.ErrCode, commError.ErrMsg)
	}
	return nil
}

//添加分组
func (group *Group) AddGroup(groupName string) (groupId int64, err error) {
	accessToken, err := group.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, groupCreateURL, accessToken)
	selfGroup := SelfGroup{
		Name: groupName,
	}
	reqCreateGroup := &reqCreateGroup{
		Group: selfGroup,
	}

	var response []byte
	response, err = util.PostJSON(uri, reqCreateGroup)
	if err != nil {
		return
	}
	var resCreateGroup ResCreateGroup
	err = json.Unmarshal(response, &resCreateGroup)
	if err != nil {
		return
	}
	if resCreateGroup.ErrCode != 0 {
		err = fmt.Errorf("AddGroup Error , errcode=%d , errmsg=%s", resCreateGroup.ErrCode, resCreateGroup.ErrMsg)
		return
	}
	groupId = resCreateGroup.Group.Id
	return
}

//获取用户所在分组
func (group *Group) GetUserGroup(openid string) (groupId int64, err error) {
	accessToken, err := group.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, groupUserURL, accessToken)
	reqGroup := &reqUserGroup{
		Openid: openid,
	}

	var response []byte
	response, err = util.PostJSON(uri, reqGroup)
	if err != nil {
		return
	}
	var resUserGroup ResUserGroup
	err = json.Unmarshal(response, &resUserGroup)
	if err != nil {
		return
	}
	if resUserGroup.ErrCode != 0 {
		err = fmt.Errorf("AddGroup Error , errcode=%d , errmsg=%s", resUserGroup.ErrCode, resUserGroup.ErrMsg)
		return
	}
	groupId = resUserGroup.GroupId
	return
}
