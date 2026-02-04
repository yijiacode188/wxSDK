package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetCallBackIp 获取微信推送服务器IP
// 该接口用于获取微信推送服务器 ip 地址（向开发者服务器推送信息的微信服务器来源地址）
// 如果开发者基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息。
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getcallbackip.html
func (wx *wxClient) GetCallBackIp() (*vo.GetCallBackIpResponse, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.GetCallBackIpResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/getcallbackip",
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
