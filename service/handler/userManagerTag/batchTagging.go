package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// BatchTagging 批量为用户打标签
// 本接口用于批量为多个用户打标签。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_batchtagging.html
func (wx *UserManagerTag) BatchTagging(body *dto.BatchTaggingRequest) (*vo.BatchTaggingResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.BatchTaggingResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging",
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
