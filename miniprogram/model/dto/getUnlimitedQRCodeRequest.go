package dto

import (
	"encoding/json"
	"github.com/yijiacode188/wxSDK/miniprogram/consts"
)

type GetUnlimitedQRCodeRequest struct {
	Scene      string            `json:"scene" comment:"最大32个可见字符，只支持数字，大小写英文以及部分特殊字符"`
	Page       string            `json:"page" comment:"默认是主页，页面 page，例如 pages/index/index"`
	CheckPath  bool              `json:"check_path" comment:"默认是true，检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）"`
	EnvVersion consts.EnvVersion `json:"env_version" comment:"要打开的小程序版本。正式版为 release,体验版为 trial 开发版为 develop,默认为正式版"`
	Width      int               `json:"width" comment:"默认430，二维码的宽度，单位 px，最小 280px，最大 1280px"`
	AutoColor  bool              `json:"auto_color" comment:"默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调"`
	LineColor  *LineColor        `json:"line_color" comment:"默认值{\"r\":0,\"g\":0,\"b\":0}	；auto_color 为 false 时生效，使用 rgb 设置颜色  十进制表示"`
	IsHyaLine  bool              `json:"is_hyaline" comment:"默认值false；是否需要透明底色，为 true 时，生成透明底色的小程序码"`
}

func (p *GetUnlimitedQRCodeRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
