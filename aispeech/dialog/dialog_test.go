package dialog

import (
	stdcontext "context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/silenceper/wechat/v2/aispeech/config"
	aispeechContext "github.com/silenceper/wechat/v2/aispeech/context"
	"github.com/silenceper/wechat/v2/aispeech/encryptor"
	"github.com/silenceper/wechat/v2/cache"
)

const (
	testAccessToken    = "access-token"
	testAccount        = "admin"
	testAESKey         = "q1Os1ZMe0nG28KUEx9lg3HjK7V5QyXvi212fzsgDqgz"
	testAppID          = "appid"
	testImportRequest  = "import-rid"
	testImportTaskID   = "task-import"
	testPublishTaskID  = "task-publish"
	testQuery          = "hello"
	testQueryAnswer    = "hello answer"
	testQueryRequestID = "query-rid"
	testToken          = "token"
	testTokenRequestID = "token-rid"
)

type emptyAccessTokenHandle struct{}

func (emptyAccessTokenHandle) GetAccessToken() (string, error) {
	return "", nil
}

func (emptyAccessTokenHandle) GetAccessTokenContext(ctx stdcontext.Context) (string, error) {
	return "", nil
}

func TestAccessTokenCache(t *testing.T) {
	var tokenRequests int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRequests++
		body := readBody(t, r)
		assertSign(t, r, testToken, body)
		if r.Header.Get("X-APPID") != testAppID {
			t.Fatalf("bad X-APPID: %s", r.Header.Get("X-APPID"))
		}
		_, _ = w.Write([]byte(accessTokenResponse("rid")))
	}))
	defer srv.Close()

	ak := NewAccessToken(&config.Config{
		AppID:   testAppID,
		Token:   testToken,
		Account: testAccount,
		BaseURL: srv.URL,
		Cache:   cache.NewMemory(),
	})

	token, err := ak.GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken error: %v", err)
	}
	if token != testAccessToken {
		t.Fatalf("bad token: %s", token)
	}
	token, err = ak.GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken second error: %v", err)
	}
	if token != testAccessToken {
		t.Fatalf("bad token second: %s", token)
	}
	if tokenRequests != 1 {
		t.Fatalf("token requests = %d, want 1", tokenRequests)
	}
}

func TestAccessTokenCacheByAccount(t *testing.T) {
	var tokenRequests int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRequests++
		body := readBody(t, r)
		assertSign(t, r, testToken, body)

		var req AccessTokenRequest
		if err := json.Unmarshal(body, &req); err != nil {
			t.Fatalf("bad token body: %v", err)
		}
		_, _ = w.Write([]byte(accountAccessTokenResponse(req.Account)))
	}))
	defer srv.Close()

	memory := cache.NewMemory()
	cfgA := testDialogConfig(srv.URL)
	cfgA.Account = "admin-a"
	cfgA.Cache = memory
	cfgB := testDialogConfig(srv.URL)
	cfgB.Account = "admin-b"
	cfgB.Cache = memory

	tokenA, err := NewAccessToken(cfgA).GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken A error: %v", err)
	}
	tokenB, err := NewAccessToken(cfgB).GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken B error: %v", err)
	}
	if tokenA != "admin-a-token" || tokenB != "admin-b-token" {
		t.Fatalf("bad tokens: %s %s", tokenA, tokenB)
	}
	if tokenRequests != 2 {
		t.Fatalf("token requests = %d, want 2", tokenRequests)
	}
}

func TestAccessTokenEmptyToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":0,"msg":"success","request_id":"rid","data":{"access_token":""}}`))
	}))
	defer srv.Close()

	ak := NewAccessToken(&config.Config{
		AppID:   testAppID,
		Token:   testToken,
		BaseURL: srv.URL,
		Cache:   cache.NewMemory(),
	})

	if _, err := ak.GetAccessToken(); err == nil {
		t.Fatal("GetAccessToken should return error")
	}
}

func TestDialogAPIs(t *testing.T) {
	srv := newDialogAPITestServer(t)
	defer srv.Close()

	d := newTestDialog(srv.URL)
	assertDialogImportJSON(t, d)
	assertDialogPublish(t, d)
	assertDialogProgress(t, d)
	assertDialogFetchAsync(t, d)
	assertDialogQuery(t, d)
}

