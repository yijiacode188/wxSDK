package custommenu

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/handler/custommenu/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetMenu 获取自定义菜单配置
// 使用接口创建自定义菜单后，开发者还可使用接口查询自定义菜单的结构
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_getmenu.html
func (menu *CustomMenu) GetMenu() (*vo.GetMenuResponse, *utils.Response, error) {
	token, _, err := menu.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/get",
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
