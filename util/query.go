package util

import "strings"

// Query 将Map序列化为Query参数
func Query(params map[string]string) string {
	finalString := make([]string, 0)
	for key, value := range params {
		finalString = append(finalString, strings.Join([]string{key, value}, "="))
	}
	return strings.Join(finalString, "&")
}
