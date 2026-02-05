package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
)

type GeRidInfoResponse struct {
	vo.Base
	Request struct {
		InvokeTime   int    `json:"invoke_time" comment:"发起请求的时间戳"`
		CostInMs     int    `json:"cost_in_ms" comment:"请求毫秒级耗时"`
		RequestUrl   string `json:"request_url" comment:"请求的URL参数"`
		RequestBody  string `json:"request_body" comment:"post请求的请求参数"`
		ResponseBody string `json:"response_body" comment:"接口请求返回参数"`
		ClientIp     string `json:"client_ip" comment:"接口请求的客户端ip"`
	} `json:"request" comment:"该rid对应的请求详情"`
}
