package dto

import (
	"encoding/json"
)

type MatchRule struct {
	TagId              string `json:"tag_id" comment:"用户标签的id，可通过用户标签管理接口获取"`
	ClientPlatformType string `json:"client_platform_type" comment:"客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配"`
}
type AddConditionalMenuRequest struct {
	Button    []Button  `json:"button" comment:"一级菜单数组(1-3个)"`
	MatchRule MatchRule `json:"matchrule" comment:"菜单匹配规则(至少一个非空字段)"`
}

func (p *AddConditionalMenuRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
