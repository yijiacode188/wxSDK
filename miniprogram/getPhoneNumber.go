package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/model/dto"
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetPhoneNumber 获取手机号
// 该接口用于将code换取用户手机号。 说明，每个code只能使用一次，code的有效期为5min。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/user-info/phone-number/api_getphonenumber.html
func (wx *wxClient) GetPhoneNumber(code string) (*vo.GetPhoneNumberResponse, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.GetPhoneNumberResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxa/business/getuserphonenumber",
		Params: params,
		Body:   (&dto.GetPhoneNumberRequest{Code: code}).ToByte(),
	})
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
