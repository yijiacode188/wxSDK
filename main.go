package main

import (
	"fmt"
	"github.com/yijiacode188/wxSDK/service"
)

func main() {
	client, err := service.NewClient("appId", "secret")
	if err != nil {
		panic(err)
		return
	}

	result, _, err := client.GetCurrentAutoReplyInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("发送成功", result)
}
