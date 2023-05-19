package util

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	result := Template("{name}={age};{with}={another}",map[string]string{
		"name":"Helan",
		"age":"33",
		"with":"Pep",
		"another":"C",
	})
	if result != "Helan=33;Pep=C"{
		t.Error("NOT PSS testing")
	}
}