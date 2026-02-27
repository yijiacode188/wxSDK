package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// UserInfo 获取用户基本信息
// 在关注者与公众号产生消息交互后，公众号可获得关注者的OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的。对于不同公众号，同一用户的openid不同）。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_userinfo.html
func (wx *UserManagerUserInfo) UserInfo(openId string, lang *string) (*vo.UserInfoResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("openid", openId)
	if lang != nil {
		params.Add("lang", *lang)
	}
	result, response, err := utils.HttpPost[vo.UserInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/user/info",
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
