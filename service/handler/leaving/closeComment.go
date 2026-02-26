package leaving

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/dto"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// CloseComment 关闭已群发文章评论
// 本接口用于关闭已群发文章评论
// https://developers.weixin.qq.com/doc/service/api/leaving/api_closecomment.html
func (wx *Leaving) CloseComment(body *dto.CloseCommentRequest) (*vo.CloseCommentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.CloseCommentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/comment/close",
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
