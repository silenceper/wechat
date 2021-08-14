package kf

import (
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/work/config"

	"time"
)

// Client 微信客服实例
type Client struct {
	corpID         string        // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	secret         string        // Secret是微信客服用于校验开发者身份的访问密钥，企业成功注册微信客服后，可在「微信客服管理后台-开发配置」处获取
	token          string        // 用于生成签名校验回调请求的合法性
	encodingAESKey string        // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	expireTime     time.Duration // 令牌过期时间
	cache          cache.Cache
	accessToken    string // 用户访问凭证
	isCloseCache   bool   // 是否自动缓存AccessToken, 默认缓存
}

// NewClient 初始化微信客服实例
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, NewSDKErr(50001)
	}

	client = &Client{
		corpID:         cfg.CorpID,
		secret:         cfg.CorpSecret,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
	}

	// token默认有效期是7200
	client.expireTime = 6000 * time.Second

	if err = client.initAccessToken(); err != nil {
		return nil, err
	}

	return client, nil
}
