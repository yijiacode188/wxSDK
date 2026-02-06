package base

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/base/dto"
	"github.com/yijiacode188/wxSDK/service/handler/base/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CallBackCheck 网络通信检测
// 为了帮助开发者排查回调连接失败的问题，提供这个网络检测的API。它可以对开发者URL做域名解析，然后对所有IP进行一次ping操作，得到丢包率和耗时。
// https://developers.weixin.qq.com/doc/service/api/base/api_callbackcheck.html
func (base *Base) CallBackCheck(body *dto.CallBackCheckRequest) (*vo.CallBackCheckResponse, *utils.Response, error) {
	token, _, err := base.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CallBackCheckResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/callback/check",
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
