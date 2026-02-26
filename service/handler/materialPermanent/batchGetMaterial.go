package materialPermanent

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/dto"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// BatchGetMaterial 获取永久素材列表
// 分类型获取永久素材列表，包含公众号在官网素材管理模块新建的素材
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_batchgetmaterial.html
func (wx *MaterialPermanent) BatchGetMaterial(body *dto.BatchGetMaterialRequest) (*vo.BatchGetMaterialResponse, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.BatchGetMaterialResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/material/batchget_material",
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
