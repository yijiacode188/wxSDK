package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// AddWXaNewTemplate 选用模板
// 从公共模板库中选用模板到私有模板库
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_addwxanewtemplate.html
func (wx *NotifyNotify) AddWXaNewTemplate(body *dto.AddWXaNewTemplateRequest) (*vo.AddWXaNewTemplateResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.AddWXaNewTemplateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate",
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
