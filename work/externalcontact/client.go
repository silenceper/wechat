package externalcontact

import (
	"errors"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/context"
)

// Client 实例
type Client struct {
	corpID string // 企业ID：企业开通的每个微信客服，都对应唯一的企业ID，企业可在微信客服管理后台的企业信息处查看
	cache  cache.Cache
	ctx    *context.Context
}

// NewClient
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, errors.New("初始化失败")
	}

	// 初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Client{
		corpID: cfg.CorpID,
		cache:  cfg.Cache,
		ctx:    ctx,
	}

	return client, nil
}
