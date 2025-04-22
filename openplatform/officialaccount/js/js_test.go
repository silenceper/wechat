// 验证 js.GetConfigContext 是否能正确传递上下文到 HTTP 请求，确保上下文正确传播，防止在获取 JSSDK 配置时发生协程泄露。
package js

import (
	"bytes"
	context2 "context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

// mockAccessTokenHandle 模拟 AccessTokenHandle
type mockAccessTokenHandle struct{}

func (m *mockAccessTokenHandle) GetAccessToken() (string, error) {
	return "mock-access-token", nil
}

func (m *mockAccessTokenHandle) GetAccessTokenContext(_ context2.Context) (string, error) {
	return "mock-access-token", nil
}

// contextCheckingRoundTripper 自定义 RoundTripper 用于检查 context
type contextCheckingRoundTripper struct {
	originalCtx context2.Context
	t           *testing.T
	key         interface{}
	expectedVal interface{}
}

func (rt *contextCheckingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// 获取请求中的 context
	reqCtx := req.Context()

	// 打印 context 比较结果
	rt.t.Logf("比较上下文的内存地址:\n")
	if reqCtx == rt.originalCtx {
		rt.t.Logf("上下文具有相同的内存地址。原始上下文: %p, 请求上下文: %p\n", rt.originalCtx, reqCtx)
	} else {
		rt.t.Logf("上下文具有不同的内存地址。原始上下文: %p, 请求上下文: %p\n", rt.originalCtx, reqCtx)
	}

	// 检查 context 中的键值对
	if rt.key != nil {
		value := reqCtx.Value(rt.key)
		rt.t.Logf("检查请求上下文中的键 %v:\n", rt.key)
		if value != rt.expectedVal {
			rt.t.Errorf("上下文键 %v 的值不匹配: 预期 %v, 实际 %v\n", rt.key, rt.expectedVal, value)
		} else {
			rt.t.Logf("上下文键 %v 的值匹配: 预期 %v, 实际 %v\n", rt.key, rt.expectedVal, value)
		}
	}

	// 检查上下文是否已取消
	select {
	case <-reqCtx.Done():
		return nil, reqCtx.Err() // 返回上下文取消错误
	default:
		// 返回模拟的 HTTP 响应，包含有效的 JSON
		responseBody := `{"ticket":"mock-ticket","expires_in":7200}`
		response := &http.Response{
			Status:        "200 OK",
			StatusCode:    http.StatusOK,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          io.NopCloser(bytes.NewReader([]byte(responseBody))),
			ContentLength: int64(len(responseBody)),
			Header:        make(http.Header),
		}
		response.Header.Set("Content-Type", "application/json")
		return response, nil
	}
}

// contextKey 定义自定义上下文键类型，避免使用内置 string 类型
type contextKey string

// setupJsInstance 初始化 Js 实例和 HTTP 客户端
func setupJsInstance(t *testing.T, ctx context2.Context, key, val interface{}) (*Js, func()) {
	cfg := &config.Config{
		AppID:     "test-app-id",
		AppSecret: "test-app-secret",
		Cache:     cache.NewMemory(),
	}
	cacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", credential.CacheKeyOfficialAccountPrefix, cfg.AppID)
	if err := cfg.Cache.Delete(cacheKey); err != nil {
		t.Fatalf("清除缓存失败: %v", err)
	}
	t.Log("清除 jsapi_ticket 的缓存:", cacheKey)

	ctxHandle := &context.Context{Config: cfg, AccessTokenHandle: &mockAccessTokenHandle{}}
	jsInstance := NewJs(ctxHandle, cfg.AppID)
	jsInstance.SetJsTicketHandle(credential.NewDefaultJsTicket(cfg.AppID, credential.CacheKeyOfficialAccountPrefix, cfg.Cache))

	originalClient := util.DefaultHTTPClient
	util.DefaultHTTPClient = &http.Client{
		Transport: &contextCheckingRoundTripper{originalCtx: ctx, t: t, key: key, expectedVal: val},
	}
	return jsInstance, func() { util.DefaultHTTPClient = originalClient }
}

// TestGetConfigContext 测试GetConfigContext的上下文传递和取消行为。
func TestGetConfigContext(t *testing.T) {
	t.Run("ContextPassing", func(t *testing.T) {
		ctxKey := contextKey("testKey111") // 使用自定义类型 contextKey
		ctxValue := "testValue222"
		ctx := context2.WithValue(context2.Background(), ctxKey, ctxValue)
		t.Logf("创建的测试上下文: %p, 添加的键值对: %v=%v\n", ctx, ctxKey, ctxValue)

		jsInstance, cleanup := setupJsInstance(t, ctx, ctxKey, ctxValue)
		defer cleanup()
		t.Log("调用 GetConfigContext")
		config2, err := jsInstance.GetConfigContext(ctx, "https://www.baidu.com", "test-app-id")
		if err != nil {
			t.Fatalf("GetConfigContext 失败: %v", err)
		}
		if config2.AppID != "test-app-id" {
			t.Errorf("预期 AppID 为 %s，实际为 %s", "test-app-id", config2.AppID)
		}
	})

	t.Run("ContextCancellation", func(t *testing.T) {
		ctx, cancel := context2.WithCancel(context2.Background())
		defer cancel()

		jsInstance, cleanup := setupJsInstance(t, ctx, nil, nil)
		defer cleanup()

		cancel()
		t.Log("调用 GetConfigContext（已取消上下文）")
		_, err := jsInstance.GetConfigContext(ctx, "https://www.baidu.com", "test-app-id")
		if err == nil {
			t.Error("预期上下文取消错误，但 GetConfigContext 未返回错误")
		} else if !errors.Is(err, context2.Canceled) {
			t.Errorf("预期错误为 context.Canceled，实际为: %v", err)
		}
	})
}
