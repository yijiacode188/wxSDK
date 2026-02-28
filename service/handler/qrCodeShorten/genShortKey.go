package qrCodeShorten

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeShorten/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeShorten/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GenShortKey 长信息转短链
// 短key托管类似于短链API，开发者可以通过GenShorten将不超过4KB的长信息转成短key，再通过FetchShorten将短key还原为长信息。
// https://developers.weixin.qq.com/doc/service/api/qrcode/shorten/api_genshortkey.html
func (wx *QrCodeShorten) GenShortKey(body *dto.GenShortKeyRequest) (*vo.GenShortKeyResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GenShortKeyResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/shorten/gen",
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
