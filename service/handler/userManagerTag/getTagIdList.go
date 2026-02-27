package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetTagIdList 获取用户的标签列表
// 本接口用于获取粉丝用户的标签列表。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_gettagidlist.html
func (wx *UserManagerTag) GetTagIdList(body *dto.GetTagIdListRequest) (*vo.GetTagIdListResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetTagIdListResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/getidlist",
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
