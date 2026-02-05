package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/handler/custommenu/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
)

type TryMatchMenuResponse struct {
	vo.Base
	Button dto.Button `json:"button" `
}
