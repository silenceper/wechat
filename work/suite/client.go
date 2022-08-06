package suite

// Client 第三方应用实例
type Client struct {
	SuiteAccessToken string
}

// NewClient 初始化实例
func NewClient(suiteAccessToken string) *Client {
	return &Client{
		SuiteAccessToken: suiteAccessToken,
	}
}
