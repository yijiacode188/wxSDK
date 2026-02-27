package qrCodeQrCodeJump

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/dto"
	"github.com/yijiacode188/wxSDK/service/handler/qrCodeQrCodeJump/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// QrCodeJumpPublish 发布已设置的二维码规则
// 需要先调接口添加二维码规则，然后调用本接口将二维码规则发布生效，发布后用户扫码命中该规则的二维码时将跳转到正式版小程序指定的页面
// https://developers.weixin.qq.com/doc/service/api/qrcode/qrcodejump/api_qrcodejumppublish.html
func (wx *QrCodeQrCodeJump) QrCodeJumpPublish(body *dto.QrCodeJumpPublishRequest) (*vo.QrCodeJumpPublishResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.QrCodeJumpPublishResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish",
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
