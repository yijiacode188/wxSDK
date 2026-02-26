package leaving

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/dto"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DelComment 删除评论
// 本接口用于删除评论
// https://developers.weixin.qq.com/doc/service/api/leaving/api_delcomment.html
func (wx *Leaving) DelComment(body *dto.DelCommentRequest) (*vo.DelCommentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DelCommentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/comment/delete",
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
