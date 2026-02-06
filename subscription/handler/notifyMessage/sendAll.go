package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SendAll 根据标签群发消息
// 本接口用于根据标签群发消息
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_sendall.html
func (wx *NotifyMessage) SendAll(body *dto.SendAllRequest) (*vo.SendAllResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.SendAllResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/sendall",
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
