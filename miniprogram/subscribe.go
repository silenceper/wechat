package miniprogram

const (
	subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
)

// SubscribeSend 发送订阅消息
func (wxa *MiniProgram) SubscribeSend(touser, templateId, page string, data map[string]interface{}) (result ResCode2Session, err error) {
	// TODO： 发送订阅消息实现
	_ = subscribeSendURL
	/*
		{
		    "touser": "orKb-41QOJ0zHo6JSwaOMPjMVCpM",
		    "template_id": "wddSQ2HuFBrbjLbrHT-fwJxjptLgQgsgOFrn_0qINt4",
		    "page": "pages/my/myHome/myHome",
		    "data": {
		        "character_string1": {
		            "value": "339208499"
		        },
		        "thing24": {
		            "value": "倒霉狐狸的消息内容哦"
		        }
		    }
		}
	*/
	return
}
