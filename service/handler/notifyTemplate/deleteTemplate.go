package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteTemplate 删除模板
// 本接口用于删除账号下的指定模板
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_deletetemplate.html
func (wx *NotifyTemplate) DeleteTemplate(body *dto.DeleteTemplateRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DeleteTemplateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/template/del_private_template",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return response, err
	}
	if result.ErrCode != 0 {
		return response, errors.New(result.ErrMsg)
	}
	return response, nil
}
