package userManagerUserInfo

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerUserInfo/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// UpdateRemark 设置用户备注名
// 本接口用于对指定用户设置备注名，该接口暂时开放给微信认证的服务号。 接口调用请求说明。
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_updateremark.html
func (wx *UserManagerUserInfo) UpdateRemark(body *dto.UpdateRemarkRequest) (*vo.UpdateRemarkResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)

	result, response, err := utils.HttpGet[vo.UpdateRemarkResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/user/info/updateremark",
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
