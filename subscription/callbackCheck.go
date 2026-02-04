package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CallBackCheck 网络通信检测
// 为了帮助开发者排查回调连接失败的问题，提供这个网络检测的API。它可以对开发者URL做域名解析，然后对所有IP进行一次ping操作，得到丢包率和耗时。
// https://developers.weixin.qq.com/doc/service/api/base/api_callbackcheck.html
func (wx *wxClient) CallBackCheck(body *dto.CallBackCheckRequest) (*vo.CallBackCheckResponse, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.CallBackCheckResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/callback/check",
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
