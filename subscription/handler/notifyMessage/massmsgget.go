package notifyMessage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/dto"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// MassMsgGet 查询群发消息发送状态
// 本接口用于查询群发消息发送状态
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_massmsgget.html
func (wx *NotifyMessage) MassMsgGet(body *dto.MassMsgGetRequest) (*vo.MassMsgGetResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.MassMsgGetResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/message/mass/get",
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
