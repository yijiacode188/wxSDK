package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetWXaPubNewTemplate 获取已有模板列表
// 该接口用于获取当前帐号下的已有的模板列表。
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_getwxapubnewtemplate.html
func (wx *NotifyNotify) GetWXaPubNewTemplate() (*vo.GetWXaPubNewTemplateResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetWXaPubNewTemplateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate",
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
