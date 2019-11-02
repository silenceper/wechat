package open

const (
	// 草稿箱列表
	TEMPLATE_DRAFT_LIST_URL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token=%s"
)


func (o *Open) TplDraftList() (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	precode, err := o.GetPreCode()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(COMPONENT_LOGIN_PAGE_URL, o.AppID, precode, urlStr, authType), nil
}
