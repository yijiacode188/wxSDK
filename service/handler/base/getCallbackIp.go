package base

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/base/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetCallBackIp 获取微信推送服务器IP
// 该接口用于获取微信推送服务器 ip 地址（向开发者服务器推送信息的微信服务器来源地址）
// 如果开发者基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息。
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getcallbackip.html
func (base *Base) GetCallBackIp() (*vo.GetCallBackIpResponse, *utils.Response, error) {
	token, _, err := base.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetCallBackIpResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/getcallbackip",
		Params: params,
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
