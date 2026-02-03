package dto

import (
	"encoding/json"
	"wxSDK/miniprogram/consts"
)

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
type GetQRCodeRequest struct {
	Path       string            `json:"path" comment:"扫码进入的小程序页面路径，最大长度 1024 个字符"`
	Width      int               `json:"width" comment:"二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px"`
	AutoColor  bool              `json:"auto_color" comment:"默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调"`
	LineColor  *LineColor        `json:"line_color" comment:"默认值{\"r\":0,\"g\":0,\"b\":0}	；auto_color 为 false 时生效，使用 rgb 设置颜色  十进制表示"`
	IsHyaLine  bool              `json:"is_hyaline" comment:"默认值false；是否需要透明底色，为 true 时，生成透明底色的小程序码"`
	EnvVersion consts.EnvVersion `json:"env_version" comment:"要打开的小程序版本。正式版为 release,体验版为 trial 开发版为 develop,默认为正式版"`
}

func (p *GetQRCodeRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
