package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// ClearQuotaByAppSecret 使用AppSecret重置API调用次数
// 本接口是通过AppSecret清空服务端接口的每日调用接口次数。
// 适用的账号类型为：公众号/服务号/小程序/小游戏/微信小店/带货助手/视频号助手/联盟带货机构/移动应用/网站应用/多端应用/第三方平台
// https://developers.weixin.qq.com/doc/subscription/api/apimanage/api_clearquotabyappsecret.html
func (wx *wxClient) ClearQuotaByAppSecret(body *dto.ClearQuotaByAppSecretRequest) error {
	params := make(url.Values)
	params.Add("appid", body.AppID)
	params.Add("appsecret", body.AppSecret)
	result, _, err := utils.HttpPost[vo.ClearQuotaByAppSecretResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/clear_quota/v2",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}
