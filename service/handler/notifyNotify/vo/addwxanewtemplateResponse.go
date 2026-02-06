package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type AddWXaNewTemplateResponse struct {
	vo.Base
	PriTmplId string `json:"priTmplId" comment:"添加至帐号下的模板id，发送小程序订阅消息时所需"`
}
