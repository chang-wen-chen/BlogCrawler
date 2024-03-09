package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	// 指定要爬取的网页地址
// 	url := "https://chang-wen-chen.github.io/2024/03/sw_gcloud_storage"

// 	// 发送 HTTP GET 请求
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("HTTP GET 请求失败:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// 读取响应内容
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("读取响应内容失败:", err)
// 		return
// 	}

// 	// 将响应内容打印到控制台
// 	fmt.Println(string(body))
// }
