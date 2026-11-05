package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingCreateCustomerShortURLURL 创建用户专属参会链接
	meetingCreateCustomerShortURLURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/create_customer_short_url?access_token=%s"
	// meetingGetCustomerShortURLURL 获取用户专属参会链接
	meetingGetCustomerShortURLURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_customer_short_url?access_token=%s"
)

type (
	// CreateCustomerShortURLRequest 创建用户专属参会链接请求
	CreateCustomerShortURLRequest struct {
		MeetingID    string `json:"meetingid"`
		CustomerData string `json:"customer_data"`
	}

	// CreateCustomerShortURLResponse 创建用户专属参会链接响应
	CreateCustomerShortURLResponse struct {
		util.CommonError
		MeetingShortURLCustomerData *CustomerShortURLData `json:"meeting_short_url_customer_data"`
	}

	// GetCustomerShortURLRequest 获取用户专属参会链接请求
	GetCustomerShortURLRequest struct {
		MeetingID string `json:"meetingid"`
	}

	// GetCustomerShortURLResponse 获取用户专属参会链接响应
	GetCustomerShortURLResponse struct {
		util.CommonError
		MeetingShortURLCustomerDataList []*CustomerShortURLData `json:"meeting_short_url_customer_data_list"`
	}

	// CustomerShortURLData 用户专属参会链接数据
	CustomerShortURLData struct {
		CustomerData    string `json:"customer_data"`
		MeetingShortURL string `json:"meeting_short_url"`
	}
)

// CreateCustomerShortURL 创建用户专属参会链接
// see https://developer.work.weixin.qq.com/document/path/98818
func (r *Client) CreateCustomerShortURL(req *CreateCustomerShortURLRequest) (*CreateCustomerShortURLResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingCreateCustomerShortURLURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &CreateCustomerShortURLResponse{}
	err = util.DecodeWithError(response, result, "CreateCustomerShortURL")
	return result, err
}

// GetCustomerShortURL 获取用户专属参会链接
// see https://developer.work.weixin.qq.com/document/path/98819
func (r *Client) GetCustomerShortURL(req *GetCustomerShortURLRequest) (*GetCustomerShortURLResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetCustomerShortURLURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetCustomerShortURLResponse{}
	err = util.DecodeWithError(response, result, "GetCustomerShortURL")
	return result, err
}
