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
)

type Group struct {
	*context.Context
}

type ResGroup struct {
	util.CommonError
	Groups []struct {
		Name    string `json:"name"`
		GroupId int64  `json:"id"`
		Count   int64  `json:"count"`
	} `json:"groups"`
}

type RecCreateGroup struct {
	util.CommonError
	Group struct {
		Id int64 `json:"id"`
	}
}

type reqGroup struct {
	Name string `json:name`
}

type reqUpdateGroup struct {
	Openid    string `json:openid`
	ToGroupid string `json:to_groupid`
}

//NewGroup 实例
func NewGroup(context *context.Context) *Group {
	group := new(Group)
	group.Context = context
	return group
}

//获取分组列表
func (group *Group) GetGroup() (resGroup ResGroup, err error) {
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
	reqGroup := &reqGroup{
		Name: groupName,
	}

	var response []byte
	response, err = util.PostJSON(uri, reqGroup)
	if err != nil {
		return
	}
	var recCreateGroup RecCreateGroup
	err = json.Unmarshal(response, &recCreateGroup)
	if err != nil {
		return
	}
	if recCreateGroup.ErrCode != 0 {
		err = fmt.Errorf("AddGroup Error , errcode=%d , errmsg=%s", recCreateGroup.ErrCode, recCreateGroup.ErrMsg)
		return
	}
	groupId = recCreateGroup.Group.Id
	return
}
