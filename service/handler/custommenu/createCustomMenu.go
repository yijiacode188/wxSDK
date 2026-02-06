package custommenu

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/custommenu/dto"
	"github.com/yijiacode188/wxSDK/service/handler/custommenu/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateCustomMenu 创建自定义菜单
// 该接口用于创建公众号/服务号的自定义菜单。
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_createcustommenu.html
func (menu *CustomMenu) CreateCustomMenu(body *dto.CreateCustomMenuRequest) (*utils.Response, error) {
	token, _, err := menu.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CreateCustomMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/create",
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
