package vo

import (
	"github.com/yijiacode188/wxSDK/service/model/vo"
)

type AddConditionalMenuResponse struct {
	vo.Base
	MenuId string `json:"menuid" comment:"菜单id"`
}
