package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DelWxaNewTemplate 删除模板
// 删除私有模板库中的模板
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_delwxanewtemplate.html
func (wx *NotifyNotify) DelWxaNewTemplate(body *dto.DelWxaNewTemplateRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DelWxaNewTemplateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate",
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
