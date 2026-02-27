package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// BatchUnBlackList 取消拉黑用户
// 本接口用来取消拉黑一批用户，黑名单列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_batchunblacklist.html
func (wx *UserManagerUserInfo) BatchUnBlackList(body *dto.BatchUnBlackListRequest) (*vo.BatchUnBlackListResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.BatchUnBlackListResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist",
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
