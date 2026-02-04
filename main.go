package main

import (
	"fmt"
	"github.com/yijiacode188/wxSDK/subscription"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
)

func main() {
	client, err := subscription.NewClient("wx4ec69cbaee743a66", "f5655dae68d660bc58f05a46ce1f91dc")
	if err != nil {
		panic(err)
		return
	}

	err = client.ClearApiQuota(&dto.ClearApiQuotaRequest{
		CgiPath: "/channels/ec/basics/info/get",
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("info")
}
