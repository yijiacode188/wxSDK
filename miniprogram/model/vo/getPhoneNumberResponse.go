package vo

import "github.com/yijiacode188/wxSDK/miniprogram/model/vo/base"

type GetPhoneNumberResponse struct {
	base.Base
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     string `json:"countryCode"`
	} `json:"phone_info"`
}
