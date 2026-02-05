package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/model/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GenerateShortLink 获取ShortLink
// 获取小程序 Short Link，适用于微信内拉起小程序的业务场景。目前对所有非个人主体小程序开放。通过该接口，可以选择生成到期失效和永久有效的小程序短链
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/short-link/api_generateshortlink.html
func (wx *wxClient) GenerateShortLink(body *dto.GenerateShortLinkRequest) (string, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return "", nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetGenerateShortLinkResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/genwxashortlink",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	return result.Link, response, nil
}
