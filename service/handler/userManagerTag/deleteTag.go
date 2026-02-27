package userManagerTag

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/dto"
	"github.com/yijiacode188/wxSDK/service/handler/userManagerTag/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DeleteTag 删除标签
// 本接口用于删除已存在的标签信息。
// https://developers.weixin.qq.com/doc/service/api/usermanage/tag/api_deletetag.html
func (wx *UserManagerTag) DeleteTag(body *dto.DeleteTagRequest) (*vo.DeleteTagResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DeleteTagResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/tags/delete",
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
