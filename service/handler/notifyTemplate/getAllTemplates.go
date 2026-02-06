package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetAllTemplates 获取已选用模板列表
// 本接口用于获取已添加至账号下所有模板列表
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_getalltemplates.html
func (wx *NotifyTemplate) GetAllTemplates() (*vo.GetAllTemplatesResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetAllTemplatesResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template",
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
