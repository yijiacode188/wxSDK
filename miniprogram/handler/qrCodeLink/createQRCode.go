package qrCodeLink

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/qrCodeLink/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/handler/qrCodeLink/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateQRCode 获取小程序二维码
// 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_createqrcode.html
func (mp *QRCodeLink) CreateQRCode(body *dto.CreateQRCodeRequest) ([]byte, *utils.Response, error) {
	token, _, err := mp.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CreateQRCodeResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return response.Body, response, nil
}
