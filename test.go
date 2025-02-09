package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("成功访问 Wikipedia，返回数据大小:", len(body))
}
