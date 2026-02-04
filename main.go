package main

import (
	"encoding/json"
	"fmt"
	"github.com/yijiacode188/wxSDK/subscription"
)

func main() {
	client, err := subscription.NewClient("wx4ec69cbaee743a66", "f5655dae68d660bc58f05a46ce1f91dc")
	if err != nil {
		panic(err)
		return
	}

	result, err := client.GetMenu()
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
