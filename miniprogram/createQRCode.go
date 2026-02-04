package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/model/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateQRCode 获取小程序二维码
// 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_createqrcode.html
func (wx *wxClient) CreateQRCode(body *dto.CreateQRCodeRequest) ([]byte, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetCreateQRCodeResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return response.Body, nil
}
