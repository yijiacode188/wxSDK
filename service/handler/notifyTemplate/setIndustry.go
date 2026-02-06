package notifyTemplate

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/dto"
	"github.com/yijiacode188/wxSDK/service/handler/notifyTemplate/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// SetIndustry 设置所属行业
// 本接口用于修改账号所属行业，修改行业后，你在原有行业中的模板将会被删除。
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_setindustry.html
func (wx *NotifyTemplate) SetIndustry(body *dto.SetIndustryRequest) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.SetIndustryResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/template/api_set_industry",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return response, err
	}
	if result.ErrCode != 0 {
		return response, errors.New(result.ErrMsg)
	}
	return response, nil
}
