package qrCodeQrCodeJump

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// QrCodeJumpDelete 删除已设置的二维码规则
// 该接口用于删除已设置的二维码规则
// https://developers.weixin.qq.com/doc/service/api/qrcode/qrcodejump/api_qrcodejumpdelete.html
func (wx *QrCodeQrCodeJump) QrCodeJumpDelete(body *dto.QrCodeJumpDeleteRequest) (*vo.QrCodeJumpDeleteResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.QrCodeJumpDeleteResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdelete",
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
