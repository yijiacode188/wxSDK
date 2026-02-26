package materialTemporary

import (
	"errors"
	"github.com/yijiacode188/wxSDK/service/handler/materialTemporary/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"net/url"
)

// GetHDVoice 获取高清语音素材
// 本接口用于获取从 JSSDK 的 uploadVoice 接口上传的临时语音素材，格式为speex，16K采样率。该音频比临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。
// https://developers.weixin.qq.com/doc/service/api/material/temporary/api_gethdvoice.html
func (wx *MaterialTemporary) GetHDVoice(mediaId string) (*utils.Response, error) {
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)
	params.Add("media_id", mediaId)
	result, response, err := utils.HttpGet[vo.GetMediaResponse](&utils.RequestParams{
		Url:    "https://api.weixin.qq.com/cgi-bin/media/get/jssdk",
		Params: params,
	})

	if err != nil {
		return response, err
	}
	if result.ErrCode != 0 {
		return response, errors.New(result.ErrMsg)
	}
	return response, nil
}
