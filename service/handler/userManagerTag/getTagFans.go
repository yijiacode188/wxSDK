package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetTagFans 获取标签下粉丝列表
// 本接口用于获取标签下粉丝列表。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_gettagfans.html
func (wx *UserManagerTag) GetTagFans(body *dto.GetTagFansRequest) (*vo.GetTagFansResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetTagFansResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/user/tag/get",
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
