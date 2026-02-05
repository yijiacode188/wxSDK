package apimanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/apimanage/dto"
	"github.com/yijiacode188/wxSDK/service/hanlder/apimanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// ClearApiQuota 重置指定API调用次数
// 本接口使用 access_token 来重置指定接口的每日调用次数
// https://developers.weixin.qq.com/doc/subscription/api/apimanage/api_clearapiquota.html
func (apiManager *ApiManager) ClearApiQuota(body *dto.ClearApiQuotaRequest) (*utils.Response, error) {
	token, _, err := apiManager.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.ClearApiQuotaResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/openapi/quota/clear",
		Body:   body.ToByte(),
		Params: params,
	})
	if err != nil {
		return response, err
	}
	if result.ErrCode != 0 {
		return response, errors.New(result.ErrMsg)
	}
	return response, nil
}
