package dialog

import (
	"os"
	"testing"
	"time"

	"github.com/silenceper/wechat/v2/aispeech/config"
	aispeechContext "github.com/silenceper/wechat/v2/aispeech/context"
	"github.com/silenceper/wechat/v2/cache"
)

func TestIntegrationAccessTokenAndQuery(t *testing.T) {
	if os.Getenv("AISPEECH_INTEGRATION") != "1" {
		t.Skip("set AISPEECH_INTEGRATION=1 to run real aispeech API test")
	}

	cfg := &config.Config{
		AppID:  os.Getenv("AISPEECH_APPID"),
		Token:  os.Getenv("AISPEECH_TOKEN"),
		AESKey: os.Getenv("AISPEECH_AES_KEY"),
		Cache:  cache.NewMemory(),
	}
	if cfg.AppID == "" || cfg.Token == "" || cfg.AESKey == "" {
		t.Fatal("AISPEECH_APPID, AISPEECH_TOKEN and AISPEECH_AES_KEY are required")
	}

	d := NewDialog(&aispeechContext.Context{
		Config:                   cfg,
		AccessTokenContextHandle: NewAccessToken(cfg),
	})

	accessToken, err := d.GetAccessToken()
	if err != nil {
		t.Fatalf("GetAccessToken error: %v", err)
	}
	if accessToken == "" {
		t.Fatal("access token is empty")
	}

	res, err := d.Query(&QueryRequest{
		Query:  "你好",
		Env:    "online",
		UserID: "gowechat-integration-test",
	})
	if err != nil {
		t.Fatalf("Query error: %v", err)
	}
	if res.RequestID == "" {
		t.Fatal("query request_id is empty")
	}
}

func TestIntegrationFullDialogFlow(t *testing.T) {
	if os.Getenv("AISPEECH_INTEGRATION_MUTATION") != "1" {
		t.Skip("set AISPEECH_INTEGRATION_MUTATION=1 to run import/publish/query real aispeech API test")
	}

	d := newIntegrationDialog(t)
	now := time.Now().Unix()
	question := "codex-smoke-test-question-" + time.Unix(now, 0).Format("20060102150405")
	answer := "codex-smoke-test-answer-" + time.Unix(now, 0).Format("20060102150405")

	importRes, err := d.ImportJSON(&ImportJSONRequest{
		Mode: 0,
		Data: []BotIntent{{
			Skill:     "CodexSmokeTest",
			Intent:    question,
			Disable:   false,
			Questions: []string{question},
			Answers:   []string{answer},
		}},
	})
	if err != nil {
		t.Fatalf("ImportJSON error: %v", err)
	}
	if importRes.TaskID == "" {
		t.Fatal("ImportJSON task_id is empty")
	}
	t.Logf("ImportJSON request_id=%s task_id=%s", importRes.RequestID, importRes.TaskID)

	importTask := waitAsyncTask(t, d, importRes.TaskID, 2*time.Minute)
	if importTask.State != 2 {
		t.Fatalf("import task failed: state=%d msg=%s progress=%d", importTask.State, importTask.Msg, importTask.Progress)
	}
	t.Logf("FetchAsync import request_id=%s state=%d progress=%d", importTask.RequestID, importTask.State, importTask.Progress)

	publishRes, err := d.Publish()
	if err != nil {
		t.Fatalf("Publish error: %v", err)
	}
	if publishRes.TaskID == "" {
		t.Fatal("Publish task_id is empty")
	}
	t.Logf("Publish request_id=%s task_id=%s", publishRes.RequestID, publishRes.TaskID)

	progress := waitEffectiveProgress(t, d, "online", 3*time.Minute)
	if progress.Status != 1 {
		t.Fatalf("publish progress failed: status=%d progress=%d", progress.Status, progress.Progress)
	}
	t.Logf("GetEffectiveProgress request_id=%s status=%d progress=%d", progress.RequestID, progress.Status, progress.Progress)

	queryRes := waitQueryAnswer(t, d, question, answer, 2*time.Minute)
	t.Logf("Query request_id=%s status=%s answer_type=%s answer=%s", queryRes.RequestID, queryRes.Status, queryRes.AnswerType, queryRes.Answer)
}

func newIntegrationDialog(t *testing.T) *Dialog {
	t.Helper()

	cfg := &config.Config{
		AppID:   os.Getenv("AISPEECH_APPID"),
		Token:   os.Getenv("AISPEECH_TOKEN"),
		AESKey:  os.Getenv("AISPEECH_AES_KEY"),
		Account: os.Getenv("AISPEECH_ACCOUNT"),
		Cache:   cache.NewMemory(),
	}
	if cfg.AppID == "" || cfg.Token == "" || cfg.AESKey == "" {
		t.Fatal("AISPEECH_APPID, AISPEECH_TOKEN and AISPEECH_AES_KEY are required")
	}

	return NewDialog(&aispeechContext.Context{
		Config:                   cfg,
		AccessTokenContextHandle: NewAccessToken(cfg),
	})
}

func waitAsyncTask(t *testing.T, d *Dialog, taskID string, timeout time.Duration) *FetchAsyncResponse {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for {
		res, err := d.FetchAsync(&FetchAsyncRequest{TaskID: taskID})
		if err != nil {
			t.Fatalf("FetchAsync error: %v", err)
		}
		if res.State == 2 || res.State == 3 {
			return res
		}
		if time.Now().After(deadline) {
			t.Fatalf("FetchAsync timeout: state=%d progress=%d msg=%s", res.State, res.Progress, res.Msg)
		}
		time.Sleep(3 * time.Second)
	}
}

func waitEffectiveProgress(t *testing.T, d *Dialog, env string, timeout time.Duration) *EffectiveProgressResponse {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for {
		res, err := d.GetEffectiveProgress(&EffectiveProgressRequest{Env: env})
		if err != nil {
			t.Fatalf("GetEffectiveProgress error: %v", err)
		}
		if res.Status == 1 || res.Status == 2 {
			return res
		}
		if time.Now().After(deadline) {
			t.Fatalf("GetEffectiveProgress timeout: status=%d progress=%d", res.Status, res.Progress)
		}
		time.Sleep(5 * time.Second)
	}
}

func waitQueryAnswer(t *testing.T, d *Dialog, question, answer string, timeout time.Duration) *QueryResponse {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for {
		res, err := d.Query(&QueryRequest{
			Query:  question,
			Env:    "online",
			UserID: "gowechat-full-integration-test",
		})
		if err != nil {
			t.Fatalf("Query error: %v", err)
		}
		if res.Answer == answer {
			return res
		}
		if time.Now().After(deadline) {
			t.Fatalf("Query timeout: status=%s answer=%s want=%s", res.Status, res.Answer, answer)
		}
		time.Sleep(5 * time.Second)
	}
}
