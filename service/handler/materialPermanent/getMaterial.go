package materialPermanent

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/dto"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetMaterial 获取永久素材
// 本接口用于根据media_id获取永久素材的详细信息
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_getmaterial.html
func (wx *MaterialPermanent) GetMaterial(body *dto.GetMaterialRequest) (*vo.GetMaterialResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.GetMaterialResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/material/get_material",
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
