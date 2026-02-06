package notifyAutoReplies

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyAutoReplies/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetCurrentAutoReplyInfo 获取自动回复规则
// 获取公众号当前使用的自动回复规则，包括关注后自动回复、消息自动回复、关键词自动回复
// https://developers.weixin.qq.com/doc/service/api/notify/autoreplies/api_getcurrentautoreplyinfo.html
func (wx *NotifyAutoReplies) GetCurrentAutoReplyInfo() (*vo.GetCurrentAutoReplyInfoResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetCurrentAutoReplyInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info",
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
