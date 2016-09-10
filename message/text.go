package message

//Text 文本消息
type Text struct {
	CommonToken
	Content string `xml:"Content"`
}
