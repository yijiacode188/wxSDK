package base

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/base/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
	"time"
)

const ACCESS_TOKEN = "subscription_access_token"

// getAccessToken 获取接口调用凭据
// 本接口用于获取获取全局唯一后台接口调用凭据（Access Token），token 有效期为 7200 秒，开发者需要进行妥善保存，使用注意事项请参考此文档。
// https://developers.weixin.qq.com/doc/subscription/api/base/api_getaccesstoken.html
func (base *Base) getAccessToken() (string, *utils.Response, error) {
	token, ok := base.Store.GetStore(ACCESS_TOKEN)
	if ok {
		return token.(string), nil, nil
	}
	params := make(url.Values)
	params.Add("grant_type", "client_credential")
	params.Add("appid", base.AppId)
	params.Add("secret", base.Secret)
	result, response, err := utils.HttpGet[vo.GetAccessTokenResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/token",
		Params: params,
	})
	if err != nil {
		return "", response, err
	}
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	err = base.Store.SetStore(ACCESS_TOKEN, result.AccessToken, time.Now().Add(time.Duration(result.ExpiresIn-10)*time.Second))
	if err != nil {
		return "", response, err
	}
	return result.AccessToken, response, nil
}
