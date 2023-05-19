package util

import (
	"strings"
)

func Template(source string, data map[string]string) string {
	sourceCopy := &source
	for k, v := range data {
		*sourceCopy = strings.Replace(*sourceCopy, strings.Join([]string{"{", k, "}"}, ""), v, 1)
	}
	return *sourceCopy
}
