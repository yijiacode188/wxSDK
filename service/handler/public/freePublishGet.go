package public

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/public/dto"
	"github.com/yijiacode188/wxSDK/service/handler/public/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FreePublishGet 发布状态查询
// 本接口用于查询发布任务的状态和详情
// https://developers.weixin.qq.com/doc/service/api/public/api_freepublish_get.html
func (wx *Public) FreePublishGet(body *dto.FreePublishGetRequest) (*vo.FreePublishGetResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FreePublishGetResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/freepublish/get",
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
