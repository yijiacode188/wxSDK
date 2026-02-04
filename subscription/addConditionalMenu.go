package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// AddConditionalMenu 创建个性化菜单
// 为了帮助公众号实现灵活的业务运营，微信公众平台新增了个性化菜单接口，开发者可以通过该接口，让公众号的不同用户群体看到不一样的自定义菜单。
// 开发者可以通过以下条件来设置用户看到的菜单：
// 用户标签（开发者的业务需求可以借助用户标签来完成）
// 使用普通自定义菜单查询接口可以获取默认菜单和全部个性化菜单信息，请见自定义菜单查询接口的说明。
// 使用普通自定义菜单删除接口可以删除所有自定义菜单（包括默认菜单和全部个性化菜单），请见自定义菜单删除接口的说明。
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_addconditionalmenu.html
func (wx *wxClient) AddConditionalMenu(body *dto.AddConditionalMenuRequest) (string, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return "", err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.AddConditionalMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/addconditional",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", errors.New(result.ErrMsg)
	}
	return result.MenuId, nil
}
