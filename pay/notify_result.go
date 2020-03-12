package pay

import (
	"encoding/xml"
	"io"
	"sort"

	"github.com/silenceper/wechat/util"
)

// NotifyResult 下单回调
type NotifyResult map[string]string

// NotifyResp 消息通知返回
type NotifyResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// UnmarshalXML 解析XML为map
func (m *NotifyResult) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = NotifyResult{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

// VerifySign 验签
func (pcf *Pay) VerifySign(notifyRes NotifyResult) bool {
	// 支付key
	sortedKeys := make([]string, 0, len(notifyRes))
	for k := range notifyRes {
		if k == "sign" {
			continue
		}
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	// STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		value := notifyRes[k]
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	// STEP3, 在键值对的最后加上key=API_KEY
	signStrings = signStrings + "key=" + pcf.PayKey
	// STEP4, 进行MD5签名并且将所有字符转为大写.
	sign := util.MD5Sum(signStrings)
	if sign != notifyRes["sign"] {
		return false
	}
	return true
}
