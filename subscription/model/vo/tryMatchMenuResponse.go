package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo/base"
)

type TryMatchMenuResponse struct {
	base.Base
	Button dto.Button `json:"button" `
}
