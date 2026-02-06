package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// AddTemplate 选用模板
// 本接口用于从类目模板库或行业模板库添加模板获得模板ID
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_addtemplate.html
func (wx *NotifyTemplate) AddTemplate(body *dto.AddTemplateRequest) (*vo.AddTemplateResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.AddTemplateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/template/api_add_template",
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
