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

	result, _, err := client.GetSpeed()
	if err != nil {
		return
	}

	fmt.Println("创建成功", result)
}
