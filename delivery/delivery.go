package delivery

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	addOrderURL        = "https://api.weixin.qq.com/cgi-bin/express/business/order/add"
	testUpdateOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/test_update_order"
	delMaterialURL     = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	getMaterialURL     = "https://api.weixin.qq.com/cgi-bin/material/get_material"
)

//Delivery 素材管理
type Delivery struct {
	*context.Context
}

//NewDelivery init
func NewDelivery(context *context.Context) *Delivery {
	delivery := new(Delivery)
	delivery.Context = context
	return delivery
}

//AddOrderFields 添加订单
type AddOrderFields struct {
	AccessToken  string   `json:"access_token"`  //必须	接口调用凭证
	AddSource    int      `json:"add_source"`    //必须 订单来源，0为小程序订单，2为App或H5订单，填2则不发送物流服务通知
	WxAppid      string   `json:"wx_appid"`      //必须 App或H5的appid，add_source=2时必填，需和开通了物流助手的小程序绑定同一open帐号
	OrderID      string   `json:"order_id"`      //必须 订单ID，须保证全局唯一，不超过512字节
	Openid       string   `json:"openid"`        //否 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID   string   `json:"delivery_id"`   //是	快递公司ID，参见getAllDelivery
	BizID        string   `json:"biz_id"`        //是	快递客户编码或者现付编码
	CustomRemark string   `json:"custom_remark"` //否	快递备注信息，比如"易碎物品"，不超过1024字节
	Tagid        int      `json:"tagid"`         //tagid	number		否	订单标签id，用于平台型小程序区分平台上的入驻方，tagid须与入驻方账号一一对应，非平台型小程序无需填写该字段
	Sender       Sender   `json:"sender"`        //sender	Object		是	发件人信息
	Receiver     Receiver `json:"receiver"`      //	Object		是	收件人信息
	Cargo        Cargo    `json:"cargo"`         //是	包裹信息，将传递给快递公司
	Shop         Shop     `json:"shop"`          //	是	商品信息，会展示到物流服务通知和电子面单中
	Insured      Insured  `json:"insured"`       //是	保价信息
	Service      Service  `json:"service"`       //是	服务类型
	ExpectTime   int64    `json:"expect_time"`   //否	Unix 时间戳, 单位秒，顺丰必须传。 预期的上门揽件时间，0表示已事先约定取件时间；否则请传预期揽件时间戳，需大于当前时间，收件员会在预期时间附近上门。例如expect_time为“1557989929”，表示希望收件员将在2019年05月16日14:58:49-15:58:49内上门取货。说明：若选择 了预期揽件时间，请不要自己打单，由上门揽件的时候打印。如果是下顺丰散单，则必传此字段，否则不会有收件员上门揽件。
}

