package vo

import (
	"github.com/yijiacode188/wxSDK/service/hanlder/custommenu/dto"
	"github.com/yijiacode188/wxSDK/service/model/vo"
)

type TryMatchMenuResponse struct {
	vo.Base
	Button dto.Button `json:"button" `
}
