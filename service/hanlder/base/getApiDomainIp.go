package base

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/base/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetApiDomainIp 获取微信API服务器IP
// 该接口用于获取微信 api 服务器 ip 地址（开发者服务器主动访问 api.weixin.qq.com 的远端地址）
// 如果开发者基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息。
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getapidomainip.html
func (base *Base) GetApiDomainIp() (*vo.GetApiDomainIpResponse, *utils.Response, error) {
	token, _, err := base.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetApiDomainIpResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/get_api_domain_ip",
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
