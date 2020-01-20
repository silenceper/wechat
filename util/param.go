package util

import (
	"bytes"
	"sort"
	"strconv"
)

// OrderParam order params
func OrderParam(source interface{}, bizKey string) (returnStr string) {
	switch v := source.(type) {
	case map[string]string:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v[k])
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			switch vv := v[k].(type) {
			case string:
				buf.WriteString(vv)
			case int:
				buf.WriteString(strconv.FormatInt(int64(vv), 10))
			default:
				panic("params type not supported")
			}
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	}
	return
}
