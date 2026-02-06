package notifySubscribe

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifySubscribe/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifySubscribe/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// TemplateSubscribe 发送一次性订阅消息
// 推送订阅模板消息给授权微信用户
// https://developers.weixin.qq.com/doc/service/api/notify/subscribe/api_templatesubscribe.html#Body__data
func (wx *NotifySubscribe) TemplateSubscribe(body *dto.TemplateSubscribeRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.TemplateSubscribeResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/template/subscribe",
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
