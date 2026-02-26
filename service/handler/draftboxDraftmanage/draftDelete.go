package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftDelete 删除草稿
// 删除指定草稿，节省空间。此操作不可撤销，请谨慎操作。
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_delete.html
func (wx *DraftBoxDraftManage) DraftDelete(body *dto.DraftDeleteRequest) (*vo.DraftDeleteResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DraftDeleteResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/delete",
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
