package public

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/public/dto"
	"github.com/yijiacode188/wxSDK/service/handler/public/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FreePublishGetArticle 获取已发布图文信息
// 本接口用于获取已发布的图文信息
// https://developers.weixin.qq.com/doc/service/api/public/api_freepublishgetarticle.html
func (wx *Public) FreePublishGetArticle(body *dto.FreePublishGetArticleRequest) (*vo.FreePublishGetArticleResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FreePublishGetArticleResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/freepublish/getarticle",
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
