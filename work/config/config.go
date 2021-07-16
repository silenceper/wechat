// Package config 企业微信config配置
package config

import (
	"github.com/silenceper/wechat/v2/cache"
)

// Config config for 企业微信
type Config struct {
	CorpID     string `json:"corp_id"`     // corp_id
	CorpSecret string `json:"corp_secret"` // corp_secret
	AgentID    string `json:"agent_id"`    // agent_id
	Cache      cache.Cache
}
