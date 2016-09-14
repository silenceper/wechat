package util

import "time"

//GetCurrTs return current timestamps
func GetCurrTs() int64 {
	return time.Now().Unix()
}
