package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// BatchUnTagging 批量为用户取消标签
// 本接口用于批量为多个用户取消标签
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_batchuntagging.html
func (wx *UserManagerTag) BatchUnTagging(body *dto.BatchUnTaggingRequest) (*vo.BatchUnTaggingResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.BatchUnTaggingResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging",
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
