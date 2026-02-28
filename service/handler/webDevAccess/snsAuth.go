package webDevAccess

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/webDevAccess/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SnsAuth 检验用户授权凭证
// 检验授权凭证（access_token）是否有效
// https://developers.weixin.qq.com/doc/service/api/webdev/access/api_snsauth.html
func (wx *WebDevAccess) SnsAuth(openId, accessToken string) (*vo.SnsAuthResponse, *utils.Response, error) {
	params := make(url.Values)
	params.Add("access_token", accessToken)
	params.Add("openid", openId)
	result, response, err := utils.HttpGet[vo.SnsAuthResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/sns/auth",
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
