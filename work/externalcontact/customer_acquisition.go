package externalcontact

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// listLinkUrl 获取获客链接列表
	listLinkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/list_link?access_token=%s"
	// getCustomerAcquisition 获取获客链接详情
	getCustomerAcquisitionURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/get?access_token=%s"
	// createCustomerAcquisitionLink 创建获客链接
	createCustomerAcquisitionLinkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/create_link?access_token=%s"
	// updateCustomerAcquisitionLink 编辑获客链接
	updateCustomerAcquisitionLinkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/update_link?access_token=%s"
	// deleteCustomerAcquisitionLink 删除获客链接
	deleteCustomerAcquisitionLinkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/delete_link?access_token=%s"
	// getCustomerInfoWithCustomerAcquisitionLinkURL 获取由获客链接添加的客户信息
	getCustomerInfoWithCustomerAcquisitionLinkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/customer?access_token=%s"
	// customerAcquisitionQuota 查询剩余使用量
	customerAcquisitionQuotaURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition_quota?access_token=%s"
	// customerAcquisitionStatistic 查询链接使用详情
	customerAcquisitionStatisticURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/statistic?access_token=%s"
)

type (
	// ListLinkRequest 获取获客链接列表请求
	ListLinkRequest struct {
		Limit  int    `json:"limit"`
		Cursor string `json:"cursor"`
	}
	// ListLinkResponse 获取获客链接列表响应
	ListLinkResponse struct {
		util.CommonError
		LinkIDList []string `json:"link_id_list"`
		NextCursor string   `json:"next_cursor"`
	}
)

