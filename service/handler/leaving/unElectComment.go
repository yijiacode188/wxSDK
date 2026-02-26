package leaving

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/dto"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// UnElectComment 评论取消精选
// 本接口将评论取消精选
// https://developers.weixin.qq.com/doc/service/api/leaving/api_unelectcomment.html
func (wx *Leaving) UnElectComment(body *dto.UnElectCommentRequest) (*vo.UnElectCommentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.UnElectCommentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/comment/unmarkelect",
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
