package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type DraftSwitchResponse struct {
	vo.Base
	IsOpen int `json:"is_open" comment:"仅 errcode==0 (即调用成功) 时返回，0 表示开关处于关闭，1 表示开启成功（或开关已开启"`
}
