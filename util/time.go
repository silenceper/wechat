package util

import "time"

//GetCurrTs return current timestamps
func GetCurrTS() int64 {
	return time.Now().Unix()
}
