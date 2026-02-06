package notifyNotify

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/hanlder/notifyNotify/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetCategory 获取类目
// 本接口用于获取小程序、公众号所属类目用于查询公共模板
// https://developers.weixin.qq.com/doc/service/api/notify/notify/api_getcategory.html
func (wx *NotifyNotify) GetCategory() (*vo.GetCategoryResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetCategoryResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory",
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
