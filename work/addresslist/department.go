// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/12

package addresslist

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// DepartmentSimpleListURL 获取子部门ID列表
	DepartmentSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"
)

type (
	// DepartmentSimpleListResponse 获取子部门ID列表响应
	DepartmentSimpleListResponse struct {
		util.CommonError
		DepartmentId []*DepartmentID `json:"department_id"`
	}
	DepartmentID struct {
		ID       int `json:"id"`
		ParentID int `json:"parentid"`
		Order    int `json:"order"`
	}
)

// DepartmentSimpleList 获取子部门ID列表
// see https://developer.work.weixin.qq.com/document/path/95350
func (r *Client) DepartmentSimpleList(departmentId int) ([]*DepartmentID, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(DepartmentSimpleListURL, accessToken, departmentId)); err != nil {
		return nil, err
	}
	result := &DepartmentSimpleListResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentSimpleList"); err != nil {
		return nil, err
	}
	return result.DepartmentId, nil
}
