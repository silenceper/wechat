package open

const (
	tmplCategoryURL         = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory"
	getpubtemplatetitlesURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles"
	addtemplateURL          = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"
	gettemplateURL          = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
	deltemplateURL          = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"
)

// GetTmplCategory 获取小程序支持的类目
func (m *MiniPrograms) GetTmplCategory() (ret []byte, err error) {
	_ = tmplCategoryURL
	return
}

// GetPubTemplateTitles 获取公共区域模板标题
func (m *MiniPrograms) GetPubSubscribeTemplateTitles() (ret []byte, err error) {
	_ = getpubtemplatetitlesURL
	return
}

// AddSubscribeTemplate 添加订阅模板到用户身上
func (m *MiniPrograms) AddSubscribeTemplate() (ret []byte, err error) {
	_ = addtemplateURL
	return
}

// GetSubscribeTemplate 获取用户订阅模板列表
func (m *MiniPrograms) GetSubscribeTemplate() (ret []byte, err error) {
	_ = gettemplateURL
	return
}

// DelSubscribeTemplate 删除订阅模板
func (m *MiniPrograms) DelSubscribeTemplate() (ret []byte, err error) {
	_ = deltemplateURL
	return
}
