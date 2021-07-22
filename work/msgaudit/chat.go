package msgaudit

import "encoding/json"

type ChatDataResponse struct {
	Error
	ChatDataList []ChatData `json:"chatdata,omitempty"`
}

func (c ChatDataResponse) IsError() bool {
	return c.ErrCode != 0
}

type ChatData struct {
	Seq              uint64 `json:"seq,omitempty"`                // 消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	MsgId            string `json:"msgid,omitempty"`              // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	PublickeyVer     uint32 `json:"publickey_ver,omitempty"`      // 加密此条消息使用的公钥版本号。
	EncryptRandomKey string `json:"encrypt_random_key,omitempty"` // 使用publickey_ver指定版本的公钥进行非对称加密后base64加密的内容，需要业务方先base64 decode处理后，再使用指定版本的私钥进行解密，得出内容。
	EncryptChatMsg   string `json:"encrypt_chat_msg,omitempty"`   // 消息密文。需要业务方使用将encrypt_random_key解密得到的内容，与encrypt_chat_msg，传入sdk接口DecryptData,得到消息明文。
}

type ChatMessage struct {
	Id         string   // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	From       string   // 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	ToList     []string // 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	Action     string   // 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	Type       string   // 消息类型
	originData []byte   // 原始消息对象
}

func (c ChatMessage) GetOriginMessage() (msg map[string]interface{}) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetTextMessage() (msg TextMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetImageMessage() (msg ImageMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetRevokeMessage() (msg RevokeMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetAgreeMessage() (msg AgreeMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetVoiceMessage() (msg VoiceMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetVideoMessage() (msg VideoMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetCardMessage() (msg CardMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetLocationMessage() (msg LocationMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetEmotionMessage() (msg EmotionMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetFileMessage() (msg FileMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetLinkMessage() (msg LinkMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetWeappMessage() (msg WeappMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetChatRecordMessage() (msg ChatRecordMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetTodoMessage() (msg TodoMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetVoteMessage() (msg VoteMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetCollectMessage() (msg CollectMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetRedpacketMessage() (msg RedpacketMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetMeetingMessage() (msg MeetingMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetDocMessage() (msg DocMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetMarkdownMessage() (msg MarkdownMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetNewsMessage() (msg NewsMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetCalendarMessage() (msg CalendarMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetMixedMessage() (msg MixedMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetMeetingVoiceCallMessage() (msg MeetingVoiceCallMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetVoipDocShareMessage() (msg VoipDocShareMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetExternalRedPacketMessage() (msg ExternalRedPacketMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetSphFeedMessage() (msg SphFeedMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}

func (c ChatMessage) GetSwitchMessage() (msg SwitchMessage) {
	_ = json.Unmarshal(c.originData, &msg)
	return msg
}
