package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetCurrentSelfMenuInfo 查询自定义菜单信息
// 本接口提供公众号当前使用的自定义菜单的配置，如果公众号是通过API调用设置的菜单，则返回菜单的开发配置，而如果公众号是在公众平台官网通过网站功能发布菜单，则本接口返回运营者设置的菜单配置
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_getcurrentselfmenuinfo.html
func (wx *wxClient) GetCurrentSelfMenuInfo() (*vo.GetCurrentSelfMenuInfoResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetCurrentSelfMenuInfoResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info",
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
