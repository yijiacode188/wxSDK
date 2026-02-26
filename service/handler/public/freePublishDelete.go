package public

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/public/dto"
	"github.com/yijiacode188/wxSDK/service/handler/public/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// FreePublishDelete 删除发布文章
// 本接口用于删除已发布的文章，此操作不可逆，请谨慎操作
// https://developers.weixin.qq.com/doc/service/api/public/api_freepublishdelete.html
func (wx *Public) FreePublishDelete(body *dto.FreePublishDeleteRequest) (*vo.FreePublishDeleteResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.FreePublishDeleteResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/freepublish/delete",
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
