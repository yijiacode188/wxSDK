# wxSDK

Integrate WeChat official account platform, commonly used interfaces for mini programs, official accounts, and service accounts

## Getting Started

```cmd
go get github.com/yijiacode188/wxSDK
```

## For Example
```go
    package main

    import "github.com/yijiacode188/wxSDK/miniprogram"

    func main()  {
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

```