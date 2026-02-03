package dto

import "encoding/json"

type GetCreateQRCodeRequest struct {
	Path  string `json:"path" comment:"扫码进入的小程序页面路径，最大长度 128 个字符，不能为空"`
	Width int    `json:"width" comment:"二维码的宽度，单位 px。最小 280px，最大 1280px;默认是430"`
}

func (p *GetCreateQRCodeRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
