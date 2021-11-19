package privacyconfig

import "github.com/silenceper/wechat/v2/util"

type GetPrivacySettingResponse struct {
	util.CommonError
	CodeExist    int8            `json:"code_exist"`    //代码是否存在， 0 不存在， 1 存在 。如果最近没有通过commit接口上传代码，则会出现 code_exist=0的情况。
	PrivacyList  []string        `json:"privacy_list"`  //代码检测出来的用户信息类型（privacy_key）
	SettingList  []*SettingItem  `json:"setting_list"`  //要收集的用户信息配置
	UpdateTime   int64           `json:"update_time"`   //更新时间
	OwnerSetting *OwnerSetting   `json:"owner_setting"` //收集方（开发者）信息配置
	PrivacyDesc  PrivacyDescList `json:"privacy_desc"`  //用户信息类型对应的中英文描述
}

type SettingItem struct {
	PrivacyKey   string `json:"privacy_key"`             //用户信息类型的英文名称
	PrivacyText  string `json:"privacy_text"`            //该用户信息类型的用途
	PrivacyLabel string `json:"privacy_label,omitempty"` //用户信息类型的中文名称
}

type OwnerSetting struct {
	ContactEmail         string `json:"contact_email"`          //信息收集方（开发者）的邮箱
	ContactPhone         string `json:"contact_phone"`          //信息收集方（开发者）的手机号
	ContactQq            string `json:"contact_qq"`             //信息收集方（开发者）的qq
	ContactWeixin        string `json:"contact_weixin"`         //信息收集方（开发者）的微信号
	NoticeMethod         string `json:"notice_method"`          //通知方式，指的是当开发者收集信息有变动时，通过该方式通知用户
	StoreExpireTimestamp string `json:"store_expire_timestamp"` //存储期限，指的是开发者收集用户信息存储多久
	ExtFileMediaId       string `json:"ext_file_media_id"`      //自定义 用户隐私保护指引文件的media_id
}

type PrivacyDescList struct {
	PrivacyDescList []*PrivacyDescItem `json:"privacy_desc_list"`
}

type PrivacyDescItem struct {
	PrivacyKey  string `json:"privacy_key"`  //用户信息类型的英文key
	PrivacyDesc string `json:"privacy_desc"` //用户信息类型的中文描述
}

type SetPrivacySettingParams struct {
	PrivacyVer   int8           `json:"privacy_ver"`            //用户隐私保护指引的版本，1表示现网版本；2表示开发版
	OwnerSetting *OwnerSetting  `json:"owner_setting"`          //收集方（开发者）信息配置
	SettingList  []*SettingItem `json:"setting_list,omitempty"` //要收集的用户信息配置，可选择的用户信息类型参考下方详情。当privacy_ver传2或者不传是必填；当privacy_ver传1时，该参数不可传，否则会报错
}
