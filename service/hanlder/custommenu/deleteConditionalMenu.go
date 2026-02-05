package custommenu

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/custommenu/dto"
	"github.com/yijiacode188/wxSDK/service/hanlder/custommenu/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteConditionalMenu 删除个性化菜单
// 删除指定个性化菜单
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_deleteconditionalmenu.html
func (menu *CustomMenu) DeleteConditionalMenu(body *dto.DeleteConditionalMenuRequest) (*utils.Response, error) {
	token, _, err := menu.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DeleteConditionalMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/delconditional",
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
