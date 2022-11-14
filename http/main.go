package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	dns := "www.baidu.com"

	// 解析cname
	cname, _ := net.LookupCNAME(dns)
	fmt.Println("cname", cname)
	// 解析ip地址
	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Println("ns", ns)
	// 反向解析(主机必须得能解析到地址)
	dnsname, _ := net.LookupAddr("127.0.0.1")
	fmt.Println("hostname:", dnsname)

	res, _ := HttpGet("https://www.baidu.com")
	fmt.Println(res)
}

func HttpGet(url string) (string, error) {
	// 设置请求超时时间，0表示没有超时限制
	client := &http.Client{Timeout: 10 * time.Second}
	// 发送请求
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	// 必须关闭响应主体
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
