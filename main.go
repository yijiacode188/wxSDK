package main

import (
	"fmt"
	"github.com/yijiacode188/wxSDK/miniprogram"
)

func main() {
	client, err := miniprogram.NewClient("appId", "secret")
	if err != nil {
		panic(err)
		return
	}

	result, _, err := client.GetPhoneNumber("123")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("创建成功", result)
}
