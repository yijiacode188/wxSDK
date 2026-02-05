package vo

import (
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
)

type GetPhoneNumberResponse struct {
	vo.Base
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     string `json:"countryCode"`
	} `json:"phone_info"`
}
