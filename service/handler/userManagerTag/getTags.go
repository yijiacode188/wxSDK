package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetTags 获取标签
// 本接口用于获取公众号已创建的标签列表。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_gettags.html
func (wx *UserManagerTag) GetTags() (*vo.GetTagsResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetTagsResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/get",
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
