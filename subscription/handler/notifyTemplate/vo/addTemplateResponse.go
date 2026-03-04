package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type AddTemplateResponse struct {
	vo.Base
	TemplateId string `json:"template_id" comment:"模板id"`
}
