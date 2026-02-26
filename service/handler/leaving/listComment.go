package leaving

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/dto"
	"github.com/yijiacode188/wxSDK/service/handler/leaving/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// ListComment 查看指定文章的评论数据
// 本接口用于查看指定文章的评论数据
// https://developers.weixin.qq.com/doc/service/api/leaving/api_listcomment.html
func (wx *Leaving) ListComment(body *dto.ListCommentRequest) (*vo.ListCommentResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.ListCommentResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/comment/list",
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
