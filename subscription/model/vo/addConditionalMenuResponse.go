package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo/base"

type AddConditionalMenuResponse struct {
	base.Base
	MenuId string `json:"menuid" comment:"菜单id"`
}
