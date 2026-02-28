package webDevAccess

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/webDevAccess/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SnsAccessToken 换取用户授权凭证
// 此接口与获取接口调用凭据不同，为用户级别的授权 token 获取。该 token 仅能获取授权用户的相关信息，无法用于调用其他接口。
// https://developers.weixin.qq.com/doc/service/api/webdev/access/api_snsaccesstoken.html
func (wx *WebDevAccess) SnsAccessToken(appId, secret, code string) (*vo.SnsAccessTokenResponse, *utils.Response, error) {
	params := make(url.Values)
	params.Add("appid", appId)
	params.Add("secret", secret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	result, response, err := utils.HttpGet[vo.SnsAccessTokenResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/sns/oauth2/access_token",
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
