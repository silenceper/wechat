package coupon

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	couponGetListURL = "https://api.weixin.qq.com/product/coupon/get_list?access_token=%s"
	couponPushURL    = "https://api.weixin.qq.com/product/coupon/push?access_token=%s"
)

const (
	// GOODSDISCOUNT 商品条件折扣券
	GOODSDISCOUNT CStatus = 1
	// GOODSFULLDISCOUNT 商品满减券
	GOODSFULLDISCOUNT CStatus = 2
	// GOODSALLDISCOUNT 商品统一折扣券
	GOODSALLDISCOUNT CStatus = 3
	// GOODSDDISCOUNT 商品直减券
	GOODSDDISCOUNT CStatus = 4
	// SHOPSDISCOUNT 店铺条件折扣券
	SHOPSDISCOUNT CStatus = 101
	// SHOPSFULLDISCOUNT 店铺满减券
	SHOPSFULLDISCOUNT CStatus = 102
	// SHOPSALLDISCOUNT 店铺统一折扣券
	SHOPSALLDISCOUNT CStatus = 103
	// SHOPSDDISCOUNT 店铺直减券
	SHOPSDDISCOUNT CStatus = 104
)

// CStatus 折扣券类型
type CStatus int

//Coupon 优惠券接口
type Coupon struct {
	*context.Context
}

//NewCoupon new order
func NewCoupon(ctx *context.Context) *Coupon {
	return &Coupon{ctx}
}

// fetchData
func (c *Coupon) fetchData(urlStr string, body interface{}) (response []byte, err error) {
	accessToken, err := c.AccessTokenHandle.GetAccessToken()
	if err != nil {
		return nil, err
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)

	v := url.Values{}
	if c.Config.ServiceID != "" {
		v.Add("service_id", c.ServiceID)
	}
	if c.Config.SpecificationID != "" {
		v.Add("specification_id", c.SpecificationID)
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

// GetListParam 请求优惠券列表参数
type GetListParam struct {
	StartCreateTime string  `json:"start_create_time"` // (必填)优惠券创建时间的搜索开始时间
	EndCreateTime   string  `json:"end_create_time"`   // (必填)优惠券创建时间的搜索结束时间
	Status          CStatus `json:"status"`            // (必填)优惠券状态
	Page            int     `json:"page"`              // (必填)第几页（最小填1）
	PageSize        int     `json:"page_size"`         // (必填)每页数量(不超过10,000)
}

// GetListResp 获取优惠券列表
type GetListResp struct {
	util.CommonError
	Coupons []struct {
		CouponID   int    `json:"coupon_id"`
		Type       int    `json:"type"`
		Status     int    `json:"status"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
		CouponInfo struct {
			Name      string `json:"name"`
			ValidInfo struct {
				ValidType   int    `json:"valid_type"`
				ValidDayNum int    `json:"valid_day_num"`
				StartTime   string `json:"start_time"`
				EndTime     string `json:"end_time"`
			} `json:"valid_info"`
		} `json:"coupon_info"`
		StockInfo struct {
			IssuedNum  int `json:"issued_num"`
			ReceiveNum int `json:"receive_num"`
			UsedNum    int `json:"used_num"`
		} `json:"stock_info"`
	} `json:"coupons"`
}

// GetList 获取优惠券列表
func (c *Coupon) GetList(req *GetListParam) (*GetListResp, error) {
	info := &GetListResp{}
	response, err := s.fetchData(couponGetListURL, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, err
}

// Push 发放优惠券
func (c *Coupon) Push(openid, coupon_id string) error {
	req := map[string]string{
		"openid":    openid,
		"coupon_id": coupon_id,
	}
	_, err := c.fetchData(orderGetURL, req)
	return err
}
