package util

import (
	"io/ioutil"
	"net/http"
)

//HTTPGet get 请求
func HTTPGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

//HTTPPost post 请求
func HTTPPost() {
}
