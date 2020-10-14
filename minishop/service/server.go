package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	serverGetList      = "https://api.weixin.qq.com/product/service/get_list?access_token=%s"
	serverGetOrderList = "https://api.weixin.qq.com/product/service/get_order_list?access_token=%s"
)

//Service 服务商接口
type Service struct {
	*context.Context
}

//NewService new order
func NewService(ctx *context.Context) *Service {
	return &Service{ctx}
}

// fetchData
func (s *Service) fetchData(urlStr string, body interface{}) (response []byte, err error) {
	accessToken, err := s.AccessTokenHandle.GetAccessToken()
	if err != nil {
		return nil, err
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)

	response, err = util.PostJSON(urlStr, body)
	if err != nil {
		return
	}
	// 返回错误信息
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err == nil && result.ErrCode != 0 {
		err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return response, err
}

// GetOrderListRsp 用户购买的服务列表
type GetOrderListRsp struct {
	util.CommonError
	ServiceOrderList []struct {
		ServiceOrderID  int    `json:"service_order_id"`
		ServiceID       int    `json:"service_id"`
		ServiceName     string `json:"service_name"`
		CreateTime      string `json:"create_time"`
		ExpireTime      string `json:"expire_time"`
		ServiceType     int    `json:"service_type"`
		SpecificationID string `json:"specification_id"`
		TotalPrice      int    `json:"total_price"`
	} `json:"service_order_list"`
}

// GetOrderList 获取用户购买的服务列表
func (s *Service) GetOrderList(start, end string, page, pageSize int) (*GetOrderListRsp, error) {
	info := &GetOrderListRsp{}
	req := map[string]string{
		"start_create_time": start,
		"end_create_time":   end,
		"page":              strconv.Itoa(page),
		"page_size":         strconv.Itoa(pageSize),
	}
	response, err := s.fetchData(serverGetOrderList, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetListRsp 用户购买的在有效期内的服务列表
type GetListRsp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	ServiceList []struct {
		ServiceID   int    `json:"service_id"`
		ServiceName string `json:"service_name"`
		ExpireTime  string `json:"expire_time"`
		ServiceType int    `json:"service_type"`
	} `json:"service_list"`
}

// GetList 获取订单详情
func (s *Service) GetList() (*GetListRsp, error) {
	info := &GetListRsp{}
	req := map[string]string{}

	response, err := s.fetchData(serverGetList, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
