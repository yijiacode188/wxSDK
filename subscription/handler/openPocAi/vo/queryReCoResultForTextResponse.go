package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type QueryReCoResultForTextResponse struct {
	vo.Base
	Result string `json:"result" comment:"识别结果"`
}
