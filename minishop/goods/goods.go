package goods

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	spuGetListURL = "https://api.weixin.qq.com/product/spu/get_list?access_token=%s"
	spuGetURL     = "https://api.weixin.qq.com/product/spu/get?access_token=%s"
)

//Goods 商品接口
type Goods struct {
	*context.Context
}

//NewGoods new goods
func NewGoods(ctx *context.Context) *Goods {
	return &Goods{ctx}
}

// GetListRsp 商品列表
type GetListRsp struct {
	util.CommonError
	Data struct {
		Spu struct {
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
			MinPrice     int    `json:"min_price"`
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
		} `json:"spu"`
	} `json:"data"`
}

// GetList 获取商品列表
func (g *Goods) GetList(status, page, pageSize int) (*GetListRsp, error) {
	response, err := g.GetListByte(status, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &GetListRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetListByte 获取商品列表
func (g *Goods) GetListByte(status, page, pageSize int) ([]byte, error) {
	req := map[string]string{
		"status":    strconv.Itoa(status),
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	return g.request(spuGetListURL, req)
}

func (g *Goods) request(reqURL string, req map[string]string) ([]byte, error) {
	return g.FetchData(reqURL, req)
}
