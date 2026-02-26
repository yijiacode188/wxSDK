package public

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/public/dto"
	"github.com/yijiacode188/wxSDK/service/handler/public/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FreePublishBatchGet 获取已发布的消息列表
// 本接口用于获取已成功发布的消息列表
// https://developers.weixin.qq.com/doc/service/api/public/api_freepublish_batchget.html
func (wx *Public) FreePublishBatchGet(body *dto.FreePublishBatchGetRequest) (*vo.FreePublishBatchGetResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FreePublishBatchGetResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/freepublish/batchget",
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
