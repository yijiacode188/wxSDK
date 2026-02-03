package miniprogram

import (
	"errors"
	"net/url"
	"wxSDK/miniprogram/model/dto"
	"wxSDK/miniprogram/model/vo"
	"wxSDK/utils"
)

// GetUnlimitedQRCode 获取不限制的小程序码
// 该接口用于获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取小程序码。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_getunlimitedqrcode.html
func (wx *wxClient) GetUnlimitedQRCode(body *dto.GetUnlimitedQRCodeRequest) ([]byte, error) {
	token, err := wx.getAccessToken()
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetUnlimitedQRCodeResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/getwxacodeunlimit",
		Params: params,
		Body:   body.ToByte(),
	})
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return response.Body, nil
}
