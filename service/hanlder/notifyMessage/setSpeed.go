package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SetSpeed 设置群发速度
// 本接口用于设置消息的群发速度
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_setspeed.html
func (wx *NotifyMessage) SetSpeed(body *dto.SetSpeedRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.SetSpeedResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/speed/set",
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
