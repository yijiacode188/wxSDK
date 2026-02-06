package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetIndustry 获取行业信息
// 本接口可获取账号设置的行业信息
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_getindustry.html
func (wx *NotifyTemplate) GetIndustry() (*vo.GetIndustryResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetIndustryResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/template/get_industry",
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
