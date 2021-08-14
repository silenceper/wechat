package syncmsg

// Event 微信客服回调事件
type Event struct {
	ToUserName string `json:"to_user_name"` // 微信客服组件ID
	CreateTime int    `json:"create_time"`  // 消息创建时间，unix时间戳
	MsgType    string `json:"msgtype"`      // 消息的类型，此时固定为 event
	Event      string `json:"event"`        // 事件的类型，此时固定为 kf_msg_or_event
	Token      string `json:"token"`        // 调用拉取消息接口时，需要传此token，用于校验请求的合法性
}
