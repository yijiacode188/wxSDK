package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateTag 创建标签
// 本接口用于创建公众号标签。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_createtag.html
func (wx *UserManagerTag) CreateTag(body *dto.CreateTagRequest) (*vo.CreateTagResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CreateTagResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/create",
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
