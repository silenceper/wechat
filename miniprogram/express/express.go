package express

import "github.com/silenceper/wechat/v2/miniprogram/content"

//Express 物流助手
type Express struct {
	*content.Content
}

//NewExpress 物流助手
func NewExpress(ctx *content.Content) *Express {
	return &Express{ctx}
}

//ListProviders 获取支持的快递公司列表
func (exp *Express) ListProviders() {

}
