package webDevJssdk

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/consts"
	"github.com/yijiacode188/wxSDK/service/handler/webDevJssdk/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetTicket 获取sdk临时票据
// Api_ticket 是用于调用 js-sdk 的临时票据，有效期为7200 秒，通过access_token 来获取。
// https://developers.weixin.qq.com/doc/service/api/webdev/jssdk/api_getticket.html
func (wx *WebDevJsSdk) GetTicket(ticketType consts.TicketType) (*vo.GetTicketResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("type", string(ticketType))
	result, response, err := utils.HttpGet[vo.GetTicketResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/ticket/getticket",
		Params: params,
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
