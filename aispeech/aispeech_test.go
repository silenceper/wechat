package aispeech

import (
	stdcontext "context"
	"testing"

	"github.com/silenceper/wechat/v2/aispeech/config"
	"github.com/silenceper/wechat/v2/cache"
)

type staticAccessToken struct {
	token string
}

func (s staticAccessToken) GetAccessToken() (string, error) {
	return s.token, nil
}

type staticAccessTokenContext struct {
	token string
}

type contextTokenKey struct{}

func (s staticAccessTokenContext) GetAccessToken() (string, error) {
	return s.GetAccessTokenContext(stdcontext.Background())
}

func (s staticAccessTokenContext) GetAccessTokenContext(ctx stdcontext.Context) (string, error) {
	if v := ctx.Value(contextTokenKey{}); v != nil {
		return v.(string), nil
	}
	return s.token, nil
}

func TestSetAccessTokenHandle(t *testing.T) {
	ai := NewAISpeech(&config.Config{Cache: cache.NewMemory()})
	ai.SetAccessTokenHandle(staticAccessToken{token: "custom-token"})

	token, err := ai.GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken error: %v", err)
	}
	if token != "custom-token" {
		t.Fatalf("bad token: %s", token)
	}
}

func TestSetAccessTokenContextHandle(t *testing.T) {
	ai := NewAISpeech(&config.Config{Cache: cache.NewMemory()})
	ai.SetAccessTokenContextHandle(staticAccessTokenContext{token: "custom-token"})

	ctx := stdcontext.WithValue(stdcontext.Background(), contextTokenKey{}, "context-token")
	token, err := ai.GetAccessTokenContext(ctx)
	if err != nil {
		t.Fatalf("GetAccessTokenContext error: %v", err)
	}
	if token != "context-token" {
		t.Fatalf("bad token: %s", token)
	}
}
