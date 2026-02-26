package draftboxDraftmanage

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/dto"
	"github.com/yijiacode188/wxSDK/service/handler/draftboxDraftmanage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DraftAdd 新增草稿
// 本接口用于新增常用的素材到草稿箱。
// 1.上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除 2.新增草稿也可在公众平台官网-草稿箱中查看和管理
// https://developers.weixin.qq.com/doc/service/api/draftbox/draftmanage/api_draft_add.html
func (wx *DraftBoxDraftManage) DraftAdd(body *dto.DraftAddRequest) (*vo.DraftAddResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DraftAddResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/draft/add",
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
