package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftCount 获取草稿的总数
// 获取草稿的总数，此接口只统计数量，不返回草稿的具体内容
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_count.html
func (wx *DraftBoxDraftManage) DraftCount() (*vo.DraftCountResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.DraftCountResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/count",
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
