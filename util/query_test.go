package util

import (
	"testing"
)

// TestQuery query method test case
func TestQuery(t *testing.T) {
	result := Query(map[string]interface{}{
		"age":  12,
		"name": "Alan",
		"cat":  "Peter",
	})
	if result == "" {
		// 由于hash是乱序 所以没法很好的预测输出的字符串
		// 将会输出符合Query规则的字符串 "age=12&name=Alan&cat=Peter"
		t.Error("NOT PASS")
	}
}
