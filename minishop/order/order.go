package order

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/minishop/context"
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
	return o.FetchData(orderGetListURL, req)
}

type GetListRsp struct {
	Errcode int `json:"errcode"`
	Orders  []struct {
		OrderID     int64  `json:"order_id"`
		CreateTime  string `json:"create_time"`
		UpdateTime  string `json:"update_time"`
		Status      int    `json:"status"`
		OrderDetail struct {
			ProductInfos []struct {
				ProductID int    `json:"product_id"`
				SkuID     int    `json:"sku_id"`
				ThumbImg  string `json:"thumb_img"`
				SalePrice int    `json:"sale_price"`
				SkuCnt    int    `json:"sku_cnt"`
				Title     string `json:"title"`
				SkuAttrs  []struct {
					AttrKey   string `json:"attr_key"`
					AttrValue string `json:"attr_value"`
				} `json:"sku_attrs"`
				OnAftersaleSkuCnt     int `json:"on_aftersale_sku_cnt"`
				FinishAftersaleSkuCnt int `json:"finish_aftersale_sku_cnt"`
			} `json:"product_infos"`
			PayInfo struct {
				PayMethod     string `json:"pay_method"`
				PrepayID      string `json:"prepay_id"`
				PrepayTime    string `json:"prepay_time"`
				TransactionID string `json:"transaction_id"`
			} `json:"pay_info"`
			PriceInfo struct {
				ProductPrice int `json:"product_price"`
				OrderPrice   int `json:"order_price"`
				Freight      int `json:"freight"`
			} `json:"price_info"`
			DeliveryInfo struct {
				AddressInfo struct {
					UserName     string `json:"user_name"`
					PostalCode   string `json:"postal_code"`
					ProvinceName string `json:"province_name"`
					CityName     string `json:"city_name"`
					CountyName   string `json:"county_name"`
					DetailInfo   string `json:"detail_info"`
					NationalCode string `json:"national_code"`
					TelNumber    string `json:"tel_number"`
				} `json:"address_info"`
				DeliveryMethod string `json:"delivery_method"`
				ExpressFee     []struct {
					Result int `json:"result"`
				} `json:"express_fee"`
				DeliveryProductInfo []interface{} `json:"delivery_product_info"`
			} `json:"delivery_info"`
		} `json:"order_detail"`
		AftersaleDetail struct {
			AftersaleOrderList  []interface{} `json:"aftersale_order_list"`
			OnAftersaleOrderCnt int           `json:"on_aftersale_order_cnt"`
		} `json:"aftersale_detail"`
		Openid  string `json:"openid"`
		ExtInfo struct {
			CustomerNotes string `json:"customer_notes"`
			MerchantNotes string `json:"merchant_notes"`
		} `json:"ext_info"`
	} `json:"orders"`
	TotalNum int `json:"total_num"`
}

func (o *Order) ParseOrderList(t []byte) (*GetListRsp, error) {
	rsp := &GetListRsp{}
	err := json.Unmarshal(t, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// Get 获取订单详情
func (o *Order) Get(orderID string) ([]byte, error) {
	req := map[string]string{
		"order_id": orderID,
	}
	return o.FetchData(orderGetURL, req)
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
	return o.FetchData(orderSearchURL, req)
}