func TestDialogEmptyAccessToken(t *testing.T) {
	d := NewDialog(&aispeechContext.Context{
		Config: &config.Config{
			AESKey: testAESKey,
		},
		AccessTokenContextHandle: emptyAccessTokenHandle{},
	})

	if _, err := d.ImportJSON(&ImportJSONRequest{}); err == nil {
		t.Fatal("ImportJSON should return error")
	}
	if _, err := d.Publish(); err == nil {
		t.Fatal("Publish should return error")
	}
	if _, err := d.Query(&QueryRequest{}); err == nil {
		t.Fatal("Query should return error")
	}
}

func TestDialogAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case tokenPath:
			_, _ = w.Write([]byte(accessTokenResponse(testTokenRequestID)))
		case importJSONPath:
			_, _ = w.Write([]byte(`{"code":210202,"msg":"forbidden","request_id":"import-rid","data":{"reason":"forbidden"}}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	d := newTestDialog(srv.URL)

	_, err := d.ImportJSON(&ImportJSONRequest{})
	if err == nil {
		t.Fatal("ImportJSON should return error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("ImportJSON error should be *APIError but %T", err)
	}
	if apiErr.Code != 210202 || apiErr.Msg != "forbidden" || apiErr.RequestID != testImportRequest {
		t.Fatalf("bad api error: %+v", apiErr)
	}
	if string(apiErr.Data) != `{"reason":"forbidden"}` {
		t.Fatalf("bad api error data: %s", apiErr.Data)
	}
}

func TestQueryPlainJSONError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case tokenPath:
			_, _ = w.Write([]byte(accessTokenResponse(testTokenRequestID)))
		case queryPath:
			_, _ = w.Write([]byte(`{"code":110002,"msg":"bad param","request_id":"query-rid"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	d := newTestDialog(srv.URL)

	_, err := d.Query(&QueryRequest{Query: testQuery})
	if err == nil {
		t.Fatal("Query should return error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("Query error should be *APIError but %T", err)
	}
	if apiErr.Code != 110002 || apiErr.RequestID != testQueryRequestID {
		t.Fatalf("bad api error: %+v", apiErr)
	}
}

func newDialogAPITestServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := readBody(t, r)
		assertSign(t, r, testToken, body)
		switch r.URL.Path {
		case tokenPath:
			handleDialogToken(t, w, r)
		case importJSONPath:
			handleDialogImport(t, w, r, body)
		case publishPath:
			handleDialogPublish(t, w, r, body)
		case effectiveProgressPath:
			assertToken(t, r)
			_, _ = w.Write([]byte(`{"code":0,"msg":"success","request_id":"progress-rid","data":{"end_time":"","progress":100,"status":1}}`))
		case fetchAsyncPath:
			assertToken(t, r)
			_, _ = w.Write([]byte(`{"code":0,"msg":"success","request_id":"fetch-rid","data":{"state":2,"msg":"","progress":100,"start":1,"end":2,"url":"","success_skill_info":[{"id":1,"name":"AAA","intents":[{"id":2,"name":"BBB"}]}]}}`))
		case queryPath:
			handleDialogQuery(t, w, r, body)
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
}

func handleDialogToken(t *testing.T, w io.Writer, r *http.Request) {
	t.Helper()
	if r.Header.Get("X-APPID") != testAppID {
		t.Fatalf("bad X-APPID: %s", r.Header.Get("X-APPID"))
	}
	_, _ = w.Write([]byte(accessTokenResponse(testTokenRequestID)))
}

func handleDialogImport(t *testing.T, w io.Writer, r *http.Request, body []byte) {
	t.Helper()
	assertToken(t, r)
	var req ImportJSONRequest
	if err := json.Unmarshal(body, &req); err != nil {
		t.Fatalf("bad import body: %v", err)
	}
	if len(req.Data) != 1 || req.Data[0].Skill != "pre-sale" {
		t.Fatalf("bad import request: %+v", req)
	}
	_, _ = w.Write([]byte(`{"code":0,"msg":"success","request_id":"import-rid","data":{"task_id":"task-import"}}`))
}

func handleDialogPublish(t *testing.T, w io.Writer, r *http.Request, body []byte) {
	t.Helper()
	assertToken(t, r)
	if len(body) != 0 {
		t.Fatalf("publish body should be empty: %s", body)
	}
	_, _ = w.Write([]byte(`{"code":0,"msg":"success","request_id":"publish-rid","data":{"task_id":"task-publish"}}`))
}

func handleDialogQuery(t *testing.T, w io.Writer, r *http.Request, body []byte) {
	t.Helper()
	assertToken(t, r)
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "text/plain") {
		t.Fatalf("bad content type: %s", r.Header.Get("Content-Type"))
	}
	plain, err := encryptor.Decrypt(testAESKey, string(body))
	if err != nil {
		t.Fatalf("decrypt query body error: %v", err)
	}
	var req QueryRequest
	if err = json.Unmarshal(plain, &req); err != nil {
		t.Fatalf("bad query body: %v", err)
	}
	if req.Query != testQuery {
		t.Fatalf("bad query: %+v", req)
	}
	_, _ = w.Write([]byte(encryptQueryResponse(t)))
}

func assertDialogImportJSON(t *testing.T, d *Dialog) {
	t.Helper()
	res, err := d.ImportJSON(&ImportJSONRequest{
		Mode: 0,
		Data: []BotIntent{{
			Skill:     "pre-sale",
			Intent:    "business-hours",
			Disable:   false,
			Questions: []string{"when are you open"},
			Answers:   []string{"9:00-18:00"},
		}},
	})
	if err != nil || res.TaskID != testImportTaskID || res.RequestID != testImportRequest {
		t.Fatalf("ImportJSON = %+v, %v", res, err)
	}
}

func assertDialogPublish(t *testing.T, d *Dialog) {
	t.Helper()
	res, err := d.Publish()
	if err != nil || res.TaskID != testPublishTaskID {
		t.Fatalf("Publish = %+v, %v", res, err)
	}
}

func assertDialogProgress(t *testing.T, d *Dialog) {
	t.Helper()
	res, err := d.GetEffectiveProgress(&EffectiveProgressRequest{Env: "online"})
	if err != nil || res.Progress != 100 {
		t.Fatalf("GetEffectiveProgress = %+v, %v", res, err)
	}
}

func assertDialogFetchAsync(t *testing.T, d *Dialog) {
	t.Helper()
	res, err := d.FetchAsync(&FetchAsyncRequest{TaskID: testImportTaskID})
	if err != nil || res.State != 2 || len(res.SuccessSkillInfo) != 1 {
		t.Fatalf("FetchAsync = %+v, %v", res, err)
	}
}

func assertDialogQuery(t *testing.T, d *Dialog) {
	t.Helper()
	res, err := d.Query(&QueryRequest{Query: testQuery, Env: "online"})
	if err != nil || res.Answer != testQueryAnswer || res.RequestID != testQueryRequestID {
		t.Fatalf("Query = %+v, %v", res, err)
	}
}

func newTestDialog(baseURL string) *Dialog {
	cfg := testDialogConfig(baseURL)
	return NewDialog(&aispeechContext.Context{
		Config:                   cfg,
		AccessTokenContextHandle: NewAccessToken(cfg),
	})
}

func testDialogConfig(baseURL string) *config.Config {
	return &config.Config{
		AppID:   testAppID,
		Token:   testToken,
		AESKey:  testAESKey,
		Account: testAccount,
		BaseURL: baseURL,
		Cache:   cache.NewMemory(),
	}
}

func accountAccessTokenResponse(account string) string {
	return `{"code":0,"msg":"success","request_id":"rid","data":{"access_token":"` + account + `-token"}}`
}

func accessTokenResponse(requestID string) string {
	return `{"code":0,"msg":"success","request_id":"` + requestID + `","data":{"access_token":"` + testAccessToken + `"}}`
}

func encryptQueryResponse(t *testing.T) string {
	t.Helper()
	cipherText, err := encryptor.Encrypt(testAESKey, []byte(queryResponse()))
	if err != nil {
		t.Fatalf("encrypt response error: %v", err)
	}
	return cipherText
}

func queryResponse() string {
	return `{"code":0,"msg":"success","request_id":"query-rid","data":{"answer":"hello answer","answer_type":"text","skill_name":"skill","intent_name":"intent","msg_id":"msg","status":"FAQ","slots":[{"name":"n","value":"v","norm":"v"}]}}`
}

func readBody(t *testing.T, r *http.Request) []byte {
	t.Helper()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("read body error: %v", err)
	}
	return body
}

func assertToken(t *testing.T, r *http.Request) {
	t.Helper()
	if r.Header.Get("X-OPENAI-TOKEN") != testAccessToken {
		t.Fatalf("bad X-OPENAI-TOKEN: %s", r.Header.Get("X-OPENAI-TOKEN"))
	}
}

func assertSign(t *testing.T, r *http.Request, token string, body []byte) {
	t.Helper()
	timestamp, err := strconv.ParseInt(r.Header.Get("timestamp"), 10, 64)
	if err != nil {
		t.Fatalf("bad timestamp: %v", err)
	}
	want := encryptor.Sign(token, timestamp, r.Header.Get("nonce"), body)
	if got := r.Header.Get("sign"); got != want {
		t.Fatalf("bad sign: got %s want %s body %s", got, want, body)
	}
}
