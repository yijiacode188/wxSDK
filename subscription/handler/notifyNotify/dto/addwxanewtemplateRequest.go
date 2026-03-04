package dto

import "encoding/json"

type AddWXaNewTemplateRequest struct {
	Tid       string `json:"tid" comment:"模板标题 id，可通过接口获取，也可登录小程序后台查看获取"`
	KidList   []int  `json:"kidList" comment:"开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合"`
	SceneDesc string `json:"sceneDesc" comment:"服务场景描述，15个字以内"`
}

func (p *AddWXaNewTemplateRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
