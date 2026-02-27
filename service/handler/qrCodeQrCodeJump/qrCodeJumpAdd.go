package qrCodeQrCodeJump

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// QrCodeJumpAdd 增加或修改二维码规则
// 该接口用于增加或修改二维码规则
// https://developers.weixin.qq.com/doc/service/api/qrcode/qrcodejump/api_qrcodejumpadd.html
func (wx *QrCodeQrCodeJump) QrCodeJumpAdd(body *dto.QrCodeJumpAddRequest) (*vo.QrCodeJumpAddResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.QrCodeJumpAddResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpadd",
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
