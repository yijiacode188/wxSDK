package materialPermanent

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/dto"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// DelMaterial 删除永久素材
// 本接口用于删除不再需要的永久素材，节省存储空间
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_delmaterial.html
func (wx *MaterialPermanent) DelMaterial(body *dto.DelMaterialRequest) (bool, *utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return false, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	result, response, err := utils.HttpPost[vo.DelMaterialResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/material/del_material",
		Params: params,
		Body:   body.ToByte(),
	})
	if err != nil {
		return false, response, err
	}
	if result.ErrCode != 0 {
		return false, response, errors.New(result.ErrMsg)
	}
	return true, response, nil
}
