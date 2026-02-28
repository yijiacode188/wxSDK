package webDevAccess

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/consts"
	"github.com/yijiacode188/wxSDK/service/handler/webDevAccess/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SnsUserInfo 获取授权用户信息
// 如果网页授权作用域为snsapi_userinfo，则此时开发者可以通过access_token和openid拉取用户信息。
// https://developers.weixin.qq.com/doc/service/api/webdev/access/api_snsuserinfo.html
func (wx *WebDevAccess) SnsUserInfo(accessToken, openID string, lang *consts.Lang) (*vo.SnsUserInfoResponse, *utils.Response, error) {
	params := make(url.Values)
	params.Add("access_token", accessToken)
	params.Add("openid", openID)
	if lang != nil {
		params.Add("lang", string(*lang))
	}
	result, response, err := utils.HttpGet[vo.SnsUserInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/sns/userinfo",
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