//Sender 发件人信息
type Sender struct {
	Name     string `json:"name"`      //是	发件人姓名，不超过64字节
	Tel      string `json:"tel"`       //否	发件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile"`    //否	发件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company"`   //否	发件人公司名称，不超过64字节
	PostCode string `json:"post_code"` //否	发件人邮编，不超过10字节
	Country  string `json:"Country"`   //否	发件人国家，不超过64字节
	Province string `json:"province"`  //是	发件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`      //是	发件人市/地区，比如："广州市"，不超过64字节
	Area     string `json:"area"`      //是	发件人区/县，比如："海珠区"，不超过64字节
	Address  string `json:"address"`   //是	发件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

//Receiver 收件人信息
type Receiver struct {
	Name     string `json:"name"`      //是	收件人姓名，不超过64字节
	Tel      string `json:"tel"`       //否	收件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile"`    //否	收件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company"`   //否	收件人公司名，不超过64字节
	PostCode string `json:"post_code"` //否	收件人邮编，不超过10字节
	Country  string `json:"country"`   //否	收件人所在国家，不超过64字节
	Province string `json:"province"`  //是	收件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`      //是	收件人地区/市，比如："广州市"，不超过64字节
	Area     string `json:"area"`      //是	收件人区/县，比如："天河区"，不超过64字节
	Address  string `json:"address"`   //是	收件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

//Cargo 包裹信息，将传递给快递公司
type Cargo struct {
	Count      int          `json:"count"`       //是	包裹数量, 需要和detail_list size保持一致
	Weight     int          `json:"weight"`      //是	包裹总重量，单位是千克(kg)
	SpaceX     int          `json:"space_x"`     //是	包裹长度，单位厘米(cm)
	SpaceY     int          `json:"space_y"`     //是	包裹宽度，单位厘米(cm)
	SpaceZ     int          `json:"space_z"`     //是	包裹高度，单位厘米(cm)
	DetailList []DetailList `json:"detail_list"` //是	包裹中商品详情列表
}

//DetailList 包裹中商品详情列表
type DetailList struct {
	Name  string `json:"name"`  //是	商品名，不超过128字节
	Count int    `json:"count"` //是	商品数量
}

//Shop 商品信息，会展示到物流服务通知和电子面单中
type Shop struct {
	WxaPath    string `json:"wxa_path"`    //是	商家小程序的路径，建议为订单页面
	ImgURL     string `json:"img_url"`     //是	商品缩略图 url
	GoodsName  string `json:"goods_name"`  //是	商品名称, 不超过128字节
	GoodsCount int    `json:"goods_count"` //是	商品数量
}

//Insured 保价信息
type Insured struct {
	UseInsured   int `json:"use_insured"`   //是	是否保价，0 表示不保价，1 表示保价
	InsuredValue int `json:"insured_value"` //是	保价金额，单位是分，比如: 10000 表示 100 元
}

//Service 服务类型
type Service struct {
	ServiceType int    `json:"service_type"` //是	服务类型ID，详见已经支持的快递公司基本信息
	ServiceName string `json:"service_name"` //是	服务名称，详见已经支持的快递公司基本信息
}

//AddOrderReturn 接口返回信息
type AddOrderReturn struct {
	util.CommonError
	OrderID            string        `json:"order_id"`            //订单ID，下单成功时返回
	WaybillID          string        `json:"waybill_id"`          //运单ID，下单成功时返回
	WaybillData        []WaybillData `json:"waybill_data"`        //运单信息，下单成功时返回
	DeliveryResultcode int           `json:"delivery_resultcode"` //快递侧错误码，下单失败时返回
	DeliveryResultmsg  string        `json:"delivery_resultmsg"`  //快递侧错误信息，下单失败时返回
}

//WaybillData 运单信息，下单成功时返回
type WaybillData struct {
	Key   string `json:"key"`   //运单信息 key
	Value string `json:"value"` //运单信息 value
}

//AddOrder 生成订单
func (delivery *Delivery) AddOrder(req *AddOrderFields) (*AddOrderReturn, error) {
	accessToken, err := delivery.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", addOrderURL, accessToken)

	responseBytes, err := util.PostJSON(uri, req)

	var rea *AddOrderReturn
	err = json.Unmarshal(responseBytes, &rea)
	if err != nil {
		return nil, err
	}

	return rea, nil
}

//TestUpdateOrder 模拟快递公司更新订单状态
type TestUpdateOrder struct {
	BizID      string `json:"biz_id"`      //是	商户id,需填test_biz_id
	OrderID    string `json:"order_id"`    //是	订单号
	DeliveryID string `json:"delivery_id"` //是	快递公司id,需填TEST
	WaybillID  string `json:"waybill_id"`  //是	运单号
	ActionTime int64  `json:"action_time"` //是	轨迹变化 Unix 时间戳
	ActionType int64  `json:"action_type"` //是	轨迹变化类型
	ActionMsg  string `json:"action_msg"`  //是	轨迹变化具体信息说明,使用UTF-8编码
}

//TestUpdateOrderReturn 返回信息
type TestUpdateOrderReturn struct {
	util.CommonError
}

//TestUpdateOrder 模拟快递公司更新订单状态
func (delivery *Delivery) TestUpdateOrder(req TestUpdateOrder) (*TestUpdateOrderReturn, error) {
	accessToken, err := delivery.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", testUpdateOrderURL, accessToken)

	responseBytes, err := util.PostJSON(uri, req)

	var rea *TestUpdateOrderReturn
	err = json.Unmarshal(responseBytes, &rea)
	if err != nil {
		return nil, err
	}

	return rea, nil
}
