package miniprogram

import (
	"errors"
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// Code2Session 小程序登录凭证校验
// 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见小程序登录。
// https://developers.weixin.qq.com/miniprogram/dev/server/API/user-login/api_code2session.html
func (wx *wxClient) Code2Session(code string) (*vo.Code2SessionResponse, error) {
	params := make(url.Values)
	params.Add("appid", wx.AppId)
	params.Add("secret", wx.Secret)
	params.Add("js_code", code)
	params.Add("grant_type", "authorization_code")
	result, _, err := utils.HttpGet[vo.Code2SessionResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/sns/jscode2session",
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
