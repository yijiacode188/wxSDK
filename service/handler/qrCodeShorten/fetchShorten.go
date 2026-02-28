package qrCodeShorten

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeShorten/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeShorten/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FetchShorten 短链转长信息
// 开发者可以通过GenShorten将不超过4KB的长信息转成短链key，再通过FetchShorten将短链key还原为长信息。
// https://developers.weixin.qq.com/doc/service/api/qrcode/shorten/api_fetchshorten.html
func (wx *QrCodeShorten) FetchShorten(body *dto.FetchShortenRequest) (*vo.FetchShortenResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FetchShortenResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/shorten/fetch",
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
