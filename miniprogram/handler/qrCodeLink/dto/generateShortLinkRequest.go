package dto

import "encoding/json"

type GenerateShortLinkRequest struct {
	PageUrl     string `json:"page_url" comment:"通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符"`
	PageTitle   string `json:"page_title" comment:"页面标题，不能包含违法信息，超过20字符会用... 截断代替"`
	IsPermanent bool   `json:"is_permanent" comment:"默认值false。生成的 Short Link 类型，短期有效：false，永久有效：true"`
}

func (p *GenerateShortLinkRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
