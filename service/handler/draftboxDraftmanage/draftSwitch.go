package draftboxDraftmanage

import (
	"errors"
	"fmt"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftSwitch 草稿箱开关设置
// 本接口用于设置或查询草稿箱和发布功能的开关状态。
//
// 内测期间会逐步放量，任何用户都可能会自动打开；
// 此开关开启后不可逆，换言之，无法从开启的状态回到关闭；
// 内测期间，无论开关开启与否，旧版的图文素材API，以及新版的草稿箱、发布等API均可以正常使用；
// 在内测结束之后，所有用户都将自动开启，即草稿箱、发布等功能将对所有用户开放，本开关连同之前的图文素材API也将随后下线
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_switch.html
func (wx *DraftBoxDraftManage) DraftSwitch(checkOnly *int) (*vo.DraftSwitchResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	if checkOnly != nil {
		params.Add("checkonly", fmt.Sprintf("%d", *checkOnly))
	}
	result, response, err := utils.HttpPost[vo.DraftSwitchResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/switch",
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
