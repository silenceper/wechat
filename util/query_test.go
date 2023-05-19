package util

import "testing"

func TeTestQuery(t *testing.T) {
	result := Query(map[string]interface{}{
		"age":12,
		"name":"Alan",
		"cat":"Peter",
	})
	if result != "age=12&name=Alan&cat=Peter"{
		t.Error("NOT PASS")
	}
}