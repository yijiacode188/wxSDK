package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// BatchUserInfo 批量获取用户基本信息
// 本接口用于批量获取用户基本信息。最多支持一次拉取100条。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_batchuserinfo.html
func (wx *UserManagerUserInfo) BatchUserInfo(body *dto.BatchUserInfoRequest) (*vo.BatchUserInfoResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.BatchUserInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/user/info/batchget",
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
