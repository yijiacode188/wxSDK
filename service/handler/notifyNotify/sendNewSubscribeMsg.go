package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SendNewSubscribeMsg 发送订阅通知
// 发送订阅通知
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_sendnewsubscribemsg.html
func (wx *NotifyNotify) SendNewSubscribeMsg(body *dto.SendNewSubscribeMsgRequest) (*vo.SendNewSubscribeMsgResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.SendNewSubscribeMsgResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/subscribe/bizsend",
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
