package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CreateCustomMenu 创建自定义菜单
// 该接口用于创建公众号/服务号的自定义菜单。
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_createcustommenu.html
func (wx *wxClient) CreateCustomMenu(body *dto.CreateCustomMenuRequest) error {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.CreateCustomMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/create",
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
