package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetDraft 获取草稿详情
// 新增草稿后，可通过该接口下载草稿
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_getdraft.html
func (wx *DraftBoxDraftManage) GetDraft(body *dto.GetDraftRequest) (*vo.GetDraftResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)

	result, response, err := utils.HttpPost[vo.GetDraftResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/get",
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
