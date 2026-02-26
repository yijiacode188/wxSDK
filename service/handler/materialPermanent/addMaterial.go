package materialPermanent

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/dto"
	"github.com/yijiacode188/wxSDK/service/handler/materialPermanent/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"mime/multipart"
	"net/url"
)

type MediaType string

const (
	MediaTypeImg   MediaType = "image" //图片
	MediaTypeVoice MediaType = "voice" //语音
	MediaTypeVideo MediaType = "video" //视频
	MediaTypeThumb MediaType = "thumb" //缩略图
)

// AddMaterial 上传永久素材
// 本接口用于新增图片/语音/视频等类型的永久素材。
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_addmaterial.html
func (wx *MaterialPermanent) AddMaterial(fileHeader *multipart.FileHeader, data *dto.AddMaterialRequest, mediaType MediaType) (*vo.AddMaterialResponse, *utils.Response, error) {
	body, contentType, err := CreateMultipartFormData("media", fileHeader, func(writer *multipart.Writer) error {
		return writer.WriteField("description", string(data.ToByte()))
	})
	if err != nil {
		return nil, nil, err
	}
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("type", string(mediaType))
	result, response, err := utils.HttpPost[vo.AddMaterialResponse](&utils.RequestParams{
		Url:         "https://api.weixin.qq.com/cgi-bin/material/add_material",
		ContentType: contentType,
		Params:      params,
		Body:        body,
	})
	if err != nil {
		return nil, response, err
	}
	if result.ErrCode != 0 {
		return nil, response, errors.New(result.ErrMsg)
	}
	return &result, response, nil
}
