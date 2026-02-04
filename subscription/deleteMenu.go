package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteMenu 删除自定义菜单
// 删除当前使用的自定义菜单。注意：调用此接口会删除默认菜单及全部个性化菜单。
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_deletemenu.html
func (wx *wxClient) DeleteMenu() error {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpGet[vo.DeleteMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/delete",
		Params: params,
	})
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}
