package materialPermanent

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetMaterialCount 获取永久素材总数
// 本接口用于获取公众号永久素材的总数信息
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_getmaterialcount.html
func (wx *MaterialPermanent) GetMaterialCount() (*vo.GetMaterialCountResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpGet[vo.GetMaterialCountResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/material/get_materialcount",
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
