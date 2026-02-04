package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetRidInfo 查询rid信息
// 本接口用于查询调用服务端接口报错返回的rid详情信息，辅助开发者高效定位问题。
// 适用的账号类型如下：公众号/服务号/小程序/小游戏/微信小店/带货助手/视频号助手/联盟带货机构/移动应用/网站应用/多端应用/第三方平台。
// https://developers.weixin.qq.com/doc/subscription/api/apimanage/api_getridinfo.html
func (wx *wxClient) GetRidInfo(body *dto.GetRidInfoRequest) (*vo.GeRidInfoResponse, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.GeRidInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/openapi/rid/get",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
