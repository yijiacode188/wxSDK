package miniprogram

import (
	"errors"
	"net/url"
	"wxSDK/miniprogram/model/dto"
	"wxSDK/miniprogram/model/vo"
	"wxSDK/utils"
)

// GetQueryUrlLink 查询加密URLLink
// 该接口用于查询小程序加密 url_link 配置
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-link/api_queryurllink.html
func (wx *wxClient) GetQueryUrlLink(body *dto.GetQueryUrlLinkRequest) (*vo.GetQueryUrlLinkResponse, error) {
	token, err := wx.getAccessToken()
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.GetQueryUrlLinkResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/query_urllink",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
