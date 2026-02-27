package qrCodeQrcodes

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrcodes/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrcodes/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateQRCode 生成带参数的二维码
// 创建二维码ticket，用于生成带参数的二维码
// https://developers.weixin.qq.com/doc/service/api/qrcode/qrcodes/api_createqrcode.html
func (wx *QrCodeQrCodes) CreateQRCode(body *dto.CreateQRCodeRequest) (*vo.CreateQRCodeResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CreateQRCodeResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/qrcode/create",
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
