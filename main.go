package main

import (
	"encoding/json"
	"fmt"
	"github.com/yijiacode188/wxSDK/subscription"
)

func main() {
	client, err := subscription.NewClient("appid", "secret")
	if err != nil {
		panic(err)
		return
	}

	result, _, err := client.GetMenu()
	if err != nil {
		panic(err)
		return
	}
	marshal, err := json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Println("创建成功", string(marshal))
}
