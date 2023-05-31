package util

import (
	"testing"
)

// TestTemplate testing case about Template method
func TestTemplate(t *testing.T) {
	result := Template("{name}={age};{with}={another};any={any};boolean={boolean}", map[string]interface{}{
		"name":    "Helan",
		"age":     "33",
		"with":    "Pep",
		"another": "C",
		"any":     33,
		"boolean": false,
	})
	if result != "Helan=33;Pep=C;any=33;boolean=false" {
		t.Error("NOT PSS testing")
	}
}
