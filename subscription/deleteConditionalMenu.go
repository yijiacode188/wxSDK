package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteConditionalMenu 删除个性化菜单
// 删除指定个性化菜单
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_deleteconditionalmenu.html
func (wx *wxClient) DeleteConditionalMenu(body *dto.DeleteConditionalMenuRequest) error {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.DeleteConditionalMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/delconditional",
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
