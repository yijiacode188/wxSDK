package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetApiQuota 查询API调用额度
// 本接口用于查询服务端接口的的每日调用接口的额度，调用次数，频率限制。
// 适用账号类型：公众号/服务号/小程序/小游戏/微信小店/带货助手/视频号助手/联盟带货机构/移动应用/网站应用/多端应用/第三方平台等接口
// https://developers.weixin.qq.com/doc/subscription/api/apimanage/api_getapiquota.html#Res__component_rate_limit
func (wx *wxClient) GetApiQuota(body *dto.GetApiQuotaRequest) (*vo.GetApiQuotaResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetApiQuotaResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/openapi/quota/get",
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
