package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetFans 获取关注用户列表
// 本接口用来获取账号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_getfans.html
func (wx *UserManagerUserInfo) GetFans(nextOpenId *string) (*vo.GetFansResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	if nextOpenId != nil {
		params.Add("next_openid", *nextOpenId)
	}
	result, response, err := utils.HttpGet[vo.GetFansResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/user/get",
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
