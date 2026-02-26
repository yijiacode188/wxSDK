package leaving

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/dto"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DelReplyComment 删除回复
// 本接口用于删除评论的回复内容
// https://developers.weixin.qq.com/doc/service/api/leaving/api_delreplycomment.html
func (wx *Leaving) DelReplyComment(body *dto.DelReplyCommentRequest) (*vo.DelReplyCommentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DelReplyCommentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/comment/reply/delete",
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
