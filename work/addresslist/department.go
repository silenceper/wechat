package addresslist

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// departmentCreateURL 创建部门
	departmentCreateURL = "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s"
	// departmentSimpleListURL 获取子部门ID列表
	departmentSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"
	// departmentListURL 获取部门列表
	departmentListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s"
)

type (
	// DepartmentCreateRequest 创建部门数据请求
	DepartmentCreateRequest struct {
		Name     string `json:"name"`
		NameEn   string `json:"name_en,omitempty"`
		ParentID int    `json:"parentid"`
		Order    int    `json:"order,omitempty"`
		Id       int    `json:"id,omitempty"`
	}
	// DepartmentCreateResponse 创建部门数据响应
	DepartmentCreateResponse struct {
		util.CommonError
		ID int `json:"id"`
	}

	// DepartmentSimpleListResponse 获取子部门ID列表响应
	DepartmentSimpleListResponse struct {
		util.CommonError
		DepartmentID []*DepartmentID `json:"department_id"`
	}
	// DepartmentID 子部门ID
	DepartmentID struct {
		ID       int `json:"id"`
		ParentID int `json:"parentid"`
		Order    int `json:"order"`
	}

	// DepartmentListResponse 获取部门列表响应
	DepartmentListResponse struct {
		util.CommonError
		Department []*Department `json:"department"`
	}
	// Department 部门列表数据
	Department struct {
		ID               int      `json:"id"`                // 创建的部门id
		Name             string   `json:"name"`              // 部门名称
		NameEn           string   `json:"name_en"`           // 英文名称
		DepartmentLeader []string `json:"department_leader"` // 部门负责人的UserID
		ParentID         int      `json:"parentid"`          // 父部门id。根部门为1
		Order            int      `json:"order"`             // 在父部门中的次序值。order值大的排序靠前
	}
)

// DepartmentCreate 创建部门
// see https://developer.work.weixin.qq.com/document/path/90205
func (r *Client) DepartmentCreate(req *DepartmentCreateRequest) (*DepartmentCreateResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(departmentCreateURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &DepartmentCreateResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentCreate"); err != nil {
		return nil, err
	}
	return result, nil
}

// DepartmentSimpleList 获取子部门ID列表
// see https://developer.work.weixin.qq.com/document/path/95350
func (r *Client) DepartmentSimpleList(departmentID int) ([]*DepartmentID, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(departmentSimpleListURL, accessToken, departmentID)); err != nil {
		return nil, err
	}
	result := &DepartmentSimpleListResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentSimpleList"); err != nil {
		return nil, err
	}
	return result.DepartmentID, nil
}

// DepartmentList 获取部门列表
// @desc https://developer.work.weixin.qq.com/document/path/90208
func (r *Client) DepartmentList() ([]*Department, error) {
	// 获取accessToken
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 发起http请求
	response, err := util.HTTPGet(fmt.Sprintf(departmentListURL, accessToken))
	if err != nil {
		return nil, err
	}
	// 按照结构体解析返回值
	result := &DepartmentListResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentList"); err != nil {
		return nil, err
	}
	// 返回数据
	return result.Department, err
}
