package context

import (
	"encoding/json"
	"testing"
)

var testdata = `{"authorizer_info":{"nick_name":"就爱浪","head_img":"http:\/\/wx.qlogo.cn\/mmopen\/xPKCxELaaj6hiaTZGv19oQPBJibb7hBoKmNOjQibCNOUycE8iaBhiaHOA6eC8hadQSAUZTuHUJl4qCIbCQGjSWialicfzWh4mdxuejY\/0","service_type_info":{"id":1},"verify_type_info":{"id":-1},"user_name":"gh_dcdbaa6f1687","alias":"ckeyer","qrcode_url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/FribWCoIzQbAX7R1PQ8iaxGonqKp0doYD2ibhC0uhx11LrRcblASiazsbQJTJ4icQnMzfH7G0SUPuKbibTA8Cs4uk5WQ\/0","business_info":{"open_pay":0,"open_shake":0,"open_scan":0,"open_card":0,"open_store":0},"idc":1,"principal_name":"个人","signature":"不折腾会死。"},"authorization_info":{"authorizer_appid":"yyyyy","authorizer_refresh_token":"xxxx","func_info":[{"funcscope_category":{"id":1}},{"funcscope_category":{"id":15}},{"funcscope_category":{"id":4}},{"funcscope_category":{"id":7}},{"funcscope_category":{"id":2}},{"funcscope_category":{"id":3}},{"funcscope_category":{"id":11}},{"funcscope_category":{"id":6}},{"funcscope_category":{"id":5}},{"funcscope_category":{"id":8}},{"funcscope_category":{"id":13}},{"funcscope_category":{"id":9}},{"funcscope_category":{"id":12}},{"funcscope_category":{"id":22}},{"funcscope_category":{"id":23}},{"funcscope_category":{"id":24},"confirm_info":{"need_confirm":0,"already_confirm":0,"can_confirm":0}},{"funcscope_category":{"id":26}},{"funcscope_category":{"id":27},"confirm_info":{"need_confirm":0,"already_confirm":0,"can_confirm":0}},{"funcscope_category":{"id":33},"confirm_info":{"need_confirm":0,"already_confirm":0,"can_confirm":0}},{"funcscope_category":{"id":35}}]}}`

// TestDecode
func TestDecode(t *testing.T) {
	var ret struct {
		AuthorizerInfo    *AuthorizerInfo `json:"authorizer_info"`
		AuthorizationInfo *AuthBaseInfo   `json:"authorization_info"`
	}
	json.Unmarshal([]byte(testdata), &ret)
	t.Logf("%+v", ret.AuthorizerInfo)
	t.Logf("%+v", ret.AuthorizationInfo)
}
