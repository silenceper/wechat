package credential

//AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}
