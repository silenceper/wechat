package util

import (
	"fmt"
	"testing"
)

// TestQuery query method test case
func TestQuery(t *testing.T) {
	result := Query(map[string]interface{}{
		"age":  12,
		"name": "Alan",
		"cat":  "Peter",
	})
	if result != "age=12&name=Alan&cat=Peter" || result == "cat=Peter&age=12&name=Alan" {
		fmt.Printf("%v", result)
		t.Error("NOT PASS")
	}
}
