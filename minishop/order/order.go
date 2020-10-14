package order

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	orderGetListURL = "https://api.weixin.qq.com/product/order/get_list?access_token=%s"
	orderGetURL     = "https://api.weixin.qq.com/product/order/get?access_token=%s"
	orderSearchURL  = "https://api.weixin.qq.com/product/order/search?access_token=%s"
)

const (
	// UNPAY 未支付
	UNPAY OStatus = 10
	// UNDELIVE 待发货
	UNDELIVE OStatus = 20
	// UNRECEVIED 待收货
	UNRECEVIED OStatus = 30
	// COMPLETE 完成
	COMPLETE OStatus = 100
	// CANCEL 全部商品售后之后，订单取消
	CANCEL OStatus = 200
	// OVERTIME 用户主动取消或待付款超时取消
	OVERTIME OStatus = 250
)

// OStatus 订单状态
type OStatus int

//Order 订单接口
type Order struct {
	*context.Context
}

//NewOrder new order
func NewOrder(ctx *context.Context) *Order {
	return &Order{ctx}
}

// fetchData
func (o *Order) fetchData(urlStr string, body interface{}) (response []byte, err error) {
	accessToken, err := o.AccessTokenHandle.GetAccessToken()
	if err != nil {
		return nil, err
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)

	v := url.Values{}

	if o.Config.ServiceID != "" {
		v.Add("service_id", o.Config.ServiceID)
	}
	if o.Config.SpecificationID != "" {
		v.Add("specification_id", o.Config.SpecificationID)
	}
	encode := v.Encode()
	if encode != "" {
		urlStr = urlStr + "&" + encode
	}

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

// GetListParam 请求订单列表参数
type GetListParam struct {
	StartCreateTime string  `json:"start_create_time"` // 订单创建时间的搜索开始时间
	EndCreateTime   string  `json:"end_create_time"`   // 订单创建时间的搜索结束时间
	StartUpdateTime string  `json:"start_update_time"` // 订单更新时间的搜索开始时间
	EndUpdateTime   string  `json:"end_update_time"`   // 订单更新时间的搜索结束时间
	Status          OStatus `json:"status"`            // (必填)订单状态
	Page            int     `json:"page"`              // (必填)第几页（最小填1）
	PageSize        int     `json:"page_size"`         // (必填)每页数量(不超过10,000)
}

// GetList 获取订单列表
func (o *Order) GetList(req *GetListParam) ([]byte, error) {
	return o.fetchData(orderGetListURL, req)
}

// Get 获取订单详情
func (o *Order) Get(orderID string) ([]byte, error) {
	req := map[string]string{
		"order_id": orderID,
	}
	return o.fetchData(orderGetURL, req)
}

// SearchParam 搜索订单请求参数
type SearchParam struct {
	StartPayTime          string `json:"start_pay_time"`           // 订单支付时间的搜索开始时间
	EndPayTime            string `json:"end_pay_time"`             // 订单支付时间的搜索结束时间
	Title                 string `json:"title"`                    // 商品标题关键词
	SkuCode               string `json:"sku_code"`                 // 商品编码
	UserName              string `json:"user_name"`                // 收件人
	TelNumber             string `json:"tel_number"`               // 收件人电话
	OnAftersaleOrderExist bool   `json:"on_aftersale_order_exist"` // 全部订单 0:没有正在售后的订单, 1:正在售后单数量大于等于1的订单
	Status                int    `json:"status"`                   // (必填)订单状态
	Page                  int    `json:"page"`                     // (必填)第几页（最小填1）
	PageSize              int    `json:"page_size"`                // (必填)每页数量(不超过10,000)
}

// Search 搜索订单
func (o *Order) Search(req SearchParam) ([]byte, error) {
	return o.fetchData(orderSearchURL, req)
}
