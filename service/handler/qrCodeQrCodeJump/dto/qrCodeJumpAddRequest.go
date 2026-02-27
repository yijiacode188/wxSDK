package dto

import "encoding/json"

type QrCodeJumpAddRequest struct {
	Prefix        string   `json:"prefix" comment:"二维码规则。"`
	AppId         string   `json:"appid" comment:"扫了服务号二维码之后要跳转的小程序的appid。增加或修改“扫服务号二维码打开小程序”的二维码规则才需要传这个参数。"`
	Path          string   `json:"path" comment:"小程序功能页面。"`
	IsEdit        int      `json:"is_edit" comment:"编辑标志位，0 表示新增二维码规则，1 表示修改已有二维码规则（注意，已经发布的规则，不支持修改）"`
	OpenVersion   int      `json:"open_version" comment:"测试范围。1表示开发版（配置只对开发者生效）；2表示体验版（配置对管理员、体验者生效)；3表示正式版"`
	DebugUrl      []string `json:"debug_url" comment:"测试链接，至多 5 个用于测试的二维码完整链接"`
	PermitSubRule int      `json:"permit_sub_rule" comment:"是否独占符合二维码前缀匹配规则的所有子规 1 为不占用，2 为占用"`
}

func (p *QrCodeJumpAddRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
