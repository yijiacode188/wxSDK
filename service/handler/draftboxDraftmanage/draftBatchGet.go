package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftBatchGet 获取草稿列表
// 新增草稿之后，可用本接口获取草稿列表信息。
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_batchget.html
func (wx *DraftBoxDraftManage) DraftBatchGet(body *dto.DraftBatchGetRequest) (*vo.DraftBatchGetResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DraftBatchGetResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/batchget",
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
