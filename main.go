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

// 	url := "https://chang-wen-chen.github.io/2024/03/sw_gcloud_storage"

// 	response, err := http.Get(url)
// 	if err != nil {

// 		return
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {

// 		return
// 	}

// 	// 将响应内容打印到控制台
// 	fmt.Println(string(body))
// }
