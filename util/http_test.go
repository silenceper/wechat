package util

import (
	"fmt"
	"testing"
)

func TestHTTPGet(t *testing.T) {
	//abc sig
	SetProxy("socks5://127.0.0.1:1080")
	url := "https://api.ip.sb/ip"
	resp,_ := HTTPGet(url)
	//如果resp打印是代理的IP，则代理成功
	fmt.Println(string(resp))

	CloseProxy()
	resp,_ = HTTPGet(url)
	//如果resp打印是本机IP，则关闭代理成功
	fmt.Println(string(resp))

	OpenProxy()

	resp,_ = HTTPGet(url)
	//如果resp打印是代理的IP，则打开代理成功
	fmt.Println(string(resp))

}
