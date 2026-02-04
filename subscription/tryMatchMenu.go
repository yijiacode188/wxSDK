package subscription

import (
	"errors"
	"github.com/yijiacode188/wxSDK/subscription/model/dto"
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// TryMatchMenu 测试个性化菜单匹配结果
// 测试用户看到的菜单配置
// https://developers.weixin.qq.com/doc/subscription/api/custommenu/api_trymatchmenu.html
func (wx *wxClient) TryMatchMenu(body *dto.TryMatchMenuRequest) (*vo.TryMatchMenuResponse, error) {
	token, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, _, err := utils.HttpPost[vo.TryMatchMenuResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/menu/trymatch",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
