package generic

import (
	context2 "context"
	"net/url"
	"strings"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

// defaultBaseURL 微信 API 默认基础地址
const defaultBaseURL = "https://api.weixin.qq.com"

// Generic 通用 API 调用，用于访问 SDK 尚未封装的小程序接口。
// 支持自定义 path、query 参数、request body 以及 response 结构体。
type Generic struct {
	*context.Context
}

// NewGeneric init
func NewGeneric(ctx *context.Context) *Generic {
	return &Generic{ctx}
}

// BuildURL 构建完整请求 URL。
// path 支持相对路径（如 "/intp/xxx"）或完整 URL（如 "https://api.weixin.qq.com/intp/xxx"）。
// query 为额外的 URL 查询参数，可为 nil；access_token 会自动追加（若 query 中已包含则不重复追加）。
func (g *Generic) BuildURL(path string, query url.Values) (string, error) {
	return g.BuildURLContext(context2.Background(), path, query)
}

// BuildURLContext 构建完整请求 URL（带 context）
func (g *Generic) BuildURLContext(ctx context2.Context, path string, query url.Values) (string, error) {
	var baseURL string
	switch {
	case strings.HasPrefix(path, "http://"), strings.HasPrefix(path, "https://"):
		baseURL = path
	case strings.HasPrefix(path, "/"):
		baseURL = defaultBaseURL + path
	default:
		baseURL = defaultBaseURL + "/" + path
	}

	// 克隆 query，避免修改调用方原始值
	q := make(url.Values, len(query)+1)
	for k, v := range query {
		q[k] = v
	}

	// 自动追加 access_token
	if q.Get("access_token") == "" {
		accessToken, err := g.GetAccessTokenContext(ctx)
		if err != nil {
			return "", err
		}
		q.Set("access_token", accessToken)
	}

	sep := "?"
	if strings.Contains(baseURL, "?") {
		sep = "&"
	}
	return baseURL + sep + q.Encode(), nil
}

// Get 发送 GET 请求，返回原始响应字节。
// 使用 util.DecodeWithError 可将响应解码到自定义结构体。
func (g *Generic) Get(path string, query url.Values) (response []byte, err error) {
	return g.GetContext(context2.Background(), path, query)
}

// GetContext 发送 GET 请求（带 context），返回原始响应字节
func (g *Generic) GetContext(ctx context2.Context, path string, query url.Values) (response []byte, err error) {
	var uri string
	if uri, err = g.BuildURLContext(ctx, path, query); err != nil {
		return
	}
	return util.HTTPGetContext(ctx, uri)
}

// PostJSON 发送 POST 请求（JSON body），返回原始响应字节。
// body 可为 nil（无请求体）；使用 util.DecodeWithError 可将响应解码到自定义结构体。
func (g *Generic) PostJSON(path string, query url.Values, body interface{}) (response []byte, err error) {
	return g.PostJSONContext(context2.Background(), path, query, body)
}

// PostJSONContext 发送 POST 请求（JSON body，带 context），返回原始响应字节
func (g *Generic) PostJSONContext(ctx context2.Context, path string, query url.Values, body interface{}) (response []byte, err error) {
	var uri string
	if uri, err = g.BuildURLContext(ctx, path, query); err != nil {
		return
	}
	return util.PostJSONContext(ctx, uri, body)
}

// GetDecode 发送 GET 请求并将响应解码到 resp。
// resp 应为指向结构体的指针，结构体中可嵌入 util.CommonError 以自动校验错误码。
// apiName 用于错误信息标识。
func (g *Generic) GetDecode(path string, query url.Values, resp interface{}, apiName string) (err error) {
	return g.GetDecodeContext(context2.Background(), path, query, resp, apiName)
}

// GetDecodeContext 发送 GET 请求并将响应解码到 resp（带 context）
func (g *Generic) GetDecodeContext(ctx context2.Context, path string, query url.Values, resp interface{}, apiName string) (err error) {
	var response []byte
	if response, err = g.GetContext(ctx, path, query); err != nil {
		return
	}
	return util.DecodeWithError(response, resp, apiName)
}

// PostJSONDecode 发送 POST 请求（JSON body）并将响应解码到 resp。
// resp 应为指向结构体的指针，结构体中可嵌入 util.CommonError 以自动校验错误码。
// apiName 用于错误信息标识。
func (g *Generic) PostJSONDecode(path string, query url.Values, body, resp interface{}, apiName string) (err error) {
	return g.PostJSONDecodeContext(context2.Background(), path, query, body, resp, apiName)
}

// PostJSONDecodeContext 发送 POST 请求（JSON body）并将响应解码到 resp（带 context）
func (g *Generic) PostJSONDecodeContext(ctx context2.Context, path string, query url.Values, body, resp interface{}, apiName string) (err error) {
	var response []byte
	if response, err = g.PostJSONContext(ctx, path, query, body); err != nil {
		return
	}
	return util.DecodeWithError(response, resp, apiName)
}
