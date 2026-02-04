# wxSDK

Integrate WeChat official account platform, commonly used interfaces for mini programs, official accounts, and service accounts

## Getting Started

```cmd
go get github.com/yijiacode188/wxSDK
```

## For Example
公众号测试号 https://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index
```go
    package main

    import (
		"github.com/yijiacode188/wxSDK/miniprogram"
		"fmt"
    )
    

    func main()  {
		client, err := miniprogram.NewClient("appId", "secret")
		if err != nil {
			panic(err)
			return
		}
		session, err := client.Code2Session("code")
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(fmt.Sprintf("openId：%s", session.OpenId))
    }

```