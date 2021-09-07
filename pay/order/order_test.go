package order

import (
	"encoding/json"
	offConfig "github.com/silenceper/wechat/v2/pay/config"
	"testing"
)

func TestOrder_CloseOrder(t *testing.T) {
	cfg := &offConfig.Config{
		AppID:     "xxxx",
		MchID:     "xxxx",
		Key:       "xxxx",
		NotifyURL: "xxxx",
	}

	wOrder := NewOrder(cfg)
	result, err := wOrder.CloseOrder(&CloseParams{
		OutTradeNo: "xxxx",
		SignType:   "",
	})

	if err != nil {
		t.Error(err)
		return
	}
	res, _ := json.Marshal(result)
	t.Log(string(res))
}

func TestOrder_QueryOrder(t *testing.T) {
	cfg := &offConfig.Config{
		AppID:     "xxxx",
		MchID:     "xxxx",
		Key:       "xxxx",
		NotifyURL: "xxxx",
	}

	wOrder := NewOrder(cfg)
	result, err := wOrder.QueryOrder(&QueryParams{
		OutTradeNo: "xxxx",
	})

	if err != nil {
		t.Error(err)
		return
	}
	res, _ := json.Marshal(result)
	t.Log(string(res))
}
