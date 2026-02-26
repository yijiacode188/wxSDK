package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftUpdate 更新草稿
// 本接口用于修改图文或图片消息草稿。
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_update.html
func (wx *DraftBoxDraftManage) DraftUpdate(body *dto.DraftUpdateRequest) (*vo.DraftUpdateResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)

	result, response, err := utils.HttpPost[vo.DraftUpdateResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/update",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
