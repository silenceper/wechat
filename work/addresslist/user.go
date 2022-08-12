// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/12

package addresslist

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// UserSimpleListURL 获取部门成员
	UserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d"
)

type (
	// UserSimpleListResponse 获取部门成员响应
	UserSimpleListResponse struct {
		util.CommonError
		UserList []*UserList
	}
	// UserList 部门成员
	UserList struct {
		UserID     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserID string `json:"open_userid"`
	}
)

// UserSimpleList 获取部门成员
// @see https://developer.work.weixin.qq.com/document/path/90200
func (r *Client) UserSimpleList(departmentID int) ([]*UserList, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(UserSimpleListURL, accessToken, departmentID)); err != nil {
		return nil, err
	}
	result := &UserSimpleListResponse{}
	err = util.DecodeWithError(response, result, "UserSimpleList")
	if err != nil {
		return nil, err
	}
	return result.UserList, nil
}
