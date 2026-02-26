package public

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/public/dto"
	"github.com/yijiacode188/wxSDK/service/handler/public/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FreePublishSubmit 发布草稿
// 本接口用于将图文草稿提交发布。
// https://developers.weixin.qq.com/doc/service/api/public/api_freepublish_submit.html
func (wx *Public) FreePublishSubmit(body *dto.FreePublishSubmitRequest) (*vo.FreePublishSubmitResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FreePublishSubmitResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/freepublish/submit",
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
