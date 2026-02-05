package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/model/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GenerateUrlLink 获取加密URLLink
// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。目前仅针对国内非个人主体的小程序开放，详见获取 URL Link
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/url-link/api_generateurllink.html
func (wx *wxClient) GenerateUrlLink(body *dto.GenerateUrlLinkRequest) (string, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return "", nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GenerateUrlLinkResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/generate_urllink",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	return result.UrlLink, response, nil
}
