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
	session, err := client.GetCode2Session("code")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(fmt.Sprintf("openId：%s", session.OpenId))
}
