package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetBlackList 获取公众号的黑名单列表
// 本接口用来获取账号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_getblacklist.html
func (wx *UserManagerUserInfo) GetBlackList(body *dto.GetBlackListRequest) (*vo.GetBlackListResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetBlackListResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
