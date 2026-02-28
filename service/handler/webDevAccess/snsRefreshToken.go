package webDevAccess

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/webDevAccess/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SnsRefreshToken 刷新用户授权凭证
// 由于access_token拥有较短的有效期，当access_token超时后，可以使用refresh_token进行刷新，refresh_token有效期为30天，当refresh_token失效之后，需要用户重新授权。
// https://developers.weixin.qq.com/doc/service/api/webdev/access/api_snsrefreshtoken.html
func (wx *WebDevAccess) SnsRefreshToken(appId, refreshToken string) (*vo.SnsRefreshTokenResponse, *utils.Response, error) {
	params := make(url.Values)
	params.Add("appid", appId)
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)
	result, response, err := utils.HttpGet[vo.SnsRefreshTokenResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/sns/oauth2/refresh_token",
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
