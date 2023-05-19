package util

import (
	"fmt"
	"strings"
)

// Query 将Map序列化为Query参数
func Query(params map[string]interface{}) string {
	finalString := make([]string, 0)
	for key, value := range params {
		finalString = append(finalString, strings.Join([]string{key, fmt.Sprintf("%s", value)}, "="))
	}
	return strings.Join(finalString, "&")
}
