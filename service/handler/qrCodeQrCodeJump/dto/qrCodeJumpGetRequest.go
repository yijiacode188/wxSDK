package dto

import "encoding/json"

type QrCodeJumpGetRequest struct {
	AppID      string   `json:"appid" comment:"小程序的appid。获取“扫服务号二维码打开小程序”已设置的二维码规则才需要传这个参数。"`
	GetType    int      `json:"get_type" comment:"默认值为0。 0：查询最近新增 10000 条（数量大建议用1或者2）；1：prefix查询；2：分页查询，按新增顺序返回。获取“扫服务号二维码打开小程序”已设置的二维码规则才需要传这个参数。"`
	PrefixList []string `json:"prefix_list" comment:"prefix查询，get_type=1 必传，最多传 200 个前缀。获取“扫服务号二维码打开小程序”已设置的二维码规则才需要传这个参数。"`
	PageNum    int64    `json:"page_num" comment:"页码，get_type=2 必传，从 1 开始。获取“扫服务号二维码打开小程序”已设置的二维码规则才需要传这个参数。"`
	PageSize   int64    `json:"page_size" comment:"每页数量，get_type=2 必传，最大为 200。获取“扫服务号二维码打开小程序”已设置的二维码规则才需要传这个参数。"`
}

func (p *QrCodeJumpGetRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
