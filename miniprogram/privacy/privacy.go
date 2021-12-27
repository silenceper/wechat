package privacy

import "github.com/silenceper/wechat/v2/miniprogram/context"

// Privacy 小程序授权隐私设置
type Privacy struct {
	*context.Context
}

// NewCustomerMessageManager 实例化消息管理者
func NewPrivacy(context *context.Context) *Privacy {
	return &Privacy{
		context,
	}
}

type OwnerSetting struct {
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	ContactQQ string `json:"contact_qq"`
	ContactWeixin string `json:"contact_weixin"`
	ExtFileMediaID string `json:"ext_file_media_id"`
	NoticeMethod string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}
