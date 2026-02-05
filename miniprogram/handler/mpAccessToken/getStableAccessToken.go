package mpAccessToken

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"time"
)

// GetStableAccessToken 获取稳定版接口调用凭据
// 本接口用于获取获取全局唯一后台接口调用凭据（Access Token），token 有效期为 7200 秒，但此接口和 getAccessToken 互相隔离，且比其更加稳定，推荐使用此接口替代。开发者需要进行妥善保存，使用注意事项请参考此文档。
// 有两种调用模式:
// 普通模式，access_token 有效期内重复调用该接口不会更新 access_token，绝大部分场景下使用该模式；
// 强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token；
// 如使用云开发，可通过云调用免维护 access_token 调用；
// 如使用云托管，也可以通过微信令牌/开放接口服务免维护 access_token 调用；
// https://developers.weixin.qq.com/miniprogram/dev/server/API/mp-access-token/api_getstableaccesstoken.html
func (mp *MPAccessToken) GetStableAccessToken(forceRefresh bool) (string, *utils.Response, error) {
	token, ok := mp.Store.GetStore(ACCESS_TOKEN)
	if ok {
		return token.(string), nil, nil
	}
	body := &dto.GetStableAccessTokenRequest{
		AppID:        mp.AppId,
		Secret:       mp.Secret,
		ForceRefresh: forceRefresh,
	}
	result, response, err := utils.HttpPost[vo.GetStableAccessTokenResponse](&utils.RequestParams{
		Url:  "https://api.weixin.qq.com/cgi-bin/stable_token",
		Body: body.ToByte(),
	})
	if err != nil {
		return "", response, nil
	}
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	err = mp.Store.SetStore(ACCESS_TOKEN, result.AccessToken, time.Now().Add(time.Duration(result.ExpiresIn-10)*time.Second))
	if err != nil {
		return "", response, err
	}
	return result.AccessToken, response, nil
}
