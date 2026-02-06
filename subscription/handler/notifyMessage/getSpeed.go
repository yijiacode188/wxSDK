package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetSpeed 获取群发速度
// 本接口用于获取消息的群发速度
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_getspeed.html
func (wx *NotifyMessage) GetSpeed() (*vo.GetSpeedResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetSpeedResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/speed/get",
		Params: params,
	})
	if err != nil {
		return nil, nil, err
	}
	if result.ErrCode != 0 {
		return nil, nil, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
