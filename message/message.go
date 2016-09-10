package message

//EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}
