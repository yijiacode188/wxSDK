package miniprogram

import (
	"errors"
	"net/url"
	"time"
	"wxSDK/miniprogram/model/vo"
	"wxSDK/utils"
)

const ACCESS_TOKEN = "miniprogram_access_token"

// getAccessToken 获取接口调用凭据
// 本接口用于获取获取全局唯一后台接口调用凭据（Access Token），token 有效期为 7200 秒，开发者需要进行妥善保存，使用注意事项请参考此文档。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/mp-access-token/api_getaccesstoken.html
func (wx *wxClient) getAccessToken() (string, error) {

	token, ok := wx.Store.GetStore(ACCESS_TOKEN)
	if ok {
		return token.(string), nil
	}
	params := make(url.Values)
	params.Add("grant_type", "client_credential")
	params.Add("appid", wx.AppId)
	params.Add("secret", wx.Secret)
	result, _, err := utils.HttpGet[vo.GetAccessTokenResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/token",
		Params: params,
	})
	if err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", errors.New(result.ErrMsg)
	}
	err = wx.Store.SetStore(ACCESS_TOKEN, result.AccessToken, time.Now().Add(time.Duration(result.ExpiresIn)*time.Second))
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}
