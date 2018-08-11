package menu

//Button 菜单按钮
type Button struct {
	Type       string    `json:"type,omitempty"`
	Name       string    `json:"name,omitempty"`
	Key        string    `json:"key,omitempty"`
	URL        string    `json:"url,omitempty"`
	MediaID    string    `json:"media_id,omitempty"`
	APPID        string    `json:"appid,omitempty"`
	PagePath        string    `json:"pagepath,omitempty"`
	SubButtons []*Button `json:"sub_button,omitempty"`
}

//SetSubButton 设置二级菜单
func (btn *Button) SetSubButton(name string, subButtons []*Button) {
	btn.Name = name
	btn.SubButtons = subButtons
	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
	btn.MediaID = ""
}

//SetClickButton btn 为click类型
func (btn *Button) SetClickButton(name, key string) {
	btn.Type = "click"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetViewButton view类型
func (btn *Button) SetViewButton(name, url string) {
	btn.Type = "view"
	btn.Name = name
	btn.URL = url
	btn.Key = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetScanCodePushButton 扫码推事件
func (btn *Button) SetScanCodePushButton(name, key string) {
	btn.Type = "scancode_push"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetScanCodeWaitMsgButton 设置 扫码推事件且弹出"消息接收中"提示框
func (btn *Button) SetScanCodeWaitMsgButton(name, key string) {
	btn.Type = "scancode_waitmsg"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetPicSysPhotoButton 设置弹出系统拍照发图按钮
func (btn *Button) SetPicSysPhotoButton(name, key string) {
	btn.Type = "pic_sysphoto"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetPicPhotoOrAlbumButton 设置弹出拍照或者相册发图类型按钮
func (btn *Button) SetPicPhotoOrAlbumButton(name, key string) {
	btn.Type = "pic_photo_or_album"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetPicWeixinButton 设置弹出微信相册发图器类型按钮
func (btn *Button) SetPicWeixinButton(name, key string) {
	btn.Type = "pic_weixin"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

// SetLocationSelectButton 设置 弹出地理位置选择器 类型按钮
func (btn *Button) SetLocationSelectButton(name, key string) {
	btn.Type = "location_select"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
}

//SetMediaIDButton  设置 下发消息(除文本消息) 类型按钮
func (btn *Button) SetMediaIDButton(name, mediaID string) {
	btn.Type = "media_id"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}

//SetViewLimitedButton  设置 跳转图文消息URL 类型按钮
func (btn *Button) SetViewLimitedButton(name, mediaID string) {
	btn.Type = "view_limited"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}
