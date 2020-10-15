package order

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	spuGetListURL = "https://api.weixin.qq.com/product/spu/get_list?access_token=%s"
)

//Goods 商品接口
type Goods struct {
	*context.Context
}

//NewGoods new goods
func NewGoods(ctx *context.Context) *Goods {
	return &Goods{ctx}
}

// fetchData
func (g *Goods) fetchData(urlStr string, body interface{}) (response []byte, err error) {
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

// GetListRsp 商品列表
type GetListRsp struct {
	util.CommonError
	Spus []struct {
		Title    string   `json:"title"`
		SubTitle string   `json:"sub_title"`
		HeadImg  []string `json:"head_img"`
		DescInfo struct {
			Imgs []string `json:"imgs"`
		} `json:"desc_info"`
		OutProductID string `json:"out_product_id"`
		ProductID    int    `json:"product_id"`
		BrandID      int    `json:"brand_id"`
		Status       int    `json:"status"`
		EditStatus   int    `json:"edit_status"`
		Cats         []struct {
			CatID int `json:"cat_id"`
			Level int `json:"level"`
		} `json:"cats"`
		Attrs []struct {
			AttrKey   string `json:"attr_key"`
			AttrValue string `json:"attr_value"`
		} `json:"attrs"`
		Model   string `json:"model"`
		Shopcat []struct {
			ShopcatID int `json:"shopcat_id"`
		} `json:"shopcat"`
		Skus []struct {
			SkuID int `json:"sku_id"`
		} `json:"skus"`
	} `json:"spus"`
}

// GetList 获取商品列表
func (o *Order) GetList(req *GetListParam) (*GetListRsp, error) {
	info := &GetListRsp{}
	response, err := s.fetchData(spuGetListURL, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
