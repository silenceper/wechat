package message

//TransferCustomer 转发客服消息
type TransferCustomer struct {
	CommonToken

	TransInfo *TransInfo `xml:"TransInfo,omitempty"`
}

//TransInfo 转发到指定客服
type TransInfo struct {
	KfAccount string `xml:"KfAccount"`
}

//NewTransferCustomer 实例化
func NewTransferCustomer(KfAccount string) *TransferCustomer {
	tc := new(TransferCustomer)
	if KfAccount != "" {
		transInfo := new(TransInfo)
		transInfo.KfAccount = KfAccount
		tc.TransInfo = transInfo
	}
	return tc
}
