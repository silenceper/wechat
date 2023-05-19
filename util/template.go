package util

import (
	"strings"
)

// 对字符串中的和map的key相同的字符串进行模板替换 仅支持 形如: {name} 
func Template(source string, data map[string]string) string {
	sourceCopy := &source
	for k, v := range data {
		*sourceCopy = strings.Replace(*sourceCopy, strings.Join([]string{"{", k, "}"}, ""), v, 1)
	}
	return *sourceCopy
}
