package context

import (
	"sync"
	"testing"
)

func TestContext_SetCustomAccessTokenFunc(t *testing.T) {
	ctx := Context{
		accessTokenLock: new(sync.RWMutex),
	}
	f := func(ctx *Context) (accessToken string, err error) {
		return "fake token", nil
	}
	ctx.SetGetAccessTokenFunc(f)
	res, err := ctx.GetAccessToken()
	if res != "fake token" || err != nil {
		t.Error("expect fake token but error")
	}
}

func TestContext_NoSetCustomAccessTokenFunc(t *testing.T) {
	ctx := Context{
		accessTokenLock: new(sync.RWMutex),
	}

	if ctx.accessTokenFunc != nil {
		t.Error("error accessTokenFunc")
	}
}
