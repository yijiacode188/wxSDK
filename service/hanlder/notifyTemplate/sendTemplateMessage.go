package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyTemplate/dto"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SendTemplateMessage 发送模板消息
// 本接口用于向用户发送模板消息
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_sendtemplatemessage.html
func (wx *NotifyTemplate) SendTemplateMessage(body *dto.SendTemplateMessageRequest) (*vo.SendTemplateMessageResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.SendTemplateMessageResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/template/send",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