// ListLink 获客助手--获取获客链接列表
// see https://developer.work.weixin.qq.com/document/path/97297
func (r *Client) ListLink(req *ListLinkRequest) (*ListLinkResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(listLinkURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &ListLinkResponse{}
	if err = util.DecodeWithError(response, result, "ListLink"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetCustomerAcquisitionRequest 获取获客链接详情请求
	GetCustomerAcquisitionRequest struct {
		LinkID string `json:"link_id"`
	}
	// GetCustomerAcquisitionResponse 获取获客链接详情响应
	GetCustomerAcquisitionResponse struct {
		util.CommonError
		Link       Link                     `json:"link"`
		Range      CustomerAcquisitionRange `json:"range"`
		SkipVerify bool                     `json:"skip_verify"`
	}
	// Link 获客链接
	Link struct {
		LinkID     string `json:"link_id"`
		LinkName   string `json:"link_name"`
		URL        string `json:"url"`
		CreateTime int64  `json:"create_time"`
	}

	// CustomerAcquisitionRange 该获客链接使用范围
	CustomerAcquisitionRange struct {
		UserList       []string `json:"user_list"`
		DepartmentList []int64  `json:"department_list"`
	}
)

// GetCustomerAcquisition 获客助手--获取获客链接详情
// see https://developer.work.weixin.qq.com/document/path/97297
func (r *Client) GetCustomerAcquisition(req *GetCustomerAcquisitionRequest) (*GetCustomerAcquisitionResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(getCustomerAcquisitionURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetCustomerAcquisitionResponse{}
	if err = util.DecodeWithError(response, result, "GetCustomerAcquisition"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// CreateCustomerAcquisitionLinkRequest 创建获客链接请求
	CreateCustomerAcquisitionLinkRequest struct {
		LinkName   string                   `json:"link_name"`
		Range      CustomerAcquisitionRange `json:"range"`
		SkipVerify bool                     `json:"skip_verify"`
	}
	// CreateCustomerAcquisitionLinkResponse 创建获客链接响应
	CreateCustomerAcquisitionLinkResponse struct {
		util.CommonError
		Link Link `json:"link"`
	}
)

// CreateCustomerAcquisitionLink 获客助手--创建获客链接
// see https://developer.work.weixin.qq.com/document/path/97297
func (r *Client) CreateCustomerAcquisitionLink(req *CreateCustomerAcquisitionLinkRequest) (*CreateCustomerAcquisitionLinkResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(createCustomerAcquisitionLinkURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &CreateCustomerAcquisitionLinkResponse{}
	if err = util.DecodeWithError(response, result, "CreateCustomerAcquisitionLink"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// UpdateCustomerAcquisitionLinkRequest 编辑获客链接请求
	UpdateCustomerAcquisitionLinkRequest struct {
		LinkID     string                   `json:"link_id"`
		LinkName   string                   `json:"link_name"`
		Range      CustomerAcquisitionRange `json:"range"`
		SkipVerify bool                     `json:"skip_verify"`
	}
	// UpdateCustomerAcquisitionLinkResponse 编辑获客链接响应
	UpdateCustomerAcquisitionLinkResponse struct {
		util.CommonError
	}
)

// UpdateCustomerAcquisitionLink 获客助手--编辑获客链接
// see https://developer.work.weixin.qq.com/document/path/97297
func (r *Client) UpdateCustomerAcquisitionLink(req *UpdateCustomerAcquisitionLinkRequest) (*UpdateCustomerAcquisitionLinkResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(updateCustomerAcquisitionLinkURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &UpdateCustomerAcquisitionLinkResponse{}
	if err = util.DecodeWithError(response, result, "UpdateCustomerAcquisitionLink"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// DeleteCustomerAcquisitionLinkRequest 删除获客链接请求
	DeleteCustomerAcquisitionLinkRequest struct {
		LinkID string `json:"link_id"`
	}
	// DeleteCustomerAcquisitionLinkResponse 删除获客链接响应
	DeleteCustomerAcquisitionLinkResponse struct {
		util.CommonError
	}
)

// DeleteCustomerAcquisitionLink 获客助手--删除获客链接
// see https://developer.work.weixin.qq.com/document/path/97297
func (r *Client) DeleteCustomerAcquisitionLink(req *DeleteCustomerAcquisitionLinkRequest) (*DeleteCustomerAcquisitionLinkResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(deleteCustomerAcquisitionLinkURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &DeleteCustomerAcquisitionLinkResponse{}
	if err = util.DecodeWithError(response, result, "DeleteCustomerAcquisitionLink"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetCustomerInfoWithCustomerAcquisitionLinkRequest 获取由获客链接添加的客户信息请求
	GetCustomerInfoWithCustomerAcquisitionLinkRequest struct {
		LinkID string `json:"link_id"`
		Limit  int64  `json:"limit"`
		Cursor string `json:"cursor"`
	}
	// GetCustomerInfoWithCustomerAcquisitionLinkResponse 获取由获客链接添加的客户信息响应
	GetCustomerInfoWithCustomerAcquisitionLinkResponse struct {
		util.CommonError
		CustomerList []CustomerList `json:"customer_list"`
		NextCursor   string         `json:"next_cursor"`
	}
	// CustomerList 客户列表
	CustomerList struct {
		ExternalUserid string `json:"external_userid"`
		Userid         string `json:"userid"`
		ChatStatus     int64  `json:"chat_status"`
		State          string `json:"state"`
	}
)

// GetCustomerInfoWithCustomerAcquisitionLink 获客助手--获取由获客链接添加的客户信息
// see https://developer.work.weixin.qq.com/document/path/97298
func (r *Client) GetCustomerInfoWithCustomerAcquisitionLink(req *GetCustomerInfoWithCustomerAcquisitionLinkRequest) (*GetCustomerInfoWithCustomerAcquisitionLinkResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(getCustomerInfoWithCustomerAcquisitionLinkURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetCustomerInfoWithCustomerAcquisitionLinkResponse{}
	if err = util.DecodeWithError(response, result, "GetCustomerInfoWithCustomerAcquisitionLink"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// CustomerAcquisitionQuotaResponse 查询剩余使用量响应
	CustomerAcquisitionQuotaResponse struct {
		util.CommonError
		Total     int64       `json:"total"`
		Balance   int64       `json:"balance"`
		QuotaList []QuotaList `json:"quota_list"`
	}
	// QuotaList 额度
	QuotaList struct {
		ExpireDate int64 `json:"expire_date"`
		Balance    int64 `json:"balance"`
	}
)

// CustomerAcquisitionQuota 获客助手额度管理与使用统计--查询剩余使用量
// see https://developer.work.weixin.qq.com/document/path/97375
func (r *Client) CustomerAcquisitionQuota() (*CustomerAcquisitionQuotaResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.HTTPGet((fmt.Sprintf(customerAcquisitionQuotaURL, accessToken))); err != nil {
		return nil, err
	}
	result := &CustomerAcquisitionQuotaResponse{}
	if err = util.DecodeWithError(response, result, "CustomerAcquisitionQuota"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// CustomerAcquisitionStatisticRequest 查询链接使用详情请求
	CustomerAcquisitionStatisticRequest struct {
		LinkID    string `json:"link_id"`
		StartTime int64  `json:"start_time"`
		EndTime   int64  `json:"end_time"`
	}
	// CustomerAcquisitionStatisticResponse 查询链接使用详情响应
	CustomerAcquisitionStatisticResponse struct {
		util.CommonError
		ClickLinkCustomerCnt int64 `json:"click_link_customer_cnt"`
		NewCustomerCnt       int64 `json:"new_customer_cnt"`
	}
)

// CustomerAcquisitionStatistic 获客助手额度管理与使用统计--查询链接使用详情
// see https://developer.work.weixin.qq.com/document/path/97375
func (r *Client) CustomerAcquisitionStatistic(req *CustomerAcquisitionStatisticRequest) (*CustomerAcquisitionStatisticResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(customerAcquisitionStatisticURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &CustomerAcquisitionStatisticResponse{}
	if err = util.DecodeWithError(response, result, "CustomerAcquisitionStatistic"); err != nil {
		return nil, err
	}
	return result, nil
}
